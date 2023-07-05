/*
 * Created on Wed Jul 05 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package s3

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// BatchUploadIterator iterates through files to be uploaded
// to S3.
type BatchDeleteIterator struct {
	keys   []string
	bucket string
	next   struct {
		key string
	}
	index int
	err   error
}

// NewBatchDeleteIterator creates and returns a new NewBatchDeleteIterator
func NewBatchDeleteIterator(bucket string, keys []string) s3manager.BatchDeleteIterator {
	return &BatchDeleteIterator{
		bucket: bucket,
		keys:   keys,
	}
}

// Next opens the next file and stops iteration if it fails to open
// a file.
func (iter *BatchDeleteIterator) Next() bool {
	if len(iter.keys) == iter.index {
		return false
	}

	iter.next.key = iter.keys[iter.index]

	iter.index++
	return true && iter.Err() == nil
}

func (iter *BatchDeleteIterator) Err() error {
	return iter.err
}

// DeleteObject returns a BatchDeleteObject
func (iter *BatchDeleteIterator) DeleteObject() s3manager.BatchDeleteObject {
	return s3manager.BatchDeleteObject{
		Object: &s3.DeleteObjectInput{
			Bucket: &iter.bucket,
			Key:    &iter.next.key,
		},
	}
}
