/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package mapset

type threadUnsafeMapSet[T comparable] map[T]bool

func NewSet[T comparable](vals ...T) threadUnsafeMapSet[T] {
	set := make(threadUnsafeMapSet[T])
	for _, val := range vals {
		set[val] = true
	}
	return set
}

func (s threadUnsafeMapSet[T]) Add(value T) {
	s[value] = true
}

func (s threadUnsafeMapSet[T]) Remove(value T) {
	delete(s, value)
}

func (s threadUnsafeMapSet[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

func (s threadUnsafeMapSet[T]) Len() int {
	return len(s)
}

func (s threadUnsafeMapSet[T]) Elements() []T {
	elements := make([]T, 0, len(s))
	for element := range s {
		elements = append(elements, element)
	}
	return elements
}
