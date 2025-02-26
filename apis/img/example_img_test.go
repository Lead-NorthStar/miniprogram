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

package img_test

import (
	"fmt"
	"net/url"

	"github.com/Lead-NorthStar/miniprogram"
	"github.com/Lead-NorthStar/miniprogram/apis/img"
)

func ExampleAiCrop() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	params := url.Values{}
	resp, err := img.AiCrop(ctx, payload, params)

	fmt.Println(resp, err)
}

func ExampleScanQRCode() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	params := url.Values{}
	resp, err := img.ScanQRCode(ctx, payload, params)

	fmt.Println(resp, err)
}

func ExampleSuperresolution() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	params := url.Values{}
	resp, err := img.Superresolution(ctx, payload, params)

	fmt.Println(resp, err)
}
