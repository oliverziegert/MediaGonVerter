package utils

import (
	"fmt"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"pc-ziegert.de/media_service/common/config"
	error2 "pc-ziegert.de/media_service/common/error"
	"pc-ziegert.de/media_service/common/log"
	"time"
)

func GenerateS3PresignUploadUrl(ctx *gin.Context, cnf *config.Config, key string) (*v4.PresignedHTTPRequest, *error2.Error) {

	scp := credentials.NewStaticCredentialsProvider(cnf.Data.S3.AccessKeyId, cnf.Data.S3.SecretAccessKey, "")

	client := s3.New(s3.Options{
		Credentials:      scp,
		Region:           cnf.Data.S3.Region,
		EndpointResolver: s3.EndpointResolverFromURL(cnf.Data.S3.Address),
	})
	ex, err := time.ParseDuration(cnf.Data.Expiration)
	if err != nil {
		//ToDo: Error handling
	}
	cacheControl := fmt.Sprintf("max-age=%d", int(ex.Seconds()))
	expires := time.Now().Add(time.Minute * 15)
	poi := &s3.PutObjectInput{
		Bucket:           &cnf.Data.S3.Bucket,
		Key:              &key,
		Body:             nil,
		BucketKeyEnabled: false,
		CacheControl:     &cacheControl,
		Expires:          &expires,
	}

	psClient := s3.NewPresignClient(client)

	resp, err := psClient.PresignPutObject(ctx, poi)
	if err != nil {
		err := error2.WrapError(error2.ValIdInvalid, "Failed to register a consumer.", err)
		log.Debug(err.StackTrace())
		return nil, err
	}
	return resp, nil
}
