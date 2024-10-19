/*
 * Created on Fri Oct 18 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */
package security

import (
	"fmt"
	"testing"
)

func TestGenRandStr(t *testing.T) {
	str, err := GenRandStr(10)
	if err != nil {
		t.Errorf("GenRandStr() error = %v", err)
		return
	}
	if len(str) != 10 {
		t.Errorf("GenRandStr() = %v, want %v", len(str), 10)
	}
	fmt.Println(str)
}
