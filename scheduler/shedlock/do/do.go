/*
 * Created on Sun Jul 09 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package do

type ShedlockDo struct {
	Name string `gorm:"primaryKey"`
	// unix milli
	LockUntil int64
	// unix milli
	LockedAt int64
	LockedBy string
}

func (ShedlockDo) TableName() string {
	return "shedlocks"
}
