/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package utils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

func GenerateSecret() (string, error) {
	key := make([]byte, 32) // AES-256
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}

// generateKey 生成指定长度的随机字符串
func generateKey(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(bytes)[:length], nil
}

// generateAKSK 生成 Access Key 和 Secret Key
func GenerateAKSK() (string, string, error) {
	ak, err := generateKey(16) // 生成 16 字节的 Access Key
	if err != nil {
		return "", "", err
	}
	sk, err := generateKey(32) // 生成 32 字节的 Secret Key
	if err != nil {
		return "", "", err
	}
	return ak, sk, nil
}
