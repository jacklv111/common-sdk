/*
 * Created on Sun Jul 09 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package shedlock

import (
	"github.com/jacklv111/common-sdk/database"
	"github.com/jacklv111/common-sdk/log"
	"github.com/jacklv111/common-sdk/scheduler/shedlock/do"
)

//go:generate mockgen -source=interface.go -destination=./mock/mock_interface.go -package=mock

type ShedlockInterface interface {

	// TryLock 尝试获取分布式锁。非阻塞。
	//  @return bool true: 获取锁成功；false: 获取锁失败
	TryLock() bool

	// UnLock 释放锁
	UnLock()
}

func Init() error {
	if database.Db.Migrator().HasTable(&do.ShedlockDo{}) {
		log.Infof("table %s already exists", do.ShedlockDo{}.TableName())
		return nil
	}
	if err := database.Db.Migrator().CreateTable(&do.ShedlockDo{}); err != nil {
		return err
	}
	return nil
}
