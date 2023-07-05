/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jacklv111/common-sdk/log"
	testdata "github.com/jacklv111/common-sdk/utils/test-data"
)

func TestDecompress(t *testing.T) {
	log.ValidateAndApply(log.LogConfig)
	tempDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	type args struct {
		srcFilePath  string
		destDir      string
		checkZipBomb bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "test.tar decompress success", args: args{srcFilePath: filepath.Join(testdata.TEST_FILE_DIR, "test.tar"), destDir: tempDir, checkZipBomb: false}, wantErr: false},
		{name: "test.zip decompress success", args: args{srcFilePath: filepath.Join(testdata.TEST_FILE_DIR, "test.zip"), destDir: tempDir, checkZipBomb: false}, wantErr: false},
		{name: "test decompress success", args: args{srcFilePath: filepath.Join(testdata.TEST_FILE_DIR, "test"), destDir: tempDir, checkZipBomb: false}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Decompress(tt.args.srcFilePath, tt.args.destDir, tt.args.checkZipBomb); (err != nil) != tt.wantErr {
				t.Errorf("Decompress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
