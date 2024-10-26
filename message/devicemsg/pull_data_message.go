/*
 * Created on Sat Oct 26 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */
package devicemsg

type PullDataMessage struct {
	DataId string `json:"id"`
	Ext    string `json:"ext"`

	Text string `json:"text"`
}
