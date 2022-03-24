package worker

import (
	"errors"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"net/http"
	"os"
	"pc-ziegert.de/media_service/common/log"
)

func UploadFile(filePath string, s3ResignedUrl *v4.PresignedHTTPRequest, contentType string) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	client := &http.Client{}
	req, err := http.NewRequest(s3ResignedUrl.Method, s3ResignedUrl.URL, file)
	req.Header = s3ResignedUrl.SignedHeader
	req.Header.Add("Content-Type", contentType)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("upload failed")
	}

	defer resp.Body.Close()
	defer client.CloseIdleConnections()

	return nil
}
