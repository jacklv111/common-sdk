/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package utils

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	gruntime "runtime"
)

func StartProf() {
	go func() {
		gruntime.SetBlockProfileRate(1)
		gruntime.SetMutexProfileFraction(1)
		err := http.ListenAndServe(":6061", nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "start prof error:: %v\n", err)
		}
	}()
}
