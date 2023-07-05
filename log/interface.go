/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package log

import (
	"os"
	"strings"
	"time"

	"github.com/jacklv111/common-sdk/env"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ENCODING_JSON    = "json"
	ENCODING_CONSOLE = "console"
)

// 使用 sugaredLogger 性能上比不上 zap.Logger，但是差距不大。而且 key value 类型的输出没有类型保护。
// 主要是接口上的考虑选择了 sugaredLogger
// 上层只需要知道 log 工具能提供什么功能，知道几组简单的接口
var logger zap.SugaredLogger

// init interface
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// ValidateAndApply combines validation and application of the logging configuration.
// This should be invoked as early as possible because then the rest of the program
// startup (including validation of other options) will already run with the final
// logging configuration.
func ValidateAndApply(logcfg *logConfig) []error {
	var errs = logcfg.Validate()
	if len(errs) > 0 {
		return errs
	}

	envType := env.EnvConfig.GetEnvType()

	// Output should also go to standard out.
	consoleDebugging := zapcore.Lock(os.Stdout)

	var encoderConfig zapcore.EncoderConfig
	var fileEncoder, consoleEncoder zapcore.Encoder
	if strings.EqualFold(envType, env.PROD) {
		encoderConfig = zap.NewProductionEncoderConfig()
	} else if strings.EqualFold(envType, env.DEV) {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	} else if strings.EqualFold(envType, env.TEST) {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	}

	encoderConfig.EncodeTime = timeEncoder

	if strings.EqualFold(logcfg.Encoding, "json") {
		fileEncoder = zapcore.NewJSONEncoder(encoderConfig)
		consoleEncoder = zapcore.NewJSONEncoder(encoderConfig)
	} else if strings.EqualFold(logcfg.Encoding, "console") {
		fileEncoder = zapcore.NewConsoleEncoder(encoderConfig)
		if logcfg.Color {
			encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		}
		consoleEncoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	var cores []zapcore.Core
	if !logcfg.Stdout && !logcfg.FilesOut || logcfg.Stdout {
		cores = append(cores, zapcore.NewCore(consoleEncoder, consoleDebugging, getLevel(logcfg.Level)))
	}

	if logcfg.FilesOut && len(logcfg.LogPath) > 0 {
		for i := 0; i < len(logcfg.LogPath); i++ {
			cores = append(cores, zapcore.NewCore(fileEncoder, zapcore.AddSync(logcfg.LogPath[i].Hook), getLevel(logcfg.LogPath[i].Level)))
		}
	}

	core := zapcore.NewTee(cores...)
	// From a zapcore.Core to construct a Logger.

	if logcfg.AddCaller {
		logger = *zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
	} else {
		logger = *zap.New(core).Sugar()
	}

	return errs
}

func getLevel(levelStr string) *zapcore.Level {
	level := new(zapcore.Level)
	_ = level.UnmarshalText([]byte(levelStr))
	return level
}

//------------------------------------------------------ log interface

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	logger.Error(args...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanic(args ...interface{}) {
	logger.DPanic(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicf(args ...interface{}) {
	logger.DPanic(args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	s.With(keysAndValues).Debug(msg)
func Debugw(msg string, keysAndValues ...interface{}) {
	logger.Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...interface{}) {
	logger.Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warnw(msg string, keysAndValues ...interface{}) {
	logger.Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...interface{}) {
	logger.Errorf(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func DPanicw(msg string, keysAndValues ...interface{}) {
	logger.DPanicw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func Panicw(msg string, keysAndValues ...interface{}) {
	logger.Panicw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues ...interface{}) {
	logger.Fatalw(msg, keysAndValues...)
}
