// Various operations for string slices for Go
//
// Go-StrArr is a collection of functions to manipulate string arrays/slices.
// Some functions were adapted from the strings package to work with string slices, other
// were ported from PHP 'array_*' function equivalents.
//
// Example:
//
//		package main
//
//		import(
//			"github.com/codehack/go-strarr"
//			"fmt"
//		)
//
//		func main() {
//			arr := []string{"Go", "nuts", "for", "Go"}
//
//			foo := strarr.Repeat("Go",3)
//			fmt.Println(foo)
//
//			fmt.Println(strarr.Count(arr, "Go"))
//
//			fmt.Println(strarr.Index(arr, "Go"))
//			fmt.Println(strarr.LastIndex(arr, "Go"))
//
//			if strarr.Contains(arr, "nuts") {
//				arr = strarr.Replace(arr, []string{"Insanely"})
//			}
//			fmt.Println(arr)
//
//			str := strarr.Shift(&arr)
//			fmt.Println(str)
//			fmt.Println(arr)
//
//			strarr.Unshift(&arr, "Really")
//			fmt.Println(arr)
//
//			fmt.Println(strarr.ToUpper(arr))
//			fmt.Println(strarr.ToLower(arr))
//			fmt.Println(strarr.ToTitle(arr))
//
//			fmt.Println(strarr.Trim(arr,"Really"))
//			fmt.Println(strarr.Filter(arr,"Go"))
//
//			fmt.Println(strarr.Diff(arr,foo))
//			fmt.Println(strarr.Intersect(arr,foo))
//
//			fmt.Println(strarr.Rand(arr,2))
//
//			fmt.Println(strarr.Reverse(arr))
// 	}
//
//
// See example_test.go for more details. Also, strarr_test.go has plenty of tests to
// serve as examples. Enjoy.
//
package strarr
