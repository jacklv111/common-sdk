/*
 * Created on Sat Oct 19 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package usermngclient

import (
	"net/http"
	"time"

	"github.com/jacklv111/common-sdk/log"
	usermngclientgo "github.com/jacklv111/usermng-client-go"
)

var usermngClient *usermngclientgo.APIClient

func InitClientV2() (err error) {
	host := Config.GetServerUrl()
	log.Infof("usermng server host: %s", host)
	clientConfig := usermngclientgo.NewConfiguration()
	clientConfig.Servers = usermngclientgo.ServerConfigurations{
		{
			URL:         host,
			Description: "No description provided",
		},
	}
	client := http.Client{}
	client.Timeout = time.Second * time.Duration(Config.TimeoutInSec)
	clientConfig.HTTPClient = &client
	usermngClient = usermngclientgo.NewAPIClient(clientConfig)
	return nil
}

func GetUsermngClient() *usermngclientgo.APIClient {
	return usermngClient
}
