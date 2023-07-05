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
	"os"
)

// 本地文件和 s3 上文件的映射关系
type FileMapper struct {
	File *os.File
	Key  string
}

type WriterMapper struct {
	Writer io.WriterAt
	Key    string
}

type ReaderMapper struct {
	Reader io.Reader
	Key    string
}

func GetWriterMapperKeyList(list []WriterMapper) []string {
	res := make([]string, 0)
	for _, mapper := range list {
		res = append(res, mapper.Key)
	}
	return res
}

func GetReaderMapperKeyList(list []ReaderMapper) []string {
	res := make([]string, 0)
	for _, mapper := range list {
		res = append(res, mapper.Key)
	}
	return res
}
