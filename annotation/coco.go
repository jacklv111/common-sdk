/*
 * Created on Tue Jul 11 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package annotation

import (
	"encoding/json"
	"os"

	"github.com/jacklv111/common-sdk/log"
)

type (
	CocoAnno struct {
		Categories []CocoCategory `json:"categories,omitempty"`
		// key is category id,  value is category name
		CategoryMap map[int]string
		Images      []CocoImage    `json:"images,omitempty"`
		Annotations []CocoAnnoType `json:"annotations,omitempty"`
	}

	CocoImage struct {
		RawDataId    string `json:"raw_data_id,omitempty"`
		Id           int    `json:"id,omitempty"`
		License      int    `json:"license,omitempty"`
		FileName     string `json:"file_name,omitempty"`
		CocoUrl      string `json:"coco_url,omitempty"`
		Height       int    `json:"height,omitempty"`
		Width        int    `json:"width,omitempty"`
		DateCaptured string `json:"date_captured,omitempty"`
		FlickrUrl    string `json:"flick_url,omitempty"`
	}

	CocoAnnoType struct {
		Id int `json:"id,omitempty"`
		// 直接提供 aifs 系统中的 raw data id
		RawDataId string `json:"raw_data_id,omitempty"`
		// 直接提供 aifs 系统中的 label id
		LabelId        string      `json:"label_id,omitempty"`
		Segmentation   interface{} `json:"segmentation,omitempty"`
		IsCrowd        int         `json:"iscrow,omitempty"`
		Area           float32     `json:"area,omitempty"`
		NumKeyPoints   int         `json:"num_keypoints,omitempty"`
		KeyPoints      []int       `json:"keypoints,omitempty"`
		ImageId        int         `json:"image_id,omitempty"`
		Bbox           []float32   `json:"bbox,omitempty"`
		CategoryId     int         `json:"category_id,omitempty"`
		PredictedIou   float32
		PointCoords    [][]float32
		CropBox        []float32
		StabilityScore float32
	}

	CocoCategory struct {
		Name          string    `json:"name,omitempty"`
		SuperCategory string    `json:"supercategory,omitempty"`
		Id            int       `json:"id,omitempty"`
		KeyPoints     []string  `json:"keypoints,omitempty"`
		Skeleton      [][]int32 `json:"skeleton,omitempty"`
	}

	// CocoAnnoFormat is the format of annotation data at server side
	CocoAnnoFormat struct {
		RawDataId            string         `json:"RawDataId,omitempty"`
		AnnotationTemplateId string         `json:"AnnotationTemplateId,omitempty"`
		AnnoData             []CocoAnnoData `json:"AnnoData,omitempty"`
	}
	CocoAnnoData struct {
		Id             int         `json:"Id,omitempty"`
		LabelId        string      `json:"LabelId,omitempty"`
		Segmentation   interface{} `json:"Segmentation,omitempty"`
		IsCrowd        int         `json:"IsCrowd,omitempty"`
		Area           float32     `json:"Area,omitempty"`
		NumKeyPoints   int         `json:"NumKeyPoints,omitempty"`
		KeyPoints      []int       `json:"KeyPoints,omitempty"`
		Bbox           []float32   `json:"Bbox,omitempty"`
		PredictedIou   float32     `json:"PredictedIou,omitempty"`
		PointCoords    [][]float32 `json:"PointCoords,omitempty"`
		CropBox        []float32   `json:"CropBox,omitempty"`
		StabilityScore float32     `json:"StabilityScore,omitempty"`
	}
)

func (anno *CocoAnno) UnmarshalJSON(data []byte) error {
	type Alias CocoAnno
	wrapper := &struct {
		RawDataId string
		*Alias
	}{
		Alias: (*Alias)(anno),
	}

	if err := json.Unmarshal(data, wrapper); err != nil {
		return err
	}
	anno.CategoryMap = make(map[int]string)
	for _, category := range anno.Categories {
		anno.CategoryMap[category.Id] = category.Name
	}
	return nil
}

func (anno *CocoAnno) ParseAnnotationFile(annoFilePath string) error {
	log.Infof("parse annotation file %s", annoFilePath)
	content, err := os.ReadFile(annoFilePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &anno)
	if err != nil {
		return err
	}
	return nil
}
