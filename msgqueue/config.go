/*
 * Created on Fri Oct 18 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package msgqueue

import (
	"github.com/spf13/pflag"
)

type msgQueueConfig struct {
	// Hosts is a list of the servers to connect to.
	// The host format is a comma separated list of "hostname:port" string.
	// Host example: "hostname1:6650"
	Host string
}

var MsgQueueConfig *msgQueueConfig

func init() {
	MsgQueueConfig = &msgQueueConfig{
		Host: "localhost:6650",
	}
}

func (config *msgQueueConfig) ReadFromFile() error {
	return nil
}

func (config *msgQueueConfig) AddFlags(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&config.Host, "msgqueue-host", "localhost:6650", "Value to indicate the host of service.")
}

func (config msgQueueConfig) Validate() []error {
	return []error{}
}
