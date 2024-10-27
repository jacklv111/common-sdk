/*
 * Created on Sun Oct 27 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */
package file

import "io"

type File interface {
	io.WriterAt
	io.Reader
	io.Closer
	io.Seeker
	Name() string
}
