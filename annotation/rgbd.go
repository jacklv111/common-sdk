/*
 * Created on Tue Jul 11 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package annotation

// rgbd
type RgbdAnnotation struct {
	RawDataId            string            `json:"RawDataId"`
	AnnotationTemplateId string            `json:"AnnotationTemplateId"`
	BoundingBoxList      []RgbdBoundingBox `json:"BoundingBoxList"`
}

type RgbdBoundingBox struct {
	LabelId       string `json:"LabelId"`
	BoundingBox2D BBox2D `json:"BoundingBox2D"`
	BoundingBox3D BBox3D `json:"BoundingBox3D"`
}
