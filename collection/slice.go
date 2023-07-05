/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package collection

import "math"

// BatchRange 将 slice 中的数据按照 batch 的方式交给 consumer 消费
//
//	@param items
//	@param batchSize
//	@param consumer
//	@return error
func BatchRange[T any](items []T, batchSize int, consumer func([]T) error) error {
	itemLen := len(items)
	for start := 0; start < itemLen; start += batchSize {
		end := start + batchSize
		if end > itemLen {
			end = itemLen
		}
		err := consumer(items[start:end])
		if err != nil {
			return err
		}
	}
	return nil
}

// Sum
//
//	@param values
//	@return T
func Sum[T int | int16 | int32 | int64 | float32 | float64](values []T) T {
	var result T
	for _, value := range values {
		result += value
	}
	return result
}

// DivideItems
//
//	@param items
//	@param ratio
//	@return len(ratio) parts of []T
func DivideItems[T any](items []T, ratio []int) [][]T {
	total := len(items)
	parts := make([][]T, len(ratio))
	partSize := float64(total) / float64(Sum(ratio))
	startIndex := 0
	for i, r := range ratio {
		endIndex := startIndex + int(math.Ceil(float64(r)*partSize))
		if i == len(ratio)-1 {
			endIndex = total
		}
		parts[i] = items[startIndex:endIndex]
		startIndex = endIndex
	}
	return parts
}
