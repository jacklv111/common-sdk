/*
 * Created on Sat Oct 19 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package msgsvcclient

import (
	"net/http"
	"time"

	"github.com/jacklv111/common-sdk/log"
	msgsvcclientgo "github.com/jacklv111/msgsvc-client-go"
)

var msgsvcClient *msgsvcclientgo.APIClient

func Init() (err error) {
	host := Config.GetServerUrl()
	log.Infof("usermng server host: %s", host)
	clientConfig := msgsvcclientgo.NewConfiguration()
	clientConfig.Servers = msgsvcclientgo.ServerConfigurations{
		{
			URL:         host,
			Description: "No description provided",
		},
	}
	client := http.Client{}
	client.Timeout = time.Second * time.Duration(Config.TimeoutInSec)
	clientConfig.HTTPClient = &client
	msgsvcClient = msgsvcclientgo.NewAPIClient(clientConfig)
	return nil
}

func GetClient() *msgsvcclientgo.APIClient {
	return msgsvcClient
}
