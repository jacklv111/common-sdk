/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package gorm

import (
	"log"
	"os"
	"time"

	"github.com/spf13/pflag"
	gormlib "gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type gormConfig struct {
	gormlib.Config
}

var GormConfig *gormConfig

func init() {
	GormConfig = &gormConfig{}
	GormConfig.CreateBatchSize = 50

	GormConfig.Logger = logger.New(
		//将标准输出作为Writer
		log.New(os.Stdout, "\r\n", log.LstdFlags),

		logger.Config{
			// 设定慢查询时间阈值为1ms
			SlowThreshold: 1000 * time.Microsecond,
			// 设置日志级别
			LogLevel: logger.Warn,
		},
	).LogMode(logger.Info)
}

func (config *gormConfig) ReadFromFile() error {
	// do nothing
	return nil
}

func (config *gormConfig) AddFlags(flagSet *pflag.FlagSet) {
	flagSet.BoolVar(&config.DryRun, "orm-dry-run", false, "if true: generate sql without execute")
}

func (config gormConfig) Validate() []error {
	// do nothing
	var errs []error
	return errs
}
