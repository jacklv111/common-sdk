/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package env

import (
	"fmt"

	"github.com/spf13/pflag"
	"golang.org/x/exp/slices"
)

const (
	DEV  = "development"
	PROD = "production"
	TEST = "test"
)

var envTypeList = []string{DEV, PROD, TEST}

type environmentConfig struct {
	envType string
}

var EnvConfig *environmentConfig

// init sets up the defaults.
func init() {
	EnvConfig = &environmentConfig{envType: DEV}
}

func (env *environmentConfig) AddFlags(flagset *pflag.FlagSet) {
	flagset.StringVar(&env.envType, "env-type", DEV, "Value to indicate the environment, e.g.: development, production, test.")
}

// do nothing
func (env *environmentConfig) ReadFromFile() error {
	return nil
}

func (env environmentConfig) Validate() []error {
	var errs []error

	if !slices.Contains(envTypeList, env.envType) {
		errs = append(errs, fmt.Errorf("env type should be in %q, the current value is %s", envTypeList, env.envType))
	}

	return errs
}

// envType getter
func (env environmentConfig) GetEnvType() string {
	return env.envType
}
