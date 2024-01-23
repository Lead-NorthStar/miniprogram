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

package pay_test

import (
	"fmt"

	"github.com/Lead-NorthStar/miniprogram"
	"github.com/Lead-NorthStar/miniprogram/minigame/apis/pay"
)

func ExampleCancelPay() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := pay.CancelPay(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetBalance() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := pay.GetBalance(ctx, payload)

	fmt.Println(resp, err)
}

func ExamplePay() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := pay.Pay(ctx, payload)

	fmt.Println(resp, err)
}

func ExamplePresent() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := pay.Present(ctx, payload)

	fmt.Println(resp, err)
}
