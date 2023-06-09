package mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	l "pc-ziegert.de/media_service/common/log"
	"sync/atomic"
	"time"
)

const delay = 3 // reconnect after delay seconds

// Connection amqp.Connection wrapper
type Connection struct {
	*amqp.Connection
}

// Channel wrap amqp.Connection.Channel, get a auto reconnect channel
func (c *Connection) Channel() (*Channel, error) {
	ch, err := c.Connection.Channel()
	if err != nil {
		return nil, err
	}

	channel := &Channel{
		Channel: ch,
	}

	go func() {
		for {
			reason, ok := <-channel.Channel.NotifyClose(make(chan *amqp.Error))
			// exit this goroutine if closed by developer
			if !ok || channel.IsClosed() {
				l.Debug("channel closed")
				_ = channel.Close() // close again, ensure closed flag set when connection closed
				break
			}
			l.Debugf("channel closed, reason: %v", reason)

			// reconnect if not closed by developer
			for {
				// wait 1s for connection reconnect
				time.Sleep(delay * time.Second)

				ch, err := c.Connection.Channel()
				if err == nil {
					l.Debug("channel recreate success")
					channel.Channel = ch
					break
				}

				l.Debugf("channel recreate failed, err: %v", err)
			}
		}

	}()

	return channel, nil
}

// DialCluster with reconnect
func DialCluster(urls []string) (*Connection, error) {
	nodeSequence := 0
	conn, err := amqp.Dial(urls[nodeSequence])

	if err != nil {
		return nil, err
	}
	connection := &Connection{
		Connection: conn,
	}

	go func(urls []string, seq *int) {
		for {
			reason, ok := <-connection.Connection.NotifyClose(make(chan *amqp.Error))
			if !ok {
				l.Debug("connection closed")
				break
			}
			l.Debugf("connection closed, reason: %v", reason)

			// reconnect with another node of cluster
			for {
				time.Sleep(time.Duration(delay) * time.Second)

				newSeq := next(urls, *seq)
				*seq = newSeq

				conn, err := amqp.Dial(urls[newSeq])
				if err == nil {
					connection.Connection = conn
					l.Debug("reconnect success")
					break
				}

				l.Debugf("reconnect failed, err: %v", err)
			}
		}
	}(urls, &nodeSequence)

	return connection, nil
}

// Next element index of slice
func next(s []string, lastSeq int) int {
	length := len(s)
	if length == 0 || lastSeq == length-1 {
		return 0
	} else if lastSeq < length-1 {
		return lastSeq + 1
	} else {
		return -1
	}
}

// Channel amqp.Channel wrapper
type Channel struct {
	*amqp.Channel
	closed int32
}

// IsClosed indicate closed by developer
func (ch *Channel) IsClosed() bool {
	return atomic.LoadInt32(&ch.closed) == 1
}

// Close ensure closed flag set
func (ch *Channel) Close() error {
	if ch.IsClosed() {
		return amqp.ErrClosed
	}

	atomic.StoreInt32(&ch.closed, 1)

	return ch.Channel.Close()
}

// Consume wrap amqp.Channel.Consume, the returned delivery will end only when channel closed by developer
func (ch *Channel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	deliveries := make(chan amqp.Delivery)

	go func() {
		for {
			d, err := ch.Channel.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
			if err != nil {
				l.Debugf("consume failed, err: %v", err)
				time.Sleep(delay * time.Second)
				continue
			}

			for msg := range d {
				deliveries <- msg
			}

			// sleep before IsClose call. closed flag may not set before sleep.
			time.Sleep(delay * time.Second)

			if ch.IsClosed() {
				break
			}
		}
	}()

	return deliveries, nil
}
