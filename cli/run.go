/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// 执行 cobra 命令，做一个包装，封装命令执行前后需要做的事情和处理执行失败的异常
func Run(cmd *cobra.Command) int {
	if err := run(cmd); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}
	return 0
}

func run(cmd *cobra.Command) (err error) {
	err = cmd.Execute()
	return
}
