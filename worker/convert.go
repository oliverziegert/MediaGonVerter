package worker

import (
	"fmt"
	"github.com/davidbyttow/govips/v2/vips"
	amqp "github.com/rabbitmq/amqp091-go"
	"io"
	"log"
	"net/http"
	"pc-ziegert.de/media_service/common/constant"
	e "pc-ziegert.de/media_service/common/error"
	l "pc-ziegert.de/media_service/common/log"
	m "pc-ziegert.de/media_service/common/model"
	"sync"
)

func Convert(in amqp.Delivery, cRes chan<- m.Image) {
	var wg sync.WaitGroup

	defer in.Ack(false)

	if in.Type != constant.RabbitMQImageMessageType {
		l.Fatalf("Wrong message type:", in.Type)
		return
	}

	i := m.Image{}
	i.JsonToImage(in.Body)

	hResp, err := http.Get(i.S3DownloadUrl)
	if err != nil {
		l.Fatal(err.Error())
	}
	if hResp.StatusCode != http.StatusOK {
		l.Fatal(err.Error())
	}
	defer hResp.Body.Close()

	rawBody, err := io.ReadAll(hResp.Body)

	for _, c := range i.Conversions {
		if c.State != m.ConversionStateCached {
			wg.Add(1)
			go convertFile(&wg, &i, c, cRes, &rawBody)
		}
	}

	wg.Wait()
	return
}

func convertFile(wg *sync.WaitGroup, i *m.Image, c *m.Conversion, cRes chan<- m.Image, rawBody *[]byte) {
	// call done when finished
	defer wg.Done()

	// Open image with VIPS
	iImage, err := vips.NewImageFromBuffer(*rawBody)
	if err != nil {
		err := e.NewError(e.ValIdInvalid, fmt.Sprintf("error opening image"))
		l.Debug(err.StackTrace())
		c.State = m.ConversionStateConversionError
		cRes <- *i
		return
	}

	// Resize image
	err = iImage.Thumbnail(int(c.Width), int(c.Height), vips.InterestingAttention)
	if err != nil {
		err := e.NewError(e.ValIdInvalid, fmt.Sprintf("error resizing image"))
		l.Debug(err.StackTrace())
		c.State = m.ConversionStateConversionError
		cRes <- *i
		return
	}

	// Save resized image
	ep := vips.NewJpegExportParams()
	oImage, _, err := iImage.ExportJpeg(ep)
	if err != nil {
		log.Fatal(err.Error())
	}

	// upload image to file
	err = UploadFile(&oImage, c.S3UploadUrl, i.NodeType)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, fmt.Sprintf("Cant upload image to s3"), err)
		l.Debug(err.StackTrace())
		c.State = m.ConversionStateConversionError
		cRes <- *i
		return
	}
	c.State = m.ConversionStateCached
	cRes <- *i
	return
}
