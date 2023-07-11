/*
 * Created on Tue Jul 11 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package aifsclient

import (
	"fmt"

	"github.com/spf13/pflag"
)

type aifsConfig struct {
	ServerIp     string
	ServerPort   int
	TimeoutInSec int
}

const (
	AIFS_SERVER_CONFIG_PATH = "./conf/aifs_server.json"
)

var AifsConfig *aifsConfig

func init() {
	AifsConfig = &aifsConfig{
		ServerIp:     "10.11.0.12",
		ServerPort:   8080,
		TimeoutInSec: 1200,
	}
}

func (cfg aifsConfig) GetServerUrl() string {
	return fmt.Sprintf("http://%s:%d/api/open/v1", cfg.ServerIp, cfg.ServerPort)
}

func (cfg *aifsConfig) ReadFromFile() error {
	return nil
}

func (cfg *aifsConfig) AddFlags(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&cfg.ServerIp, "aifs-server-ip", cfg.ServerIp, "aifs server ip")
	flagSet.IntVar(&cfg.ServerPort, "aifs-server-port", cfg.ServerPort, "aifs server port")
	flagSet.IntVar(&cfg.TimeoutInSec, "aifs-client-timeout-in-sec", cfg.TimeoutInSec, "aifs timeout in second")
}

func (cfg *aifsConfig) Validate() []error {
	return []error{}
}
