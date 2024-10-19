/*
 * Created on Sat Oct 19 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package pusherclient

import (
	"fmt"

	"github.com/spf13/pflag"
)

type config struct {
	Ip           string
	Port         int
	TimeoutInSec int
}

const (
	SERVER_CONFIG_PATH = "./conf/pusherclient_server.json"
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
	flagSet.StringVar(&cfg.Ip, "pusherclient-server-ip", cfg.Ip, "pusherclient server ip")
	flagSet.IntVar(&cfg.Port, "pusherclient-server-port", cfg.Port, "pusherclient server port")
	flagSet.IntVar(&cfg.TimeoutInSec, "pusherclient-client-timeout-in-sec", cfg.TimeoutInSec, "pusherclient timeout in second")
}

func (cfg *config) Validate() []error {
	return []error{}
}
