package utils

import (
	"fmt"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"pc-ziegert.de/media_service/common/config"
	e "pc-ziegert.de/media_service/common/error"
	"pc-ziegert.de/media_service/common/log"
	"time"
)

func GenerateS3PresignUploadUrl(ctx *gin.Context, cfg *config.Config, key string, contentType string) (*v4.PresignedHTTPRequest, *time.Time, *e.Error) {

	scp := credentials.NewStaticCredentialsProvider(cfg.Data.S3.AccessKeyId, cfg.Data.S3.SecretAccessKey, "")

	client := s3.New(s3.Options{
		Credentials:      scp,
		Region:           cfg.Data.S3.Region,
		EndpointResolver: s3.EndpointResolverFromURL(cfg.Data.S3.Address),
	})

	cacheControl := fmt.Sprintf("max-age=%d", int(cfg.Data.ExpirationDuration.Seconds()))
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
		log.Debug(err.StackTrace())
		return nil, nil, err
	}
	return resp, &expires, nil
}

func GenerateS3PresignDownloadUrl(ctx *gin.Context, cfg *config.Config, key string, contentType string) (*v4.PresignedHTTPRequest, *e.Error) {

	scp := credentials.NewStaticCredentialsProvider(cfg.Data.S3.AccessKeyId, cfg.Data.S3.SecretAccessKey, "")

	client := s3.New(s3.Options{
		Credentials:      scp,
		Region:           cfg.Data.S3.Region,
		EndpointResolver: s3.EndpointResolverFromURL(cfg.Data.S3.Address),
	})

	goi := &s3.GetObjectInput{
		Bucket: &cfg.Data.S3.Bucket,
		Key:    &key,
	}

	psClient := s3.NewPresignClient(client)

	resp, err := psClient.PresignGetObject(ctx, goi)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
		log.Debug(err.StackTrace())
		return nil, err
	}
	return resp, nil
}
