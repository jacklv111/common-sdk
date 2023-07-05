/*
 * Created on Wed Jul 05 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package config

import "github.com/spf13/pflag"

// 配置类的接口，可用于配置的统一管理
type Interface interface {
	// read config from file
	ReadFromFile() error
	// AddFlags is for explicitly initializing the flags.
	AddFlags(*pflag.FlagSet)
	// validate checks the config and return a slice of found errs
	Validate() []error
}
