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
	args []string
}

func (appError AppError) Error() string {
	return appError.err.Error()
}

func (appError AppError) Code() string {
	return appError.code
}

func (appError AppError) Args() []string {
	return appError.args
}

// NewAppErr create a new AppError
// code: error code
// err: error
// args: args in the error message template
func NewAppErr(code string, err error, args ...string) AppError {
	return AppError{
		code: code,
		err:  err,
		args: args,
	}
}
