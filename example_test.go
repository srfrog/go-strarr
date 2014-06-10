// Copyright 2014 Codehack.com All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strarr_test

import(
	"github.com/codehack/go-strarr"
	"fmt"
)

// This example shows basic usage of various functions by manipulating
// the array 'arr'.
func Example_basic() {
	arr := []string{"Go", "nuts", "for", "Go"}

	foo := strarr.Repeat("Go",3)
	fmt.Println(foo)

	fmt.Println(strarr.Count(arr, "Go"))

	fmt.Println(strarr.Index(arr, "Go"))
	fmt.Println(strarr.LastIndex(arr, "Go"))

	if strarr.Contains(arr, "nuts") {
		arr = strarr.Replace(arr, []string{"Insanely"})
	}
	fmt.Println(arr)

	str := strarr.Shift(&arr)
	fmt.Println(str)
	fmt.Println(arr)

	strarr.Unshift(&arr, "Really")
	fmt.Println(arr)

	fmt.Println(strarr.ToUpper(arr))
	fmt.Println(strarr.ToLower(arr))
	fmt.Println(strarr.ToTitle(arr))

	fmt.Println(strarr.Trim(arr,"Really"))
	fmt.Println(strarr.Filter(arr,"Go"))

	fmt.Println(strarr.Diff(arr,foo))
	fmt.Println(strarr.Intersect(arr,foo))

	// this will return 2 random elements from arr
	// but it will fail testing, so...
	// fmt.Println(strarr.Rand(arr,2))

	fmt.Println(strarr.Reverse(arr))

	// Output:
	// [Go Go Go]
	// 2
	// 0
	// 3
	// [Insanely nuts for Go]
	// Insanely
	// [nuts for Go]
	// [Really nuts for Go]
	// [REALLY NUTS FOR GO]
	// [really nuts for go]
	// [Really Nuts For Go]
	// [nuts for Go]
	// [Go]
	// [Really nuts for]
	// [Go]
	// [Go for nuts Really]
}

