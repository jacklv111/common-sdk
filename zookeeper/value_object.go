/*
 * Created on Sat Oct 19 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */
package zookeeper

// 封装服务节点的变动信息
type ServiceChange struct {
	Added   []string
	Removed []string
}
