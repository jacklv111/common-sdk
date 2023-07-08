/*
 * Created on Sun Jul 09 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package shedlock

import "time"

type ShedlockConfig struct {
	Enabled bool
	Name    string

	LockAtLeastFor time.Duration
	LockAtMostFor  time.Duration
}
