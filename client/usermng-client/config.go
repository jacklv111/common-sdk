/*
 * Created on Sat Oct 19 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package usermngclient

import (
	"fmt"

	"github.com/spf13/pflag"
)

type config struct {
	Ip           string `json:"ip"`
	Port         int    `json:"port"`
	TimeoutInSec int    `json:"timeout_in_sec"`
}

const (
	SERVER_CONFIG_PATH = "./conf/usermngclient_server.json"
)

var Config *config

func init() {
	Config = &config{
		Ip:           "10.11.0.12",
		Port:         8080,
		TimeoutInSec: 1200,
	}
}

func (cfg config) GetServerUrl() string {
	return fmt.Sprintf("https://%s:%d/api/open/v1", cfg.Ip, cfg.Port)
}

func (cfg *config) ReadFromFile() error {
	return nil
}

func (cfg *config) AddFlags(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&cfg.Ip, "usermngclient-server-ip", cfg.Ip, "usermngclient server ip")
	flagSet.IntVar(&cfg.Port, "usermngclient-server-port", cfg.Port, "usermngclient server port")
	flagSet.IntVar(&cfg.TimeoutInSec, "usermngclient-client-timeout-in-sec", cfg.TimeoutInSec, "usermngclient timeout in second")
}

func (cfg *config) Validate() []error {
	return []error{}
}
