// Copyright (c) Huawei Technologies Co., Ltd. 2019. All rights reserved.
// iSulad-img licensed under the Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//     http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
// PURPOSE.
// See the Mulan PSL v2 for more details.
// Description: iSulad image kit
// Author: lifeng
// Create: 2019-05-06

package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func containerRemove(gopts *globalOptions, idOrName string) error {
	imageService, err := getImageService(gopts)
	if err != nil {
		return err
	}

	storageRuntimeService := getRuntimeService("", imageService)
	if storageRuntimeService == nil {
		return errors.New("Failed to get storageRuntimeService")
	}

	err = storageRuntimeService.RemoveContainer(idOrName)
	if err != nil {
		return fmt.Errorf("failed to remove container %s: %v", idOrName, err)
	}
	return err
}
