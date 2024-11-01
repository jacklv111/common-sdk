/*
 * Created on Fri Oct 18 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package msgqueue

import (
	"context"
	"fmt"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/jacklv111/common-sdk/log"
)

var client pulsar.Client

// 初始化消息队列
func InitMsgQueue() (err error) {
	// 创建 Pulsar 客户端
	client, err = pulsar.NewClient(pulsar.ClientOptions{
		// Pulsar 服务地址
		// examples: pulsar://localhost:6650
		URL: fmt.Sprintf("pulsar://%s", MsgQueueConfig.Host),
	})
	if err != nil {
		log.Fatal(err)
	}
	return
}

// Send 发送消息
// @param ctx
// @param topic
// @param message 要发送的消息
func Send(ctx context.Context, topic string, message []byte) error {
	// 创建生产者
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})
	if err != nil {
		log.Errorf("Failed to create producer: %v", err)
		return err
	}
	defer producer.Close()

	// 发送消息
	_, err = producer.Send(ctx, &pulsar.ProducerMessage{
		Payload: message,
	})
	if err != nil {
		log.Errorf("Failed to send message, %v", err)
		return err
	}
	return nil
}

// GetConsumer 获取消费者
func GetConsumer(topic string) (pulsar.Consumer, error) {
	// 创建消费者
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,            // 订阅的主题
		SubscriptionName: topic,            // 订阅名
		Type:             pulsar.Exclusive, // 订阅类型, Exclusive 表示独占消费
	})

	if err != nil {
		log.Errorf("Failed to create consumer: %v", err)
		return nil, err
	}
	return consumer, nil
}
