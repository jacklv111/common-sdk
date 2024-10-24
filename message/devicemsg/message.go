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
	CreateAt int64           `json:"create_at"` // 可以使用时间格式
	Type     MessageType     `json:"type"`
	Data     json.RawMessage `json:"data"` // 使用空接口存储任何类型的数据
}

func NewMessage(id string, createAt int64, messageType MessageType, data interface{}) *Message {
	dataJson, _ := json.Marshal(data)
	return &Message{
		ID:       id,
		CreateAt: createAt,
		Type:     messageType,
		Data:     dataJson,
	}
}
