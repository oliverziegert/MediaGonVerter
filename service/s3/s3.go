package s3

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
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

	// max-age: response can be used up to MAX-AGE, before it gets 'stale'
	// private: only store in local cache. We use this as the images might be somewhat sensitive
	// no-transform: do not convert the images in cache (e.g. to reduce size).
	// stale-if-error: a stale object can be reused for X if a refresh attempt failed because of 5XX
	cacheControl := fmt.Sprintf("private, no-transform, max-age=%d, stale-if-error=86400", int(cfg.Data.ExpirationDuration.Seconds()))
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

func GenerateS3PresignedDownloadUrl(ctx *gin.Context, s3Config *m.S3Config, key string) (*v4.PresignedHTTPRequest, *e.Error) {

	scp := credentials.NewStaticCredentialsProvider(s3Config.AccessKeyId, s3Config.SecretAccessKey, "")

	client := s3.New(s3.Options{
		Credentials:      scp,
		Region:           s3Config.Region,
		EndpointResolver: s3.EndpointResolverFromURL(s3Config.Address),
		UsePathStyle:     true,
	})

	// 'refresh' timestamp so that the file expiration gets reset
	coi := s3.CopyObjectInput{
		Bucket:     &s3Config.Bucket,
		CopySource: aws.String(fmt.Sprintf("%s/%s", s3Config.Bucket, key)),
		Key:        &key,
	}

	if _, err := client.CopyObject(ctx, &coi); err != nil {
		l.Infof("could not refresh %s/%s: %v", s3Config.Bucket, key, err)
	}

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
