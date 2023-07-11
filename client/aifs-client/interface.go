/*
 * Created on Tue Jul 11 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package aifsclient

import (
	"net/http"
	"time"

	aifsclientgo "github.com/jacklv111/aifs-client-go"
	httpclient "github.com/jacklv111/aifs-client-go/http-client"
	"github.com/jacklv111/common-sdk/log"
	"github.com/jacklv111/common-sdk/utils"
)

var aifsClient *aifsclientgo.APIClient

func InitAifsClient() (err error) {
	host, err := utils.GetAifsConfigFromDataClient()
	if err != nil {
		return err
	}
	log.Infof("aifs host: %s", host)
	clientConfig := aifsclientgo.NewConfiguration()
	clientConfig.Servers = aifsclientgo.ServerConfigurations{
		{
			URL:         host,
			Description: "No description provided",
		},
	}
	aifsClient = aifsclientgo.NewAPIClient(clientConfig)
	return nil
}

func InitAifsClientV2() (err error) {
	host := AifsConfig.GetServerUrl()
	log.Infof("aifs host: %s", host)
	clientConfig := aifsclientgo.NewConfiguration()
	clientConfig.Servers = aifsclientgo.ServerConfigurations{
		{
			URL:         host,
			Description: "No description provided",
		},
	}
	client := http.Client{}
	client.Timeout = time.Second * time.Duration(AifsConfig.TimeoutInSec)
	clientConfig.HTTPClient = &client
	aifsClient = aifsclientgo.NewAPIClient(clientConfig)
	return nil
}

// InitAifsClientV3 is used for testing
func InitAifsClientV3(client httpclient.HTTPClient) (err error) {
	host := AifsConfig.GetServerUrl()
	log.Infof("aifs host: %s", host)
	clientConfig := aifsclientgo.NewConfiguration()
	clientConfig.Servers = aifsclientgo.ServerConfigurations{
		{
			URL:         host,
			Description: "No description provided",
		},
	}
	clientConfig.HTTPClient = client

	aifsClient = aifsclientgo.NewAPIClient(clientConfig)
	return nil
}

func GetAifsClient() *aifsclientgo.APIClient {
	return aifsClient
}
