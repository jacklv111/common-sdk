/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package utils

import (
	"reflect"
	"testing"

	testdata "github.com/jacklv111/common-sdk/utils/test-data"
	. "github.com/smartystreets/goconvey/convey"
)

func TestReadAllFiles(t *testing.T) {
	Convey("read all files", t, func() {
		Convey("Success", func() {
			result, _ := ReadAllFiles(testdata.TEST_PATH)
			So(result, ShouldResemble, []string{"test-data/os-utils/a.jpg", "test-data/os-utils/b.jpg", "test-data/os-utils/c.jpg", "test-data/os-utils/exclude/d.jpg"})
		})
	})
}

func TestGetFileSha256(t *testing.T) {
	Convey("cal image sha256", t, func() {
		res, err := GetFileSha256FromFile(testdata.TEST_IMAGE_JPG)
		So(err, ShouldEqual, nil)
		So(res, ShouldEqual, "39728b1b7ab36d4540a6b16ad2cba26d28e7e60e81ffbffb13899f3f8606102f")
	})
}

func TestReadAllFilesExclude(t *testing.T) {
	type args struct {
		path    string
		exclude []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "TestReadAllFilesExclude1", args: args{path: testdata.TEST_PATH, exclude: []string{"a.jpg"}}, want: []string{"test-data/os-utils/b.jpg", "test-data/os-utils/c.jpg", "test-data/os-utils/exclude/d.jpg"}, wantErr: false},
		{name: "TestReadAllFilesExclude2", args: args{path: testdata.TEST_PATH, exclude: []string{"a.jpg", "b.jpg"}}, want: []string{"test-data/os-utils/c.jpg", "test-data/os-utils/exclude/d.jpg"}, wantErr: false},
		{name: "TestReadAllFilesExclude3", args: args{path: testdata.TEST_PATH, exclude: []string{"exclude"}}, want: []string{"test-data/os-utils/a.jpg", "test-data/os-utils/b.jpg", "test-data/os-utils/c.jpg"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadAllFilesExclude(tt.args.path, tt.args.exclude)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAllFilesExclude() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadAllFilesExclude() = %v, want %v", got, tt.want)
			}
		})
	}
}
