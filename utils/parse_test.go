/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseInt(t *testing.T) {
	type args struct {
		input      string
		minVal     int
		maxVal     int
		defaultVal int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
		wantErr    bool
	}{
		{"string is empty", args{"", 0, 50, 10}, 10, false},
		{"string is invalid", args{"tt", 0, 50, 10}, 10, true},
		{"string is valid but number smaller than minVal", args{"-10", 0, 50, 10}, -10, true},
		{"string is valid but number bigger than maxVal", args{"100", 0, 50, 10}, 100, true},
		{"string is valid and number is in the range", args{"20", 0, 50, 10}, 20, false},
	}
	for _, tt := range tests {
		Convey(tt.name, t, func() {
			num, err := ParseInt(tt.args.input, tt.args.minVal, tt.args.maxVal, tt.args.defaultVal)
			if tt.wantErr {
				So(err, ShouldNotBeNil)
			} else {
				So(err, ShouldBeNil)
				if tt.args.input == "" {
					So(num, ShouldEqual, tt.args.defaultVal)
				} else {
					So(num, ShouldEqual, tt.wantResult)
				}
			}
		})
	}
}
