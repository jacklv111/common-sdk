/*
 * Created on Wed Jul 05 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */
package errors

type AppError struct {
	err  error
	code string
	msg  string
}

func (appError AppError) Error() error {
	return appError.err
}

func (appError AppError) Code() string {
	return appError.code
}

func (appError AppError) Message() string {
	return appError.msg
}

func NewAppErr(code string, err error, msg string) AppError {
	return AppError{
		err:  err,
		code: code,
		msg:  msg,
	}
}
