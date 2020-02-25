// Copyright (c) Huawei Technologies Co., Ltd. 2020. All rights reserved.
// iSulad-img licensed under the Mulan PSL v1.
// You can use this software according to the terms and conditions of the Mulan PSL v1.
// You may obtain a copy of Mulan PSL v1 at:
//     http://license.coscl.org.cn/MulanPSL
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
// PURPOSE.
// See the Mulan PSL v1 for more details.
// Description: iSulad image kit
// Author: wangfengtu
// Create: 2020-02-25

package main

import (
	"testing"
)

func TestGetDiskUsageStats(t *testing.T) {
	// Make sure it do not panic
	_, _, err := GetDiskUsageStats("/proc")
	if err != nil {
		t.Errorf("GetDiskUsageStats failed: %v", err)
	}
}
