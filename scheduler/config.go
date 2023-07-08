/*
 * Created on Sun Jul 09 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package scheduler

import "time"

type ScheduleConfig struct {
	Name         string
	Interval     time.Duration
	InitialDelay time.Duration
	Runnable     func()
}
