/*
 * Created on Sun Jul 09 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package scheduler

import (
	"time"

	"github.com/jacklv111/common-sdk/log"
	"github.com/jacklv111/common-sdk/scheduler/shedlock"
)

// Schedule 创建一个定时任务
//
//	@param shedlockConfig
//	@param scheduleConfig
func Schedule(shedlockConfig shedlock.ShedlockConfig, scheduleConfig ScheduleConfig) {
	go run(shedlockConfig, scheduleConfig)
}

func run(shedlockConfig shedlock.ShedlockConfig, scheduleConfig ScheduleConfig) {
	var err error
	time.Sleep(scheduleConfig.InitialDelay)
	tick := time.NewTicker(time.Duration(scheduleConfig.Interval))
	var lock shedlock.ShedlockInterface
	if shedlockConfig.Enabled {
		lock, err = shedlock.NewShedLock(shedlockConfig)
		if err != nil {
			log.Errorf("schedule %s initialize failed", scheduleConfig.Name)
			return
		}
	}

	for range tick.C {
		if shedlockConfig.Enabled {
			succ := lock.TryLock()
			if !succ {
				continue
			}
			scheduleConfig.Runnable()

			lock.UnLock()
		} else {
			scheduleConfig.Runnable()
		}
	}
}

func WaitCondition(intervalInSec int, condition func() bool) {
	for {
		if condition() {
			break
		}
		time.Sleep(time.Duration(intervalInSec) * time.Second)
	}
}
