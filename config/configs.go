/*
 * Created on Wed Jul 05 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package config

import (
	"github.com/spf13/pflag"
)

type Configs struct {
	configList []Interface
}

func (cfg *Configs) AddConfig(config Interface) {
	cfg.configList = append(cfg.configList, config)
}

func (cfg Configs) GetFlags() (flagSet *pflag.FlagSet) {
	flagSet = &pflag.FlagSet{}
	for _, config := range cfg.configList {
		config.AddFlags(flagSet)
	}
	return
}

func (cfg Configs) Validate() []error {
	var errs []error
	for _, config := range cfg.configList {
		errs = append(errs, config.Validate()...)
	}
	return errs
}

func (cfg *Configs) ReadFromFile() []error {
	var errs []error
	for _, config := range cfg.configList {
		if err := config.ReadFromFile(); err != nil {
			errs = append(errs, config.ReadFromFile())
		}
	}
	return errs
}
