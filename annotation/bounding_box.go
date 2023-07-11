/*
 * Created on Tue Jul 11 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package annotation

type BBox2D struct {
	X      float32 `json:"X"`
	Y      float32 `json:"Y"`
	Width  float32 `json:"Width"`
	Height float32 `json:"Height"`
}

type BBox3D struct {
	X     float32 `json:"X"`
	Y     float32 `json:"Y"`
	Z     float32 `json:"Z"`
	XSize float32 `json:"XSize"`
	YSize float32 `json:"YSize"`
	ZSize float32 `json:"ZSize"`
	YawX  float32 `json:"YawX"`
	YawY  float32 `json:"YawY"`
	YawZ  float32 `json:"YawZ"`
}
