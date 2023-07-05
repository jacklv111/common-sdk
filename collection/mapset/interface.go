/*
 * Created on Mon Jul 03 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package mapset

//go:generate mockgen -source=interface.go -destination=./mock/mock_interface.go -package=mock

type Set[T comparable] interface {
	Add(value T)
	Remove(value T)
	Contains(value T) bool
	Len() int
	Elements() []T
}
