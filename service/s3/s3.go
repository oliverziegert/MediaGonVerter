package s3

import (
	"fmt"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	c "pc-ziegert.de/media_service/common/config"
	e "pc-ziegert.de/media_service/common/error"
	l "pc-ziegert.de/media_service/common/log"
	m "pc-ziegert.de/media_service/common/model"
	"pc-ziegert.de/media_service/service/s3/repo"
	"time"
)

type S3 struct {
	sRepo *repo.S3ConfigRepo
}

// NewS3 creates a new S3 repository
func NewS3() *S3 {
	return &S3{}
}

// GetS3ConfigRepo provides the S3ConfigRepo.
func (r *S3) GetS3ConfigRepo() *repo.S3ConfigRepo {
	if r.sRepo == nil {
		r.sRepo = repo.NewS3Repo()
	}
	return r.sRepo
}

func GenerateS3PresignUploadUrl(ctx *gin.Context, cfg *c.Config, s3Config *m.S3Config, key string, contentType string) (*v4.PresignedHTTPRequest, *time.Time, *e.Error) {

	scp := credentials.NewStaticCredentialsProvider(s3Config.AccessKeyId, s3Config.SecretAccessKey, "")

	client := s3.New(s3.Options{
		Credentials:      scp,
		Region:           s3Config.Region,
		EndpointResolver: s3.EndpointResolverFromURL(s3Config.Address),
		UsePathStyle:     true,
	})

	cacheControl := fmt.Sprintf("max-age=%d", int(cfg.Data.ExpirationDuration.Seconds()))
	expires := time.Now().Add(cfg.Data.ExpirationDuration)
	poi := &s3.PutObjectInput{
		Bucket:       &s3Config.Bucket,
		Key:          &key,
		ContentType:  &contentType,
		CacheControl: &cacheControl,
		Expires:      &expires,
	}

	psClient := s3.NewPresignClient(client)

	resp, err := psClient.PresignPutObject(ctx, poi)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
		l.Debug(err.StackTrace())
		return nil, nil, err
	}
	return resp, &expires, nil
}

func GenerateS3PresignDownloadUrl(ctx *gin.Context, s3Config *m.S3Config, key string) (*v4.PresignedHTTPRequest, *e.Error) {

	scp := credentials.NewStaticCredentialsProvider(s3Config.AccessKeyId, s3Config.SecretAccessKey, "")

	client := s3.New(s3.Options{
		Credentials:      scp,
		Region:           s3Config.Region,
		EndpointResolver: s3.EndpointResolverFromURL(s3Config.Address),
		UsePathStyle:     true,
	})

	goi := &s3.GetObjectInput{
		Bucket: &s3Config.Bucket,
		Key:    &key,
	}

	psClient := s3.NewPresignClient(client)

	resp, err := psClient.PresignGetObject(ctx, goi)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
		l.Debug(err.StackTrace())
		return nil, err
	}
	return resp, nil
}
