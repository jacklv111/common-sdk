/*
 * Created on Thu Feb 02 2023
 *
 * Copyright (c) 2023 Gddi
 */
package shedlock

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/jacklv111/common-sdk/log"
	. "github.com/jacklv111/common-sdk/test"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_shedlockImpl(t *testing.T) {
	defer DbSetUpAndTearDown()()
	log.ValidateAndApply(log.LogConfig)

	Convey("try lock", t, func() {

		lockConfig := ShedlockConfig{Enabled: true, Name: "test", LockAtLeastFor: time.Second, LockAtMostFor: time.Second * 2}
		lock, _ := NewShedLock(lockConfig)

		Convey("Success", func() {
			Convey("no lock record and insert a record", func() {
				Sqlmocker.ExpectBegin()
				Sqlmocker.ExpectExec(regexp.QuoteMeta("INSERT INTO `shedlocks` (`name`,`lock_until`,`locked_at`,`locked_by`) VALUES (?,?,?,?)")).
					WithArgs(lockConfig.Name, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
				Sqlmocker.ExpectCommit()

				res := lock.TryLock()

				So(res, ShouldBeTrue)
			})

			Convey("lock record existed, update record", func() {
				err := &mysql.MySQLError{Number: 1062}
				Sqlmocker.ExpectBegin()
				Sqlmocker.ExpectExec(regexp.QuoteMeta("INSERT INTO `shedlocks` (`name`,`lock_until`,`locked_at`,`locked_by`) VALUES (?,?,?,?)")).
					WithArgs(lockConfig.Name, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(err)
				Sqlmocker.ExpectRollback()

				Sqlmocker.ExpectBegin()
				Sqlmocker.ExpectExec(regexp.QuoteMeta("UPDATE `shedlocks` SET `lock_until`=?,`locked_at`=?,`locked_by`=? WHERE lock_until < (?) AND `name` = ?")).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), lockConfig.Name).
					WillReturnResult(sqlmock.NewResult(1, 1))
				Sqlmocker.ExpectCommit()

				res := lock.TryLock()

				So(res, ShouldBeTrue)
			})
		})

		Convey("Failed", func() {
			Convey("insert error", func() {
				err := &mysql.MySQLError{Number: 1060}
				Sqlmocker.ExpectBegin()
				Sqlmocker.ExpectExec(regexp.QuoteMeta("INSERT INTO `shedlocks` (`name`,`lock_until`,`locked_at`,`locked_by`) VALUES (?,?,?,?)")).
					WithArgs(lockConfig.Name, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(err)
				Sqlmocker.ExpectRollback()

				res := lock.TryLock()

				So(res, ShouldBeFalse)
			})

			Convey("record existed, update affect 0", func() {
				err := &mysql.MySQLError{Number: 1062}
				Sqlmocker.ExpectBegin()
				Sqlmocker.ExpectExec(regexp.QuoteMeta("INSERT INTO `shedlocks` (`name`,`lock_until`,`locked_at`,`locked_by`) VALUES (?,?,?,?)")).
					WithArgs(lockConfig.Name, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(err)
				Sqlmocker.ExpectRollback()

				Sqlmocker.ExpectBegin()
				Sqlmocker.ExpectExec(regexp.QuoteMeta("UPDATE `shedlocks` SET `lock_until`=?,`locked_at`=?,`locked_by`=? WHERE lock_until < (?) AND `name` = ?")).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), lockConfig.Name).
					WillReturnResult(sqlmock.NewResult(0, 0))
				Sqlmocker.ExpectCommit()

				res := lock.TryLock()

				So(res, ShouldBeFalse)
			})
		})
	})

	Convey("unlock", t, func() {

		lockConfig := ShedlockConfig{Enabled: true, Name: "test", LockAtLeastFor: time.Second, LockAtMostFor: time.Second * 2}
		lock, _ := NewShedLock(lockConfig)

		Convey("Success", func() {
			Sqlmocker.ExpectBegin()
			Sqlmocker.ExpectExec(regexp.QuoteMeta("UPDATE `shedlocks` SET `lock_until`=? WHERE locked_by = ? AND `name` = ?")).
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), lockConfig.Name).
				WillReturnResult(sqlmock.NewResult(1, 1))
			Sqlmocker.ExpectCommit()

			lock.UnLock()
		})
	})
}
