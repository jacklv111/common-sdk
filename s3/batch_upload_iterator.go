/*
 * Created on Wed Jul 05 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package s3

import (
	"io"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// BatchUploadIterator iterates through files to be uploaded
// to S3.
type BatchUploadIterator struct {
	readerMappers []ReaderMapper
	bucket        string
	next          struct {
		reader io.Reader
		key    string
	}
	index int
	err   error
}

// NewBatchUploadIterator creates and returns a new BatchUploadIterator
func NewBatchUploadIterator(bucket string, readerMappers []ReaderMapper) s3manager.BatchUploadIterator {
	return &BatchUploadIterator{
		bucket:        bucket,
		readerMappers: readerMappers,
	}
}

// Next opens the next file and stops iteration if it fails to open
// a file.
func (iter *BatchUploadIterator) Next() bool {
	if len(iter.readerMappers) == iter.index {
		iter.next.reader = nil
		return false
	}

	iter.next.reader = iter.readerMappers[iter.index].Reader
	iter.next.key = iter.readerMappers[iter.index].Key

	iter.index++
	return true && iter.Err() == nil
}

// Err returns an error that was set during opening the file
func (iter *BatchUploadIterator) Err() error {
	return iter.err
}

// UploadObject returns a BatchUploadObject and sets the After field to
// close the file.
func (iter *BatchUploadIterator) UploadObject() s3manager.BatchUploadObject {
	return s3manager.BatchUploadObject{
		Object: &s3manager.UploadInput{
			Bucket: &iter.bucket,
			Key:    &iter.next.key,
			Body:   iter.next.reader,
		},
	}
}
