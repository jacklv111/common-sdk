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

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// BatchDownloadIterator iterates through files to be downloaded
// to S3.
type BatchDownloadIterator struct {
	writers []WriterMapper
	bucket  string
	next    next
	index   int
	err     error
}

type next struct {
	writer io.WriterAt
	key    string
}

// NewBatchDownloadIterator creates and returns a new BatchDownloadIterator
func NewBatchDownloadIterator(bucket string, writers []WriterMapper) s3manager.BatchDownloadIterator {
	return &BatchDownloadIterator{
		bucket:  bucket,
		writers: writers,
	}
}

// Next opens the next file and stops iteration if it fails to open
// a file.
func (iter *BatchDownloadIterator) Next() bool {
	if len(iter.writers) == iter.index {
		iter.next.writer = nil
		return false
	}

	iter.next.writer = iter.writers[iter.index].Writer
	iter.next.key = iter.writers[iter.index].Key

	iter.index++

	return true && iter.Err() == nil
}

// Err returns an error that was set during opening the file
func (iter *BatchDownloadIterator) Err() error {
	return iter.err
}

// UploadObject returns a BatchUploadObject and sets the After field to
// close the file.
func (iter *BatchDownloadIterator) DownloadObject() s3manager.BatchDownloadObject {
	return s3manager.BatchDownloadObject{
		Object: &s3.GetObjectInput{
			Bucket: &iter.bucket,
			Key:    &iter.next.key,
		},
		Writer: iter.next.writer,
	}
}
