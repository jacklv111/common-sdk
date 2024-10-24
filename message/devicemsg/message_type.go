/*
 * Created on Wed Oct 23 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */
package devicemsg

// MessageType 定义消息类型
type MessageType int

const (
	ACK MessageType = iota
	PULLDATA
)
