# Golist

[![Build Status](https://travis-ci.com/schoukri/golist.svg?branch=master)](https://travis-ci.com/schoukri/golist)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](./LICENSE)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/schoukri/golist)
[![Go Report Card](https://goreportcard.com/badge/github.com/schoukri/golist?style=flat-square)](https://goreportcard.com/report/github.com/schoukri/golist)

Golist provides useful functions for common list operations in Go.

## Introduction

Working with lists of data in Go using the built-in slice data type and operators can sometimes be a pain. Simple and common operations such as sorting, removing duplicates, and manipulating data can often requires a lot of boilerplate code to accomplish. This package simplifies working with slices by providing easy to use methods for common operations. The methods can also be chained together to enable complex operations on the underlying data with a lot less code, but still highly-readable and easy to maintain.

## Example

Here is an example that starts with a slice of integers, multiplies each item by 2, selects only those items that are greater than 20, removes any duplicates, then finally sorts the integers in ascending order.

```go
	s := golist.NewSliceInt(-5, 3, 2, 21, 0, 40, -4, -20, 100, 40, -9, 18, 21, -33, 40)

	s.Transform(func(a int) int {
		return a * 2
	}).Filter(func(a int) bool {
		return a > 20
	}).Unique().Sort()

	for _, value := range s.Data() {
		fmt.Println(value)
	}

	// Output:
	// 36
	// 42
	// 80
	// 200
```

## Inspiration

* Go [SliceTricks](https://github.com/golang/go/wiki/SliceTricks) examples and sample code
* C# [Enumerable Class](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable) LINQ Methods
* Mark Phelps' [Optional](https://github.com/markphelps/optional) package for using `go generate` to create custom wrappers for built-in types.


## Warning

Golist is in active development and is not yet ready for use in production. The API is experimental and is subject to change without notice. Suggestions or contributions to improve Golist and help shape its design are welcome and encouraged!
