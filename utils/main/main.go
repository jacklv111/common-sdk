/*
 * Created on Thu Oct 31 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */
package main

import (
	"fmt"

	"github.com/jacklv111/common-sdk/utils"
)

func main() {
	ak, sk, _ := utils.GenerateAKSK()
	fmt.Printf("ak: %s\nsk: %s\n", ak, sk)
}
