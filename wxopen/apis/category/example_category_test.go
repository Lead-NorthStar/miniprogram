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

package category_test

import (
	"fmt"

	"github.com/Lead-NorthStar/miniprogram"
	"github.com/Lead-NorthStar/miniprogram/wxopen/apis/category"
)

func ExampleGetAllCategories() {
	var ctx *miniprogram.Miniprogram

	resp, err := category.GetAllCategories(ctx)

	fmt.Println(resp, err)
}

func ExampleGetCategory() {
	var ctx *miniprogram.Miniprogram

	resp, err := category.GetCategory(ctx)

	fmt.Println(resp, err)
}

func ExampleAddCategory() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := category.AddCategory(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDeleteCategory() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := category.DeleteCategory(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleModifyCategory() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := category.ModifyCategory(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleWxaCategory() {
	var ctx *miniprogram.Miniprogram

	resp, err := category.GetCategory(ctx)

	fmt.Println(resp, err)
}
