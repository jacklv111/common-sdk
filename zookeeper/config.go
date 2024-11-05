/*
 * Created on Fri Oct 18 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package zookeeper

import (
	"errors"
	"regexp"
	"strings"

	"github.com/spf13/pflag"
)

type zookeeperConfig struct {
	// Hosts is a list of the Zookeeper servers to connect to.
	// The host format is a comma separated list of "hostname:port" string.
	// Hosts example: "zk1:2181,zk2:2181,zk3:2181"
	Hosts string
}

var ZkConfig *zookeeperConfig

func init() {
	ZkConfig = &zookeeperConfig{
		Hosts: "localhost:2181",
	}
}

func (config *zookeeperConfig) ReadFromFile() error {
	return nil
}

func (config *zookeeperConfig) AddFlags(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&config.Hosts, "zk-hosts", "localhost:2181", "Value to indicate the hosts of zookeeper service. The host format is a comma separated list of 'hostname:port' string. Hosts example: 'zk1:2181,zk2:2181,zk3:2181'")
}

func (config zookeeperConfig) Validate() []error {
	re := regexp.MustCompile(`^([\w\.-]+:\d+)(,[\w\.-]+:\d+)*$`)
	// 使用正则表达式匹配
	if re.MatchString(config.Hosts) {
		return []error{}
	}
	return []error{errors.New("invalid zookeeper hosts")}
}

func (config *zookeeperConfig) GetHosts() []string {
	return strings.Split(config.Hosts, ",")
}

const (
	ZOOKEEPER_CONFIG_PATH = "./conf/zookeeper.json"
)
