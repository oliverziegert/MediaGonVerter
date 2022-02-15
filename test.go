package main

import (
	"fmt"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"pc-ziegert.de/media_service/common/log"
	"pc-ziegert.de/media_service/common/model"
	"pc-ziegert.de/media_service/worker"
	"sync"
)

func main() {
	c := []*model.Conversion{{
		Width:       2000,
		Height:      2000,
		Crop:        true,
		S3UploadUrl: "",
	},
		{
			Width:       2000,
			Height:      2000,
			Crop:        false,
			S3UploadUrl: "",
		}}
	i := model.Image{
		TenantId:      "asd",
		NodeId:        72825,
		ServiceName:   "asd",
		NodeType:      "image/jpeg",
		S3DownloadUrl: "https://cloud.s3.pc-ziegert.de/5/8aa-c4c/0000000000000072825-483036554463664b4a4d74753478526b?response-content-disposition=attachment%3Bfilename%2A%3DUTF-8%27%271211529591426352.jpg&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20220215T202804Z&X-Amz-SignedHeaders=host&X-Amz-Expires=17999&X-Amz-Credential=cloud%2F20220215%2Feu-central-1%2Fs3%2Faws4_request&X-Amz-Signature=4da8d5f432cddbbda09a2174f1fdc7ec0759e64343940d9408bdd265647aed74",
		NodeSize:      315902,
		State:         0,
		Conversions:   c,
	}
	convert2(&i)
}

func convert2(i *model.Image) {
	var wg sync.WaitGroup

	tds := fmt.Sprintf("./%v/%v-%d", "worker", i.TenantId, i.NodeId)

	if _, err := os.Stat(tds); os.IsNotExist(err) {
		err := os.Mkdir(tds, os.ModePerm)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	resp, err := http.Get(i.S3DownloadUrl)
	if err != nil {
		log.Fatal(err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(err.Error())
	}

	out, err := os.Create(tds + "/original")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Write the body to file
	n, err := io.Copy(out, resp.Body)
	if int64(i.NodeSize) != n {
		log.Fatal("size mismatch")
	}

	resp.Body.Close()
	out.Close()

	for _, c := range i.Conversions {
		wg.Add(1)
		go convertFile2(&wg, i, c)
	}

	wg.Wait()
}

func convertFile2(wg *sync.WaitGroup, i *model.Image, c *model.Conversion) {
	// call done when finished
	defer wg.Done()

	tds := fmt.Sprintf("./%v/%v-%d", "worker", i.TenantId, i.NodeId)

	// open files
	file, err := os.Open(tds + "/original")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	newFilePath := fmt.Sprintf("%v/%dx%d-%t", tds, c.Width, c.Height, c.Crop)
	outFile := openOrCreate2(newFilePath)
	defer outFile.Close()

	// decode
	imageData, _, err := image.Decode(file)
	if err != nil {
		log.Fatal("error decoding image")
	}

	croppedImg, err := worker.Crop(imageData, worker.Config{
		Width:   c.Width,
		Height:  c.Height,
		Mode:    worker.Centered,
		Options: worker.Ratio,
	})

	// encode in new type
	switch "image/jpeg" {
	case "image/jpeg":
		// Set the expected size that you want:
		//rect := image.Rect(0, 0, int(c.Width), int(c.Height))

		//dst := image.NewRGBA(rect)

		//draw.Draw(dst, dst.Rect, &image.Uniform{C: image.White}, image.Point{}, draw.Src)

		// Resize:
		//draw.NearestNeighbor.Scale(dst, dst.Rect, imageData, imageData.Bounds(), draw.Over, nil)
		//draw.NearestNeighbor
		//draw.ApproxBiLinear
		//draw.BiLinear
		//draw.CatmullRom

		opt := jpeg.Options{Quality: 80}
		err := jpeg.Encode(outFile, croppedImg, &opt)
		if err != nil {
			log.Fatal("error converting to jpeg")
		}
	case "image/png":
		err := png.Encode(outFile, imageData)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "image/gif":
		err := gif.Encode(outFile, imageData, nil)
		if err != nil {
			log.Fatal("error converting to gif")
		}
	case "image/bmp":
		err := bmp.Encode(outFile, imageData)
		if err != nil {
			log.Fatal("error converting to bmp")
		}
	case "image/tiff":
		err := tiff.Encode(outFile, imageData, nil)
		if err != nil {
			log.Fatal("error converting to tiff")
		}
	default:
		log.Fatal("unsupported format")
	}
}

func openOrCreate2(filename string) *os.File {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal("error creating output file")
		}
		return file
	} else {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal("error opening output file")
		}
		return file
	}
}
