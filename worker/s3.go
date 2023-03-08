package worker

import (
	"bytes"
	"errors"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"net/http"
)

func UploadFile(rawBody *[]byte, s3ResignedUrl *v4.PresignedHTTPRequest, contentType string) error {

	client := &http.Client{}
	req, err := http.NewRequest(s3ResignedUrl.Method, s3ResignedUrl.URL, bytes.NewReader(*rawBody))
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
