/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package utils

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/jacklv111/common-sdk/log"
)

const (
	MAX_FILE_SIZE               = 100 * 1024 * 1024 // 100 MB
	MAX_TOTAL_UNCOMPRESSED_SIZE = 100 * 1024 * 1024 // 100 MB
	MAX_FILES_IN_ARCHIVE        = 1000
)

type Decompression interface {
	Decompress(srcFilePath string, destDir string, checkZipBomb bool) error
}

type ZipDecompression struct{}

func (ZipDecompression) Decompress(srcFilePath string, destDir string, checkZipBomb bool) error {
	// Open the zip archive for reading
	r, err := zip.OpenReader(srcFilePath)
	if err != nil {
		return err
	}
	defer r.Close()

	if checkZipBomb {
		// Check the total uncompressed size
		var totalSize uint64
		for _, f := range r.File {
			totalSize += f.UncompressedSize64
			if f.UncompressedSize64 > MAX_FILE_SIZE {
				return errors.New("file size exceeds the limit")
			}
			if totalSize > MAX_TOTAL_UNCOMPRESSED_SIZE {
				return errors.New("uncompressed archive size exceeds the limit")
			}
		}

		// Check the number of files in the archive
		if len(r.File) > MAX_FILES_IN_ARCHIVE {
			return errors.New("number of files in the archive exceeds the limit")
		}
	}

	// Iterate through the files in the archive
	for _, file := range r.File {
		// Create the target file path
		fpath := filepath.Join(destDir, file.Name)

		// Validate the file path to protect against zip slip
		if !strings.HasPrefix(fpath, filepath.Clean(destDir)+string(os.PathSeparator)) {
			return errors.New("illegal file path detected")
		}

		// Check for directory
		if file.FileInfo().IsDir() {
			// Create the directory if it doesn't exist
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			// Create the parent directory if it doesn't exist
			if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			copy(file, fpath)
		}
	}

	return nil
}

func copy(src *zip.File, destPath string) error {
	// Create and open the target file for writing
	dstFile, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, src.Mode())
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Open the source file for reading
	srcFile, err := src.Open()
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Copy the contents from the source file to the target file
	if _, err = io.Copy(dstFile, srcFile); err != nil {
		return err
	}
	return nil
}

// ---------------------------------------------------------------------------------------

type TarDecompression struct{}

func (TarDecompression) Decompress(srcFilePath, destDir string, checkZipBomb bool) error {
	log.Info("ExtractTarFile srcFilePath: %s, destDir: %s", srcFilePath, destDir)
	file, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return decompressTar(file, destDir, checkZipBomb)
}

type Tar2Decompression struct{}

func (Tar2Decompression) Decompress(srcFilePath, destDir string, checkZipBomb bool) error {
	log.Info("ExtractTarGzFile srcFilePath: %s, destDir: %s", srcFilePath, destDir)
	file, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	// 将打开的文件先解压
	gr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	return decompressTar(gr, destDir, checkZipBomb)
}

func decompressTar(reader io.Reader, destDir string, checkZipBomb bool) error {
	tarReader := tar.NewReader(reader)
	var totalSize int64
	fileCount := 0

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break // End of archive
		}

		if err != nil {
			return err
		}

		// Check file count
		fileCount++
		if checkZipBomb {
			if fileCount > MAX_FILES_IN_ARCHIVE {
				return errors.New("too many files in tar archive")
			}

			// Check file size
			if header.Size > MAX_FILE_SIZE {
				return fmt.Errorf("file %s is too large", header.Name)
			}

			// Check total size
			totalSize += header.Size
			if totalSize > MAX_TOTAL_UNCOMPRESSED_SIZE {
				return errors.New("tar archive size exceeds the limit")
			}
		}

		target := filepath.Join(destDir, header.Name)

		// Defend against Zip Slip vulnerability
		cleanTarget, err := filepath.Abs(target)
		if err != nil {
			return err
		}

		cleanDest, err := filepath.Abs(destDir)
		if err != nil {
			return err
		}

		if !strings.HasPrefix(cleanTarget, cleanDest) {
			return fmt.Errorf("invalid file path: %s", header.Name)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(cleanTarget, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			// Ensure the destination directory exists
			if err := os.MkdirAll(filepath.Dir(cleanTarget), 0755); err != nil {
				return err
			}

			err = func(filePath string, src io.Reader, size int64) error {
				// Create the file
				outFile, err := os.Create(filePath)
				if err != nil {
					return err
				}
				defer outFile.Close()

				if _, err := io.CopyN(outFile, src, size); err != nil {
					return err
				}
				return nil
			}(cleanTarget, tarReader, header.Size)
			if err != nil {
				return err
			}
		default:
			fmt.Printf("Unsupported type: %c for file: %s\n", header.Typeflag, header.Name)
		}
	}
	return nil
}

// ---------------------------------------------------------------------------------------

var decompressionList []Decompression

func init() {
	decompressionList = []Decompression{ZipDecompression{}, TarDecompression{}, Tar2Decompression{}}
}

func Decompress(srcFilePath string, destDir string, checkZipBomb bool) error {
	for _, decompression := range decompressionList {
		if err := decompression.Decompress(srcFilePath, destDir, checkZipBomb); err != nil {
			log.Info("Decompress failed: %s", err)
			continue
		}
		return nil
	}
	return errors.New("decompression failed")
}
