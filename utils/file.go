/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package utils

import (
	"path/filepath"
	"strings"
)

func GetFileNameWithoutSuffix(fileName string) string {
	fileExt := filepath.Ext(fileName)
	nameOnly := strings.TrimSuffix(fileName, fileExt)
	return nameOnly
}
