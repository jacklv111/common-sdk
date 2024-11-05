/*
 * Created on Fri Oct 18 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */
package zookeeper

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConfigFunc(t *testing.T) {
	tests := []struct {
		name    string
		hosts   string
		wantErr bool
	}{
		{"test1", "localhost:2181", false},
		{"test2", "localhost:2181,localhost:2182", false},
		{"test3", "localhost:2181,localhost:2182,localhost:2183", false},
		{"test4", "localhost", true},
		{"test5", "localhost:gggg", true},
		{"test6", "localhost:2181,localhost", true},
		{"test7", "139.9.36.40:32043", false},
	}
	for _, tt := range tests {
		ZkConfig.Hosts = tt.hosts
		Convey(tt.name, t, func() {
			errs := ZkConfig.Validate()
			if tt.wantErr {
				So(len(errs), ShouldBeGreaterThan, 0)
			} else {
				So(len(errs), ShouldEqual, 0)
			}
		})
	}
}
