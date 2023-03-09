package worker

import (
	"bytes"
	"fmt"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"io"
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

	defer resp.Body.Close()
	defer client.CloseIdleConnections()

	if resp.StatusCode != 200 {
		b, _ := io.ReadAll(resp.Body)
		errStr := string(b)
		return fmt.Errorf("upload failed: %s\n%s\n%s", s3ResignedUrl.URL, s3ResignedUrl.Method, errStr)
	}

	return nil
}
