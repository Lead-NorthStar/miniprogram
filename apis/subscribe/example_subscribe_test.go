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

package subscribe_test

import (
	"fmt"

	"github.com/Lead-NorthStar/miniprogram"
	"github.com/Lead-NorthStar/miniprogram/apis/subscribe"
)

func ExampleAddTemplate() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := subscribe.AddTemplate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDeleteTemplate() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := subscribe.DeleteTemplate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetCategory() {
	var ctx *miniprogram.Miniprogram

	resp, err := subscribe.GetCategory(ctx)

	fmt.Println(resp, err)
}

func ExampleGetPubTemplateKeyWordsById() {
	var ctx *miniprogram.Miniprogram

	resp, err := subscribe.GetPubTemplateKeyWordsById(ctx)

	fmt.Println(resp, err)
}

func ExampleGetPubTemplateTitleList() {
	var ctx *miniprogram.Miniprogram

	resp, err := subscribe.GetPubTemplateTitleList(ctx)

	fmt.Println(resp, err)
}

func ExampleGetTemplateList() {
	var ctx *miniprogram.Miniprogram

	resp, err := subscribe.GetTemplateList(ctx)

	fmt.Println(resp, err)
}

func ExampleSend() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := subscribe.Send(ctx, payload)

	fmt.Println(resp, err)
}
