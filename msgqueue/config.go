/*
 * Created on Fri Oct 18 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package msgqueue

import (
	"errors"
	"regexp"
	"strings"

	"github.com/spf13/pflag"
)

type msgQueueConfig struct {
	// Hosts is a list of the servers to connect to.
	// The host format is a comma separated list of "hostname:port" string.
	// Hosts example: "hostname1:6650,hostname1:6650,hostname1:6650"
	Hosts string
}

var MsgQueueConfig *msgQueueConfig

func init() {
	MsgQueueConfig = &msgQueueConfig{
		Hosts: "localhost:6650",
	}
}

func (config *msgQueueConfig) ReadFromFile() error {
	return nil
}

func (config *msgQueueConfig) AddFlags(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&config.Hosts, "msgqueue-hosts", "localhost:6650", "Value to indicate the hosts of service. The host format is a comma separated list of 'hostname:port' string. Hosts example: 'hostname1:6650,hostname1:6650,hostname1:6650'")
}

func (config msgQueueConfig) Validate() []error {
	re := regexp.MustCompile(`^(\w+:\d+)(,\w+:\d+)*$`)
	// 使用正则表达式匹配
	if re.MatchString(config.Hosts) {
		return []error{}
	}
	return []error{errors.New("invalid msgqueue hosts")}
}

func (config *msgQueueConfig) GetHosts() []string {
	return strings.Split(config.Hosts, ",")
}
