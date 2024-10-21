/*
 * Created on Sat Oct 19 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package zookeeper

import (
	"time"

	"github.com/go-zookeeper/zk"
	"github.com/jacklv111/common-sdk/collection/mapset"
	"github.com/jacklv111/common-sdk/log"
)

var conn *zk.Conn

func InitZooKeeper() (err error) {
	conn, _, err = zk.Connect(ZkConfig.GetHosts(), time.Second*10)
	if err != nil {
		log.Fatal("Unable to connect to ZooKeeper: %s", err)
	}
	return
}

// CreateNode 创建一个服务节点
func CreateNode(path, ip string) (err error) {
	// 将服务节点作为临时节点 (EPHEMERAL)，这样当服务断开时，ZooKeeper 会自动删除这个节点
	_, err = conn.Create(path, []byte(ip), int32(zk.FlagEphemeral), zk.WorldACL(zk.PermAll))
	if err != nil {
		log.Errorf("Unable to create node: %s", err)
	}
	return
}

// GetAllData 获取指定路径下的所有节点数据
func GetAllData(path string) (result [][]byte, err error) {
	result = make([][]byte, 0)
	children, _, err := conn.Children(path)
	if err != nil {
		log.Errorf("Unable to get children: %s", err)
		return
	}
	log.Infof("service children found under [%s]: %v", path, children)
	var data []byte
	for _, child := range children {
		data, _, err = conn.Get(path + "/" + child)
		if err != nil {
			log.Errorf("Unable to get data: %s", err)
			return
		}
		result = append(result, data)
	}
	return
}

// 处理服务实例变化，将变化结果发送到 channel 中
// 该函数会一直阻塞，直到服务实例发生变化
func HandleServiceChanges(servicePath string, changesChan chan<- ServiceChange) {
	// 获取初始服务列表，并设置监听器
	children, _, eventCh, err := conn.ChildrenW(servicePath)
	if err != nil {
		log.Errorf("Failed to watch service instances:", err)
		return
	}

	log.Errorf("Initial service instances:", children)

	for {
		// 等待子节点变化事件
		event := <-eventCh

		log.Infof("Child nodes have changed:", event)

		// 获取更新后的服务列表
		updatedChildren, _, newCh, err := conn.ChildrenW(servicePath)
		if err != nil {
			log.Errorf("Failed to re-watch service instances:", err)
			return
		}

		// 处理新增和删除的节点
		added, removed := diff(children, updatedChildren)

		// 将变化信息传递到 channel 中
		changesChan <- ServiceChange{
			Added:   added,
			Removed: removed,
		}

		// 更新当前子节点列表和事件通道
		children = updatedChildren
		eventCh = newCh
	}
}

// 计算两个子节点列表的差异
func diff(oldList, newList []string) (added []string, removed []string) {
	oldSet := mapset.NewSet[string]()
	newSet := mapset.NewSet[string]()

	for _, node := range oldList {
		oldSet.Add(node)
	}
	for _, node := range newList {
		newSet.Add(node)
	}

	// 找出新增的节点
	for node := range newSet {
		if oldSet.Contains(node) {
			added = append(added, node)
		}
	}

	// 找出移除的节点
	for node := range oldSet {
		if newSet.Contains(node) {
			removed = append(removed, node)
		}
	}

	return added, removed
}

// GetNodeIp 获取指定节点的 ip
// @param path 节点路径
func GetNodeIp(path string) (string, error) {
	data, _, err := conn.Get(path)
	if err != nil {
		log.Errorf("Unable to get data: %s", err)
		return "", err
	}
	return string(data), nil
}
