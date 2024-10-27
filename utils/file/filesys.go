/*
 * Created on Sun Oct 27 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */
package file

type FileSys interface {
	GetTempFile() (File, error)
	Remove(file File) error
}