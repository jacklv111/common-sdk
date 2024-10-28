/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package log

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type (
	logConfig struct {
		Level     string         `json:"level"`
		Stdout    bool           `json:"stdout"`
		Encoding  string         `json:"encoding"`
		AddCaller bool           `json:"addcaller"`
		Color     bool           `json:"color"`
		FilesOut  bool           `json:"filesout"`
		LogPath   []*logFilePath `json:"logpath"`
	}

	logFilePath struct {
		Level string             `json:"level"`
		Hook  *lumberjack.Logger `json:"hook"`
	}
)

const (
	LOG_CONFIG_PATH = "./conf/logs.json"
)

var LogConfig *logConfig

func init() {
	LogConfig = &logConfig{
		Level:     "debug",
		Stdout:    true,
		Encoding:  "console",
		AddCaller: true,
		Color:     true,
		FilesOut:  false,
	}
}

func (logcfg *logConfig) ReadFromFile() error {
	content, err := os.ReadFile(LOG_CONFIG_PATH)
	if err != nil {
		fmt.Fprintf(os.Stderr, "logger: there is no any config file, using default config, error: %v\n", err)
	} else {
		if err := json.Unmarshal(content, &logcfg); err != nil {
			return err
		}
	}
	return nil
}

func (logcfg *logConfig) AddFlags(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&logcfg.Level, "log-level", "info", "Value to indicate the lowest level log to be output, e.g.: debug, info, warn, error, dpanic, panic, fatal.")
	flagSet.BoolVar(&logcfg.Stdout, "log-stdout", true, "Value to indicate whether to output log to stdout.")
	flagSet.StringVar(&logcfg.Encoding, "log-encoding", "console", "Value to indicate the encoding of log, e.g.: console, json.")
	flagSet.BoolVar(&logcfg.AddCaller, "log-add-caller", true, "Value to indicate whether to add caller info to log.")
	flagSet.BoolVar(&logcfg.Color, "log-color", true, "Value to indicate whether to add color to log.")
	flagSet.BoolVar(&logcfg.FilesOut, "log-files-out", false, "Value to indicate whether to output log to files.")
}

func (logcfg logConfig) Validate() []error {
	var errs []error

	errs = append(errs, validateAllLogLevel(logcfg)...)
	if err := validateEncoding(logcfg.Encoding); err != nil {
		errs = append(errs, err)
	}

	return errs
}

func validateAllLogLevel(logcfg logConfig) []error {
	var errs []error

	logLevels := []string{logcfg.Level}
	for _, val := range logcfg.LogPath {
		logLevels = append(logLevels, val.Level)
	}
	for _, val := range logLevels {
		if err := validateLogLevel(val); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func validateLogLevel(levelStr string) error {
	level := new(zapcore.Level)
	return level.UnmarshalText([]byte(levelStr))
}

func validateEncoding(encoding string) error {
	if !strings.EqualFold(encoding, ENCODING_JSON) && !strings.EqualFold(encoding, ENCODING_CONSOLE) {
		return fmt.Errorf("encoding [%s] is invalid, it could be [%s] or [%s]", encoding, ENCODING_JSON, ENCODING_CONSOLE)
	}
	return nil
}
