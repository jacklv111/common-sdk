/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net"
	"os"
	"path/filepath"

	"github.com/jacklv111/common-sdk/collection/mapset"
	"github.com/jacklv111/common-sdk/log"
)

// 读取一个目录下的所有文件，包括文件夹中的文件
//
//	@param path 目录的路径
//	@return []string 目录下的文件路径列表
func ReadAllFiles(path string) ([]string, error) {
	var result []string
	fileInfoList, err := os.ReadDir(path)
	if err != nil {
		return result, err
	}

	for _, file := range fileInfoList {
		if file.IsDir() {
			subRes, err := ReadAllFiles(filepath.Join(path, file.Name()))
			if err != nil {
				return subRes, err
			}
			result = append(result, subRes...)
		} else {
			result = append(result, filepath.Join(path, file.Name()))
		}
	}

	return result, nil
}

// 读取一个目录下的所有文件，包括文件夹中的文件
//
//	@param path 目录的路径
//	@return []string 目录下的文件路径列表
func ReadAllFilesExclude(path string, exclude []string) ([]string, error) {
	excludeSet := mapset.NewSet(exclude...)

	var result []string
	fileInfoList, err := os.ReadDir(path)
	if err != nil {
		return result, err
	}

	for _, file := range fileInfoList {
		if excludeSet.Contains(file.Name()) {
			continue
		}
		if file.IsDir() {
			subRes, err := ReadAllFiles(filepath.Join(path, file.Name()))
			if err != nil {
				return subRes, err
			}
			result = append(result, subRes...)
		} else {
			result = append(result, filepath.Join(path, file.Name()))
		}
	}

	return result, nil
}

// ReadFilesReturnMap 扫描文件夹中的文件，得到所有的文件，以 map 的形式返回
//
//	@param path
//	@return map[string]string key: 文件名，不带后缀; value: 文件路径
//	@return error
func ReadFilesReturnMap(dir string) (map[string]string, error) {
	result := make(map[string]string, 0)
	fileInfoList, err := os.ReadDir(dir)
	if err != nil {
		return result, err
	}
	for _, file := range fileInfoList {
		if file.IsDir() {
			continue
		}
		nameOnly := GetFileNameWithoutSuffix(file.Name())
		result[nameOnly] = filepath.Join(dir, file.Name())
	}

	return result, nil
}

// IsDir 判断文件是否为文件夹
//
//	@param path
//	@return bool
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// PathExists 判断 path 是否存在
//
//	@param path
//	@return bool
//	@return error
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetFileSha256FromFile 获取文件的 sha256 值
//
//	@param filePath
//	@return string
//	@return error
func GetFileSha256FromFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		return "", err
	}
	hash := sha256.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// GetFileSha256Bytes 获取 bytes 的 sha256 值
//
//	@param buf []byte
//	@return string
func GetFileSha256Bytes(buf []byte) (string, error) {

	hash := sha256.New()
	_, err := io.Copy(hash, bytes.NewReader(buf))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func GetHostIp() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		log.Errorf("net get interface address error [%s]", err.Error())
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}
	log.Errorf("can not find the client ip address")

	return ""
}

func CloseFiles(files []io.Closer) {
	// close file
	for _, file := range files {
		if err := file.Close(); err != nil {
			log.DPanicf("close file error: %s", err)
		}
	}
}
