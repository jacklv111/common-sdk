/*
 * Created on Sun Oct 27 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */
package file

import "os"

type DiskFileSys struct {
	tempDir string
}

func NewDiskFileSys(tempDir string) *DiskFileSys {
	return &DiskFileSys{
		tempDir: tempDir,
	}
}

func (d *DiskFileSys) GetTempFile() (File, error) {
	return os.CreateTemp(d.tempDir, TEMP_FILE_PATTERN)
}

func (d *DiskFileSys) Remove(file File) error {
	return os.Remove(file.Name())
}
