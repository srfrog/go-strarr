// Copyright 2014-present Codehack. All rights reserved.
// For mobile and web development visit http://codehack.com
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package strarr is a collection of functions to manipulate string arrays/slices.
// Some functions were adapted from the strings package to work with string slices, other
// were ported from PHP 'array_*' function equivalents.
package strarr

import (
	"math/rand"
	"strings"
)

// streq just compares 2 strings, used as base comparison.
func streq(s1, s2 string) bool { return s1 == s2 }

// indexFunc runs through 'a' and compares each element with 's' using 'f' function.
// It returns the index of the first occurance of 's' in 'a', or -1 if not found.
func indexFunc(a []string, s string, f func(string, string) bool) int {
	for i := range a {
		if f(a[i], s) {
			return i
		}
	}
	return -1
}

// Index returns the index of the first instance of 's' in 'a', or -1 if not found
func Index(a []string, s string) int {
	return indexFunc(a, s, streq)
}

// IndexPrefix returns the index of the first entry in 'a' with prefix 'prefix', or -1 if not found
func IndexPrefix(a []string, prefix string) int {
	return indexFunc(a, prefix, strings.HasPrefix)
}

// IndexSuffix returns the index of the first entry in 'a' with suffix 'suffix', or -1 if not found
func IndexSuffix(a []string, suffix string) int {
	return indexFunc(a, suffix, strings.HasSuffix)
}

// Search returns the index of the first entry containing the substring 's' in 'a', or -1 if not found
func Search(a []string, s string) int {
	return indexFunc(a, s, strings.Contains)
}

// Contains returns true if 's' is in 'a', false otherwise
func Contains(a []string, s string) bool {
	return Index(a, s) >= 0
}

// ContainsPrefix returns true if any entry in 'a' has prefix 'prefix', false otherwise
func ContainsPrefix(a []string, prefix string) bool {
	return IndexPrefix(a, prefix) >= 0
}

// ContainsSuffix returns true if any entry in 'a' has suffix 'suffix', false otherwise
func ContainsSuffix(a []string, suffix string) bool {
	return IndexSuffix(a, suffix) >= 0
}

// Count returns the number of occurrances of 's' in 'a'
func Count(a []string, s string) int {
	n := 0
	for i := range a {
		if a[i] == s {
			n++
		}
	}
	return n
}

// lastIndexFunc runs through 'a' and compares each element with 's' using 'f' function.
// It returns the index of the last occurance of 's' in 'a', or -1 if not found.
func lastIndexFunc(a []string, s string, f func(string, string) bool) int {
	for i := len(a) - 1; i >= 0; i-- {
		if f(a[i], s) {
			return i
		}
	}
	return -1
}

// LastIndex returns the index of the last instance of 's' in 'a', or -1 if not found
func LastIndex(a []string, s string) int {
	return lastIndexFunc(a, s, streq)
}

// LastIndexPrefix returns the index of the last entry in 'a' with prefix 'prefix', or -1 if not found
func LastIndexPrefix(a []string, prefix string) int {
	return lastIndexFunc(a, prefix, strings.HasPrefix)
}

// LastIndexSuffix returns the index of the last entry in 'a' with suffix 'suffix', or -1 if not found
func LastIndexSuffix(a []string, suffix string) int {
	return lastIndexFunc(a, suffix, strings.HasSuffix)
}

// LastSearch returns the index of the last entry containing the substring 's' in 'a', or -1 if not found
func LastSearch(a []string, s string) int {
	return lastIndexFunc(a, s, strings.Contains)
}

// Map returns a slice with the function 'mapping' applied to each element of 'a'
func Map(mapping func(string) string, a []string) []string {
	b := make([]string, len(a))
	for i := range a {
		b[i] = mapping(a[i])
	}
	return b
}

// ToUpper returns a slice with all entries of 'a' changed to upper case
func ToUpper(a []string) []string {
	return Map(strings.ToUpper, a)
}

// ToLower returns a slice with all entries of 'a' changed to lower case
func ToLower(a []string) []string {
	return Map(strings.ToLower, a)
}

// ToTitle returns a slice with all entries of 'a' changed to title case
func ToTitle(a []string) []string {
	return Map(strings.Title, a)
}

// trimFunc runs through 'a' and compares each element with 's' using 'f' comparison function (bool).
// It returns a slice of the elements in 'a' where 'f' returns false.
func trimFunc(a []string, s string, f func(string, string) bool) []string {
	b := []string{}
	for i := range a {
		if !f(a[i], s) {
			b = append(b, a[i])
		}
	}
	return b
}

// Trim returns a slice with all the entries of 'a' that don't match string 's'
func Trim(a []string, s string) []string {
	return trimFunc(a, s, streq)
}

// TrimFunc returns a slice with all the entries of 'a' that don't match string 's' using a callback function 'f'
// Callback is f(value, key string) where value is in 'a' and is checked for key. If true, value will be trimmed.
func TrimFunc(a []string, s string, f func(string, string) bool) []string {
	return trimFunc(a, s, f)
}

// TrimPrefix returns a slice with all the entries of 'a' that don't have prefix 'prefix'
func TrimPrefix(a []string, prefix string) []string {
	return trimFunc(a, prefix, strings.HasPrefix)
}

// TrimSuffix returns a slice with all the entries of 'a' that don't have suffix 'suffix'
func TrimSuffix(a []string, suffix string) []string {
	return trimFunc(a, suffix, strings.HasSuffix)
}

// filterFunc runs through 'a' and compares each element with 's' using 'f' comparison function (bool).
// It returns a slice of the elements in 'a' where 'f' returns true.
func filterFunc(a []string, s string, f func(string, string) bool) []string {
	// complement trimFunc comparison
	return trimFunc(a, s, func(v string, s string) bool { return !f(v, s) })
}

// Filter returns a slice with all the entries of 'a' that match string 's'
func Filter(a []string, s string) []string {
	return filterFunc(a, s, streq)
}

// FilterFunc returns a slice with all the entries of 'a' that match string 's' using a callback function 'f'
// Callback is f(value, key string) where value is in 'a' and is checked for key. If true, value will be filtered.
func FilterFunc(a []string, s string, f func(string, string) bool) []string {
	return filterFunc(a, s, f)
}

// FilterPrefix returns a slice with all the entries of 'a' that have prefix 'prefix'
func FilterPrefix(a []string, prefix string) []string {
	return filterFunc(a, prefix, strings.HasPrefix)
}

// FilterSuffix returns a slice with all the entries of 'a' that have suffix 'suffix'
func FilterSuffix(a []string, suffix string) []string {
	return filterFunc(a, suffix, strings.HasSuffix)
}

// diffFunc compares the elements of 'a' with those of 'b' using 'f' comparison function (bool).
// It returns a slice of the elements in 'a' that are not found in 'b' which f() == true.
func diffFunc(a, b []string, f func(a []string, s string) bool) []string {
	c := []string{}
	for i := range a {
		if f(b, a[i]) {
			c = append(c, a[i])
		}
	}
	return c
}

// Diff returns a slice with all the entries of 'a' that are not found in 'b'
func Diff(a, b []string) []string {
	return diffFunc(a, b, func(a []string, s string) bool { return !Contains(a, s) })
}

// Intersect returns a slice with all the entries of 'a' that are found in 'b'
func Intersect(a, b []string) []string {
	// complement diffFunc comparison
	return diffFunc(a, b, func(a []string, s string) bool { return Contains(a, s) })
}

// Repeat returns a slice consisting of 'n' copies of 's'
func Repeat(s string, n int) []string {
	a := make([]string, n)
	for i := range a {
		a[i] = s
	}
	return a
}

// Fill is an alias of Repeat (for PHP converts)
func Fill(n int, s string) []string {
	return Repeat(s, n)
}

// Replace returns a slice with the values of 'a' replaced with the index-matching
// values of 'b'. If 'b' has more entries than 'a' they will be appended.
func Replace(a, b []string) []string {
	m, n := len(a), len(b)
	c := make([]string, m)
	copy(c, a)
	for i := 0; i < m; i++ {
		if i >= n {
			break
		}
		c[i] = b[i]
	}
	for i := m; i < n; i++ {
		c = append(c, b[i])
	}
	return c
}

// Rand returns a slice with 'n' number of random entries of 'a'.
// Note: You may want initialize the rand seed once in your program.
//
//    rand.Seed(time.Now().UnixNano())
//
func Rand(a []string, n int) []string {
	m := len(a)
	b := make([]string, n)
	for i := range a {
		if i >= n {
			break
		}
		b[i] = a[rand.Intn(m)]
	}
	return b
}

// Shuffle returns a slice with randomized order of elements in 'a'.
// Note: You may want initialize the rand seed once in your program.
//
//    rand.Seed(time.Now().UnixNano())
//
func Shuffle(a []string) []string {
	m := len(a)
	b := make([]string, m)
	p := rand.Perm(m)
	for i := range a {
		b[i] = a[p[i]]
	}
	return b
}

// Reverse returns a slice of the reverse index order elements of 'a'.
func Reverse(a []string) []string {
	m := len(a)
	b := make([]string, m)
	for i := range a {
		b[i] = a[m-i-1]
	}
	return b
}

// Shift shifts the first element of '*a' and returns it, shortening the array by one.
// If '*a' is empty returns empty string "".
// Note that this function will change the array pointed by 'a'.
func Shift(a *[]string) string {
	if m := len(*a); m > 0 {
		s := (*a)[0]
		*a = (*a)[1:]
		return s
	}
	return ""
}

// Unshift prepends one or more elements to '*a' and returns the number of entries.
// Note that this function will change the array pointed by 'a'
func Unshift(a *[]string, s ...string) int {
	if s != nil {
		*a = append(s, *a...)
	}
	return len(*a)
}

// Pop removes the last element in '*a' and returns it, shortening the array by one.
// If '*a' is empty returns empty string "".
// Note that this function will change the array pointed by 'a'.
func Pop(a *[]string) string {
	if m := len(*a); m > 0 {
		s := (*a)[m-1]
		*a = (*a)[0 : m-1]
		return s
	}
	return ""
}

// Push appends one or more elements to '*a' and returns the number of entries.
// Note that this function will change the array pointed by 'a'.
func Push(a *[]string, s ...string) int {
	if s != nil {
		*a = append(*a, s...)
	}
	return len(*a)
}
