package golist_test

import (
	"fmt"

	"github.com/schoukri/golist"
)

func ExampleSliceInt_Insert() {
	s := golist.NewSliceInt(10, 40)
	s.Insert(1, 20, 30)

	for _, value := range s.Data() {
		fmt.Println(value)
	}

	// Output:
	// 10
	// 20
	// 30
	// 40
}

func ExampleSliceInt_Count() {
	s := golist.NewSliceInt(-5, 3, 2, 21, 0, 40, -4, -20, 100, 40, -9, 18, 21, -33, 40)

	count := s.Filter(func(a int) bool {
		return a > 10
	}).Count()

	uniqCount := s.Filter(func(a int) bool {
		return a > 10
	}).Unique().Count()

	fmt.Println("total count:", count)
	fmt.Println("unique count:", uniqCount)

	// Output:
	// total count: 7
	// unique count: 4
}

func ExampleSliceInt_Transform() {
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
}
