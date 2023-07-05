/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package database

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

const (
	DB_CONFIG_PATH = "./conf/dbconfig.yml"
)

type dbConfig struct {
	Name     string `yaml:"name"`
	Ip       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`

	// the maximum number of connections in the idle connection pool
	MaxIdleConns int
	// the maximum number of open connections to the database
	MaxOpenConns int
	// the maximum amount of time a connection may be reused
	ConnMaxLifetime time.Duration
}

var DbConfig *dbConfig

func init() {
	DbConfig = &dbConfig{
		Name:            "default",
		Ip:              "",
		Port:            0,
		DbName:          "",
		User:            "",
		Password:        "",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: time.Hour,
	}
}

func (config *dbConfig) ReadFromFile() error {
	content, err := os.ReadFile(DB_CONFIG_PATH)
	if err != nil {
		fmt.Fprintf(os.Stderr, "logger: there is no any config file, using default config, error: %v\n", err)
	} else {
		if err := yaml.Unmarshal(content, &config); err != nil {
			return err
		}
	}
	return nil
}

func (config *dbConfig) AddFlags(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&config.Ip, "db-ip", "", "Value to indicate the host of database")
	flagSet.IntVar(&config.Port, "db-port", 0, "Value to indicate the port of database")
	flagSet.StringVar(&config.DbName, "db-name", "", "Value to indicate the name of database")
	flagSet.StringVar(&config.User, "db-user", "", "Value to indicate the user of database")
	flagSet.StringVar(&config.Password, "db-password", "", "Value to indicate the password of database")
	flagSet.IntVar(&config.MaxIdleConns, "db-max-idle-conns", 10, "Value to indicate the max idle connections of database")
	flagSet.IntVar(&config.MaxOpenConns, "db-max-open-conns", 100, "Value to indicate the max open connections of database")
}

func (config dbConfig) Validate() []error {
	// do nothing
	var errs []error
	return errs
}
