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
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

func DownloadZipCmd(dataViewId string, localDir string) (filePath string, err error) {
	args := []string{
		"download",
		fmt.Sprintf("--data_view_id=%s", dataViewId),
		fmt.Sprintf("--local_dir=%s", localDir),
	}
	cmd := exec.Command("aifs", args...)
	var stderr bytes.Buffer
	var stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	err = cmd.Run()

	if err != nil {
		return "", fmt.Errorf(stderr.String())
	}
	return strings.Trim(stdout.String(), " \t\n\r"), nil
}

func GetAifsConfigFromDataClient() (aifsUrl string, err error) {
	args := []string{
		"showconfig",
	}
	cmd := exec.Command("aifs", args...)
	var stderr bytes.Buffer
	var stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	err = cmd.Run()

	if err != nil {
		return "", fmt.Errorf(stderr.String())
	}
	dict := strings.ReplaceAll(strings.Trim(stdout.String(), " \t\n\r"), "'", "\"")
	type (
		aifsConfig struct {
			Ip   string `json:"ip"`
			Port string `json:"port"`
			Host string `json:"host"`
		}
		dataClientConfig struct {
			Aifs aifsConfig `json:"aifs"`
		}
	)
	var config dataClientConfig
	err = json.Unmarshal([]byte(dict), &config)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("http://%s", config.Aifs.Host), err
}
