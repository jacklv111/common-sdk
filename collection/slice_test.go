/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package collection

import (
	"reflect"
	"testing"
)

func TestDivideItems(t *testing.T) {
	// Test case 1: divide 10 items into 3 parts with ratio 30:60:10
	items1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ratio1 := []int{30, 60, 10}
	expected1 := [][]int{
		{1, 2, 3},
		{4, 5, 6, 7, 8, 9},
		{10},
	}
	actual1 := DivideItems(items1, ratio1)
	if len(actual1) != len(expected1) {
		t.Errorf("Test case 1: Expected %d parts, but got %d", len(expected1), len(actual1))
	}
	for i := range expected1 {
		if !reflect.DeepEqual(actual1[i], expected1[i]) {
			t.Errorf("Test case 1: Part %d does not match expected result", i)
		}
	}

	// Test case 2: divide 5 items into 2 parts with ratio 50:50
	items2 := []int{1, 2, 3, 4, 5}
	ratio2 := []int{50, 50}
	expected2 := [][]int{
		{1, 2, 3},
		{4, 5},
	}
	actual2 := DivideItems(items2, ratio2)
	if len(actual2) != len(expected2) {
		t.Errorf("Test case 2: Expected %d parts, but got %d", len(expected2), len(actual2))
	}
	for i := range expected2 {
		if !reflect.DeepEqual(actual2[i], expected2[i]) {
			t.Errorf("Test case 2: Part %d does not match expected result", i)
		}
	}

	// Test case 3: divide 7 items into 3 parts with ratio 10:20:70
	items3 := []int{1, 2, 3, 4, 5, 6, 7}
	ratio3 := []int{10, 20, 70}
	expected3 := [][]int{
		{1},
		{2, 3},
		{4, 5, 6, 7},
	}
	actual3 := DivideItems(items3, ratio3)
	if len(actual3) != len(expected3) {
		t.Errorf("Test case 3: Expected %d parts, but got %d", len(expected3), len(actual3))
	}
	for i := range expected3 {
		if !reflect.DeepEqual(actual3[i], expected3[i]) {
			t.Errorf("Test case 3: Part %d does not match expected result. actual: %v, expected: %v", i, actual3, expected3)
		}
	}

	// Test case 4: divide 3 items into 1 part with ratio 100
	items4 := []int{1, 2, 3}
	ratio4 := []int{100}
	expected4 := [][]int{
		{1, 2, 3},
	}
	actual4 := DivideItems(items4, ratio4)
	if len(actual4) != len(expected4) {
		t.Errorf("Test case 4: Expected %d parts, but got %d", len(expected4), len(actual4))
	}
	for i := range expected4 {
		if !reflect.DeepEqual(actual4[i], expected4[i]) {
			t.Errorf("Test case 4: Part %d does not match expected result", i)
		}
	}
}
