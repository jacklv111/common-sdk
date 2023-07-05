/*
 * Created on Wed Jul 05 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package s3

import (
	"github.com/spf13/pflag"
)

type s3Config struct {
	accessKey        string
	secretKey        string
	token            string
	endPoint         string
	region           string
	disableSSL       bool
	s3ForcePathStyle bool

	Bucket           string
	downloadPartSize int
	uploadPartSize   int
}

const (
	S3_CONFIG_PATH = "./conf/s3.json"
)

var S3Config *s3Config

func init() {
	S3Config = &s3Config{
		accessKey: "",
		secretKey: "",
		token:     "",
		endPoint:  "",
		region:    "",
		Bucket:    "",
		// download
		downloadPartSize: 64 * 1024 * 1024, // 64MB per part
		uploadPartSize:   64 * 1024 * 1024, // 64MB per part
		disableSSL:       true,
		s3ForcePathStyle: true,
	}
}

func (config *s3Config) ReadFromFile() error {
	return nil
}

func (config *s3Config) AddFlags(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&config.accessKey, "s3-ak", "", "Value to indicate the access key of s3 service")
	flagSet.StringVar(&config.secretKey, "s3-sk", "", "Value to indicate the secret key of s3 service")
	flagSet.StringVar(&config.token, "s3-token", "", "Value to indicate the token of s3 service")
	flagSet.StringVar(&config.endPoint, "s3-endpoint", "", "Value to indicate the endpoint of s3 service")
	flagSet.StringVar(&config.region, "s3-region", "", "Value to indicate the region of s3 service")
	flagSet.StringVar(&config.Bucket, "s3-bucket", "", "Value to indicate the bucket of s3 service")
	flagSet.IntVar(&config.downloadPartSize, "s3-download-part-size", 64*1024*1024, "Value to indicate the download part size (byte) of s3 service")
	flagSet.IntVar(&config.uploadPartSize, "s3-upload-part-size", 64*1024*1024, "Value to indicate the upload part size (byte) of s3 service")
	flagSet.BoolVar(&config.disableSSL, "s3-disable-ssl", true, "Value to indicate whether to disable ssl of s3 service")
	flagSet.BoolVar(&config.s3ForcePathStyle, "s3-force-path-style", true, "Value to indicate whether to force path style of s3 service")
}

func (config s3Config) Validate() []error {
	return []error{}
}
