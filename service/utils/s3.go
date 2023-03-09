package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"pc-ziegert.de/media_service/common/config"
	e "pc-ziegert.de/media_service/common/error"
	l "pc-ziegert.de/media_service/common/log"
	"time"
)

func GenerateS3PresignUploadUrl(ctx *gin.Context, cfg *config.Config, key string, contentType string) (*v4.PresignedHTTPRequest, *time.Time, *e.Error) {

	scp := credentials.NewStaticCredentialsProvider(cfg.Data.S3.AccessKeyId, cfg.Data.S3.SecretAccessKey, "")

	client := s3.New(s3.Options{
		Credentials:      scp,
		Region:           cfg.Data.S3.Region,
		EndpointResolver: s3.EndpointResolverFromURL(cfg.Data.S3.Address),
		UsePathStyle:     true,
	})

	// max-age: response can be used up to MAX-AGE, before it gets 'stale'
	// private: only store in local cache. We use this as the images might be somewhat sensitive
	// no-transform: do not convert the images in cache (e.g. to reduce size).
	// stale-if-error: a stale object can be reused for X if a refresh attempt failed because of 5XX
	cacheControl := fmt.Sprintf("private, no-transform, max-age=%d, stale-if-error=86400", int(cfg.Data.ExpirationDuration.Seconds()))
	expires := time.Now().Add(cfg.Data.ExpirationDuration)
	poi := &s3.PutObjectInput{
		Bucket:       &cfg.Data.S3.Bucket,
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

func GenerateS3PresignedDownloadUrl(ctx *gin.Context, cfg *config.Config, key string, contentType string) (*v4.PresignedHTTPRequest, *e.Error) {
	scp := credentials.NewStaticCredentialsProvider(cfg.Data.S3.AccessKeyId, cfg.Data.S3.SecretAccessKey, "")

	client := s3.New(s3.Options{
		Credentials:      scp,
		Region:           cfg.Data.S3.Region,
		EndpointResolver: s3.EndpointResolverFromURL(cfg.Data.S3.Address),
		UsePathStyle:     true,
	})

	bucket := cfg.Data.S3.Bucket

	// 'refresh' timestamp so that the file expiration gets reset
	coi := s3.CopyObjectInput{
		Bucket:     &cfg.Data.S3.Bucket,
		CopySource: aws.String(fmt.Sprintf("%s/%s", bucket, key)),
		Key:        &key,
	}

	if _, err := client.CopyObject(ctx, &coi); err != nil {
		l.Infof("could not refresh %s/%s: %v", bucket, key, err)
	}

	goi := &s3.GetObjectInput{
		Bucket: &cfg.Data.S3.Bucket,
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
