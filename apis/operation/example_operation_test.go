// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package operation_test

import (
	"fmt"

	"github.com/Lead-NorthStar/miniprogram"
	"github.com/Lead-NorthStar/miniprogram/apis/operation"
)

func ExampleGetFeedback() {
	var ctx *miniprogram.Miniprogram

	resp, err := operation.GetFeedback(ctx)

	fmt.Println(resp, err)
}

func ExampleGetJsErrSearch() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := operation.GetJsErrSearch(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetPerformance() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := operation.GetPerformance(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetSceneList() {
	var ctx *miniprogram.Miniprogram

	resp, err := operation.GetSceneList(ctx)

	fmt.Println(resp, err)
}

func ExampleGetVersionList() {
	var ctx *miniprogram.Miniprogram

	resp, err := operation.GetVersionList(ctx)

	fmt.Println(resp, err)
}

func ExampleRealtimelogSearch() {
	var ctx *miniprogram.Miniprogram

	resp, err := operation.RealtimelogSearch(ctx)

	fmt.Println(resp, err)
}
