/*
 * Created on Wed Oct 23 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */
package devicemsg

import (
	"encoding/json"
)

// Message 定义消息结构体
type Message struct {
	ID       string          `json:"id"`
	CreateAt int64           `json:"create_at"`
	Type     MessageType     `json:"type"`
	From     string          `json:"from"`
	ToUser   string          `json:"to_user"`
	ToGroup  string          `json:"to_group"`
	Data     json.RawMessage `json:"data"`
}

func NewMessage(id string, createAt int64, messageType MessageType,
	from string, toUser string, toGroup string, data interface{}) *Message {

	dataJson, _ := json.Marshal(data)
	return &Message{
		ID:       id,
		CreateAt: createAt,
		Type:     messageType,
		From:     from,
		ToUser:   toUser,
		ToGroup:  toGroup,
		Data:     dataJson,
	}
}
