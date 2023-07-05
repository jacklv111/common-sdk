/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package database

import (
	"database/sql"
	"fmt"

	gormlocal "github.com/jacklv111/common-sdk/database/gorm"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() (err error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbConfig.User, DbConfig.Password, DbConfig.Ip, DbConfig.Port, DbConfig.DbName)

	Db, err = gorm.Open(mysql.Open(dataSourceName), gormlocal.GormConfig)

	if err != nil {
		return err
	}

	sqlDB, err := Db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(DbConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(DbConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(DbConfig.ConnMaxLifetime)
	return nil
}

// for test
func SetUpSqlMocker() (*sql.DB, sqlmock.Sqlmock, error) {
	dbConn, sqlMocker, err := sqlmock.New()
	if err != nil {
		return nil, nil, fmt.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	Db, err = gorm.Open(mysql.New(mysql.Config{Conn: dbConn, SkipInitializeWithVersion: true}), gormlocal.GormConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("gorm open error %s", err)
	}
	return dbConn, sqlMocker, nil
}
