/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package utils

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// 将字符串表示形式的数字解析为 int
// 如果是合法数字，值要在 [minVal, maxVal] 区间内
// 如果是空字符串，则返回 defaultVal
// 如果字符串不是数字，返回 error
func ParseInt(input string, minVal int, maxVal int, defaultVal int) (result int, err error) {
	if input == "" {
		return defaultVal, nil
	}
	result, err = strconv.Atoi(input)
	if err != nil {
		return
	}
	if !(result >= minVal && result <= maxVal) {
		return result, errors.Errorf("value is not in range [%d, %d]", minVal, maxVal)
	}
	return
}

// 解析如 “a,b,c" -> ["a", "b", "c"] 形式的参数
func ParseListStr(param string, hasParam bool, seq string) []string {
	if !hasParam {
		return nil
	}
	return strings.Split(param, seq)
}
