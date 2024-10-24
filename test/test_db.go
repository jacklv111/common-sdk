/*
 * Created on Sun Jul 09 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package test

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jacklv111/common-sdk/database"
)

var sqlDb *sql.DB
var Sqlmocker sqlmock.Sqlmock

func DbSetUpAndTearDown() func() {
	var err error

	sqlDb, Sqlmocker, err = database.SetUpSqlMocker()
	if err != nil {
		panic(err)
	}

	// tear down
	return func() {
		sqlDb.Close()
	}
}
