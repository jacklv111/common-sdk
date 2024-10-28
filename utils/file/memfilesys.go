/*
 * Created on Sun Oct 27 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */
package file

import "github.com/spf13/afero"

type MemFileSys struct {
	memFs   afero.Fs
	tempDir string
}

func NewMemFileSys(tempDir string) *MemFileSys {
	return &MemFileSys{
		memFs:   afero.NewMemMapFs(),
		tempDir: tempDir,
	}
}

func (m *MemFileSys) GetTempFile() (File, error) {
	return afero.TempFile(m.memFs, m.tempDir, TEMP_FILE_PATTERN)
}

func (m *MemFileSys) Remove(file File) error {
	return m.memFs.Remove(file.Name())
}
