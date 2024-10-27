/*
 * Created on Sun Oct 27 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */
package file

var Fs FileSys

func Init() {
	switch Config.FileSys {
	case "mem":
		Fs = NewMemFileSys(Config.TempDir)
	case "disk":
		Fs = NewDiskFileSys(Config.TempDir)
	default:
		Fs = NewMemFileSys(Config.TempDir)
	}
}
