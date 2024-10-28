/*
 * Created on Sun Oct 27 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */
package file

import (
	"errors"
	"strings"

	"github.com/spf13/pflag"
)

type fileSysConfig struct {
	TempDir string
	FileSys string
}

var Config *fileSysConfig
var fileSysSet map[string]struct{}

func init() {
	fileSysSet = make(map[string]struct{}, 2)
	fileSysSet["disk"] = struct{}{}
	fileSysSet["mem"] = struct{}{}

	Config = &fileSysConfig{
		TempDir: "",
		FileSys: "mem",
	}
}

func (config *fileSysConfig) ReadFromFile() error {
	return nil
}

func (config *fileSysConfig) AddFlags(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&config.TempDir, "temp-dir", "", "temp dir to put temporary files")
	flagSet.StringVar(&config.FileSys, "file-sys", "mem", "file system to use, disk or mem")
}

func (config fileSysConfig) Validate() []error {
	_, ok := fileSysSet[strings.ToLower(config.FileSys)]
	if !ok {
		return []error{errors.New("invalid file system")}
	}
	return []error{}
}
