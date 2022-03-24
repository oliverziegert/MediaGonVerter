package worker

import (
	"fmt"
	"github.com/streadway/amqp"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"os"
	"pc-ziegert.de/media_service/common/constant"
	e "pc-ziegert.de/media_service/common/error"
	l "pc-ziegert.de/media_service/common/log"
	m "pc-ziegert.de/media_service/common/model"
	"sync"
)

func Convert(in amqp.Delivery, cRes chan<- m.Image) {
	var wg sync.WaitGroup
	defer close(cRes)

	defer in.Ack(false)

	if in.Type != constant.RabbitMQImageMessageType {
		l.Fatalf("Wrong message type:", in.Type)
		return
	}

	i := m.Image{}
	i.JsonToImage(in.Body)

	tds := fmt.Sprintf("./%v/%v-%d", "worker", i.TenantUuid, i.NodeId)

	if _, err := os.Stat(tds); os.IsNotExist(err) {
		err := os.Mkdir(tds, os.ModePerm)
		if err != nil {
			l.Fatal(err.Error())
		}
	}
	defer os.RemoveAll(tds)

	hResp, err := http.Get(i.S3DownloadUrl)
	if err != nil {
		l.Fatal(err.Error())
	}
	if hResp.StatusCode != http.StatusOK {
		l.Fatal(err.Error())
	}

	out, err := os.Create(tds + "/original")
	if err != nil {
		l.Fatal(err.Error())
	}

	// Write the body to file
	n, err := io.Copy(out, hResp.Body)
	if int64(i.NodeSize) != n {
		l.Fatal("size mismatch")
	}

	hResp.Body.Close()
	out.Close()

	for _, c := range i.Conversions {
		if c.State != m.ConversionStateCached {
			wg.Add(1)
			go convertFile(&wg, &i, c, cRes)
		}
	}

	wg.Wait()
	return
}

func convertFile(wg *sync.WaitGroup, i *m.Image, c *m.Conversion, cRes chan<- m.Image) {
	// call done when finished
	defer wg.Done()

	tds := fmt.Sprintf("./%v/%v-%d", "worker", i.TenantUuid, i.NodeId)

	// open files
	file, err := os.Open(tds + "/original")
	if err != nil {
		err := e.NewError(e.ValIdInvalid, fmt.Sprintf("error opening filehandler of src image"))
		l.Debug(err.StackTrace())
		c.State = m.ConversionStateConversionError
		cRes <- *i

	}
	defer file.Close()

	newFilePath := fmt.Sprintf("%v/%dx%d-%t", tds, c.Width, c.Height, c.Crop)
	outFile := openOrCreate(newFilePath)

	// decode
	imageData, _, err := image.Decode(file)
	if err != nil {
		err := e.NewError(e.ValIdInvalid, fmt.Sprintf("error decoding image"))
		l.Debug(err.StackTrace())
		c.State = m.ConversionStateConversionError
		cRes <- *i
	}

	// resize
	dst := Resize(imageData, ResizeConfig{
		Width:           int(c.Width),
		Height:          int(c.Height),
		Crop:            c.Crop,
		BackgroundColor: image.White,
	})

	// encode in jpeg
	opt := jpeg.Options{Quality: 80}
	err = jpeg.Encode(outFile, dst, &opt)
	if err != nil {
		err := e.NewError(e.ValIdInvalid, fmt.Sprintf("error converting to jpeg"))
		l.Debug(err.StackTrace())
		c.State = m.ConversionStateConversionError
		cRes <- *i
	}

	// sync image to file
	err = outFile.Close()
	if err != nil {
		err := e.NewError(e.ValIdInvalid, fmt.Sprintf("error writing image to file"))
		l.Debug(err.StackTrace())
		c.State = m.ConversionStateConversionError
		cRes <- *i
	}

	// upload image to file
	err = UploadFile(newFilePath, c.S3UploadUrl, i.NodeType)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, fmt.Sprintf("Cant get tenant state"), err)
		l.Debug(err.StackTrace())
		c.State = m.ConversionStateConversionError
		cRes <- *i
	}
	c.State = m.ConversionStateCached
	cRes <- *i
	return
}

func openOrCreate(filename string) *os.File {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			l.Fatal("error creating output file")
		}
		return file
	} else {
		file, err := os.Open(filename)
		if err != nil {
			l.Fatal("error opening output file")
		}
		return file
	}
}
