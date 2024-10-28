/*
 * Created on Sun Oct 27 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package file

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/jacklv111/common-sdk/log"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMemFileSys(t *testing.T) {
	log.ValidateAndApply(log.LogConfig)
	fs := NewMemFileSys("")
	file, _ := fs.GetTempFile()
	defer func() {
		fmt.Println("Close file")
		fmt.Println(file.Close())
		fmt.Println("Remove file")
		fmt.Println(fs.Remove(file))
	}()
	// 写入数据到临时文件
	_, _ = file.WriteAt([]byte("Hello, World!\n"), 0)

	// 移动文件指针到开头，以便进行复制
	_, _ = file.Seek(0, io.SeekStart)

	// 创建一个字节缓冲区
	var buf bytes.Buffer

	// 使用 io.Copy 将临时文件的内容复制到字节缓冲区
	_, _ = io.Copy(&buf, file)

	// 打印读取的数据
	fmt.Println("Copied data from temp file:")
	fmt.Println(buf.String())
	Convey("Test MemFile", t, func() {
		So(buf.String(), ShouldEqual, "Hello, World!\n")
	})
}
