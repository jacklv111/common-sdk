/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package utils

import (
	"bytes"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"testing"

	testdata "github.com/jacklv111/common-sdk/utils/test-data"
	. "github.com/smartystreets/goconvey/convey"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
)

func TestGetImageMeta(t *testing.T) {
	Convey("get image meta", t, func() {
		Convey("jpg Success", func() {
			meta, err := GetImageMeta(testdata.TEST_IMAGE_JPG)
			So(err, ShouldEqual, nil)
			So(meta.Size, ShouldEqual, 11890)
			So(meta.Height, ShouldEqual, 265)
			So(meta.Width, ShouldEqual, 265)
		})
		Convey("png Success", func() {
			meta, err := GetImageMeta(testdata.TEST_IMAGE_PNG)
			So(err, ShouldEqual, nil)
			So(meta.Size, ShouldEqual, 9466)
			So(meta.Height, ShouldEqual, 225)
			So(meta.Width, ShouldEqual, 225)
		})
		Convey("gif Success", func() {
			meta, err := GetImageMeta(testdata.TEST_IMAGE_GIF)
			So(err, ShouldEqual, nil)
			So(meta.Size, ShouldEqual, 3012729)
			So(meta.Height, ShouldEqual, 360)
			So(meta.Width, ShouldEqual, 480)
		})
		Convey("bmp Success", func() {
			meta, err := GetImageMeta(testdata.TEST_IMAGE_BMP)
			So(err, ShouldEqual, nil)
			So(meta.Size, ShouldEqual, 818058)
			So(meta.Height, ShouldEqual, 426)
			So(meta.Width, ShouldEqual, 640)
		})
		Convey("tiff Success", func() {
			meta, err := GetImageMeta(testdata.TEST_IMAGE_TIFF)
			So(err, ShouldEqual, nil)
			So(meta.Size, ShouldEqual, 1131930)
			So(meta.Height, ShouldEqual, 434)
			So(meta.Width, ShouldEqual, 650)
		})
	})
}

func TestReadImage(t *testing.T) {
	Convey("get image meta", t, func() {
		Convey("jpg Success", func() {
			file, _ := os.Open(testdata.TEST_IMAGE_JPG)
			byteData, _ := io.ReadAll(file)
			image, err := ReadImage(bytes.NewReader(byteData))
			So(err, ShouldEqual, nil)
			So(len(byteData), ShouldEqual, 67294)
			So(image.Bounds().Dy(), ShouldEqual, 720)
			So(image.Bounds().Dx(), ShouldEqual, 1280)
		})
		Convey("png Success", func() {
			file, _ := os.Open(testdata.TEST_IMAGE_PNG)
			byteData, _ := io.ReadAll(file)
			image, err := ReadImage(bytes.NewReader(byteData))
			So(err, ShouldEqual, nil)
			So(len(byteData), ShouldEqual, 9466)
			So(image.Bounds().Dy(), ShouldEqual, 225)
			So(image.Bounds().Dx(), ShouldEqual, 225)
		})
		Convey("gif Success", func() {
			file, _ := os.Open(testdata.TEST_IMAGE_GIF)
			byteData, _ := io.ReadAll(file)
			image, err := ReadImage(bytes.NewReader(byteData))
			So(err, ShouldEqual, nil)
			So(len(byteData), ShouldEqual, 3012729)
			So(image.Bounds().Dy(), ShouldEqual, 360)
			So(image.Bounds().Dx(), ShouldEqual, 480)
		})
		Convey("bmp Success", func() {
			file, _ := os.Open(testdata.TEST_IMAGE_BMP)
			byteData, _ := io.ReadAll(file)
			image, err := ReadImage(bytes.NewReader(byteData))
			So(err, ShouldEqual, nil)
			So(len(byteData), ShouldEqual, 818058)
			So(image.Bounds().Dy(), ShouldEqual, 426)
			So(image.Bounds().Dx(), ShouldEqual, 640)
		})
		Convey("tiff Success", func() {
			file, _ := os.Open(testdata.TEST_IMAGE_TIFF)
			byteData, _ := io.ReadAll(file)
			image, err := ReadImage(bytes.NewReader(byteData))
			So(err, ShouldEqual, nil)
			So(len(byteData), ShouldEqual, 1131930)
			So(image.Bounds().Dy(), ShouldEqual, 434)
			So(image.Bounds().Dx(), ShouldEqual, 650)
		})
	})
}

func TestIsImageFromFile(t *testing.T) {
	Convey("get image meta", t, func() {
		Convey("jpg Success", func() {
			isImage := IsImageFromFile(testdata.TEST_IMAGE_JPG)
			So(isImage, ShouldEqual, true)
		})
		Convey("png Success", func() {
			isImage := IsImageFromFile(testdata.TEST_IMAGE_PNG)
			So(isImage, ShouldEqual, true)
		})
		Convey("gif Success", func() {
			isImage := IsImageFromFile(testdata.TEST_IMAGE_GIF)
			So(isImage, ShouldEqual, true)
		})
		Convey("bmp Success", func() {
			isImage := IsImageFromFile(testdata.TEST_IMAGE_BMP)
			So(isImage, ShouldEqual, true)
		})
		Convey("tiff Success", func() {
			isImage := IsImageFromFile(testdata.TEST_IMAGE_TIFF)
			So(isImage, ShouldEqual, true)
		})
	})
}
