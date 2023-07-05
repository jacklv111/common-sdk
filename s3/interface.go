/*
 * Created on Wed Jul 05 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package s3

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jacklv111/common-sdk/log"
)

var s3Session *session.Session

var S3Uploader *s3manager.Uploader

var S3Downloader *s3manager.Downloader

var S3Deleter *s3manager.BatchDelete

var S3Client *s3.S3

func InitS3() (err error) {
	s3Session, err = session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(S3Config.accessKey, S3Config.secretKey, S3Config.token),
		Endpoint:         aws.String(S3Config.endPoint),
		Region:           aws.String(S3Config.region),
		DisableSSL:       aws.Bool(S3Config.disableSSL),
		S3ForcePathStyle: aws.Bool(S3Config.s3ForcePathStyle),
	})
	if err != nil {
		return err
	}

	S3Uploader = s3manager.NewUploader(s3Session, func(d *s3manager.Uploader) {
		d.PartSize = int64(S3Config.uploadPartSize)
	})
	S3Downloader = s3manager.NewDownloader(s3Session, func(d *s3manager.Downloader) {
		d.PartSize = int64(S3Config.downloadPartSize)
	})
	S3Deleter = s3manager.NewBatchDelete(s3Session)
	S3Client = s3.New(s3Session)
	return nil
}

func Download(objectKey string, localPath string) error {
	file, err := os.Create(localPath)
	if err != nil {
		return err
	}
	n, err := S3Downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(S3Config.Bucket),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return fmt.Errorf("failed to download file, %v", err)
	}
	log.Info("file downloaded, %d bytes", n)
	return nil
}

func GetObjectRequest(objectKey string) (req *request.Request, output *s3.GetObjectOutput) {
	return S3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(S3Config.Bucket),
		Key:    aws.String(objectKey),
	})
}
