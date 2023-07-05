/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package utils

import "testing"

func TestCamelToSnake(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestCamelToSnake1", args: args{s: "AnnotationTemplate"}, want: "annotation_template"},
		{name: "TestCamelToSnake2", args: args{s: "BoundingBoxForImage"}, want: "bounding_box_for_image"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelToSnake(tt.args.s); got != tt.want {
				t.Errorf("CamelToSnake() = %v, want %v", got, tt.want)
			}
		})
	}
}
