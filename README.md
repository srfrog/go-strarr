# Go-StrArr [![GoDoc](https://godoc.org/github.com/codehack/go-strarr?status.png)](https://godoc.org/github.com/codehack/go-strarr)

*Various operations for string slices for [Go](http://golang.org)*

**Go-StrArr** is a collection of functions that operate on string slices. Most are functions from `package strings` that have been adapted to work with string slices, and some were ported from PHP array_* functions.


## Features

- Using a thin layer of idiomatic Go; correctness over performance.
- Provide most basic array-ish operations: index, trim, filter, map
- Some PHP favorites like: pop, push, shift, unshift, shuffle, etc...
- Non-destructive returns (won't alter original slice), except for explicit tasks.
- The perfect companion to work with `net/http` headers!

## Installation

Using "go get":

	go get github.com/codehack/go-strarr

Then import:

	import ("github.com/codehack/go-strarr")

## Example

// This example shows basic usage of various functions by manipulating
// the array 'arr'.

``` go
package main

import(
   "github.com/codehack/go-strarr"
   "fmt"
)

// This example shows basic usage of various functions by manipulating
// the array 'arr'.
func main() {
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

   fmt.Println(strarr.Rand(arr,2))

   fmt.Println(strarr.Reverse(arr))
}

```

See the testing source [strarr_test.go](https://github.com/codehack/go-strarr/blob/master/strarr_test.go) for more examples.

## Documentation

The full code documentation is located at GoDoc:

[http://godoc.org/github.com/codehack/go-strarr](http://godoc.org/github.com/codehack/go-strarr)

**Go-StrArr** is Copyright (c) 2014 [Codehack](http://codehack.com).
Published under [MIT License](https://raw.githubusercontent.com/codehack/go-strarr/master/LICENSE)

_Go nuts for Go!_
