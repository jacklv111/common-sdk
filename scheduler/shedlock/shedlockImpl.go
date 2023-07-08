/*
 * Created on Sun Jul 09 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package shedlock

import (
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jacklv111/common-sdk/database"
	"github.com/jacklv111/common-sdk/log"
	"github.com/jacklv111/common-sdk/scheduler/shedlock/do"
	"github.com/jacklv111/common-sdk/utils"
)

type shedlockImpl struct {
	config        ShedlockConfig
	shedlockDo    do.ShedlockDo
	hasLockRecord bool
}

func NewShedLock(slc ShedlockConfig) (ShedlockInterface, error) {
	if !slc.Enabled {
		return nil, fmt.Errorf("shed lock is disabled")
	}
	hostName, err := getHostName()
	if err != nil {
		return nil, fmt.Errorf("get host name error %s", err)
	}
	return &shedlockImpl{config: slc, shedlockDo: do.ShedlockDo{LockedBy: hostName, Name: slc.Name}}, nil
}

func (lock *shedlockImpl) TryLock() bool {
	lock.shedlockDo.LockedAt = time.Now().UnixMilli()
	lock.shedlockDo.LockUntil = time.Now().Add(lock.config.LockAtMostFor).UnixMilli()

	if !lock.hasLockRecord {
		res := database.Db.Create(&lock.shedlockDo)
		if res.Error != nil {
			mysqlErr, ok := res.Error.(*mysql.MySQLError)
			// 1062 为主键冲突
			if !ok || mysqlErr.Number != 1062 {
				log.Errorf("create lock record error %s", res.Error)
				return false
			} else { // 主键冲突，record 存在
				lock.hasLockRecord = true
				log.Infof("mysql error %s", res.Error)
			}
		}
		// 插入成功
		if res.RowsAffected == 1 {
			lock.hasLockRecord = true
			return true
		}
	}

	res := database.Db.Model(&lock.shedlockDo).Where("lock_until < (?)", lock.shedlockDo.LockedAt).
		Updates(do.ShedlockDo{LockUntil: lock.shedlockDo.LockUntil, LockedAt: lock.shedlockDo.LockedAt, LockedBy: lock.shedlockDo.LockedBy})
	if res.Error != nil {
		log.Errorf("update lock record error %s", res.Error)
		return false
	}
	if res.RowsAffected == 1 {
		return true
	}

	return false
}

func (lock *shedlockImpl) UnLock() {

	lock.shedlockDo.LockUntil = utils.MaxInt64(time.UnixMilli(lock.shedlockDo.LockedAt).Add(lock.config.LockAtLeastFor).UnixMilli(), time.Now().UnixMilli())
	res := database.Db.Model(&lock.shedlockDo).Where("locked_by = ?", lock.shedlockDo.LockedBy).Update("lock_until", lock.shedlockDo.LockUntil)
	if res.RowsAffected == 0 {
		log.Infof("%s unlock failed, the lock may be acquired by others", lock.shedlockDo.LockedBy)
	}
}

func getHostName() (string, error) {
	hostName, err := os.Hostname()
	if err != nil {
		return "", err
	}
	return hostName + "_" + utils.GetHostIp(), nil
}
