/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package utils

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"

	"github.com/jacklv111/common-sdk/log"
	valueobject "github.com/jacklv111/common-sdk/utils/value-object"
)

// GetImageMeta 获取图片的 metadata
//
//	@param filePath
//	@return valueobject.ImageMeta
//	@return error
func GetImageMeta(filePath string) (valueobject.ImageMeta, error) {
	var meta valueobject.ImageMeta

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return meta, err
	}
	meta.Size = fileInfo.Size()

	file, err := os.Open(filePath)
	if err != nil {
		return meta, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.DPanicf("close file error: %s", err)
		}
	}()

	imgCfg, _, err := image.DecodeConfig(file)
	if err != nil {
		log.Errorf("%s decode config err [%s]", filePath, err)
	} else {
		meta.Width = int32(imgCfg.Width)
		meta.Height = int32(imgCfg.Height)
	}

	return meta, nil
}

func ReadImage(reader io.Reader) (image.Image, error) {
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func IsImageFromFile(filePath string) bool {
	f, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer f.Close()

	_, _, err = image.DecodeConfig(f)
	return err == nil
}
