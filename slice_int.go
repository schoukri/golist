package golist

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
)

// SliceInt is a Slice of integers.
type SliceInt struct {
	data []int
}

// SliceIntFilterer defines the interface for Filter functions.
// type SliceIntFilterer interface {
// 	func(elem int) bool
// }

// SliceIntTransformer defines the interface for Transfor functions.
// type SliceIntTransformer interface {
// 	func(elem int) int
// }

// NewSliceInt returns a pointer to a SliceInt initialized with the specified elements.
func NewSliceInt(elems ...int) *SliceInt {
	s := newSliceInt()
	s.data = append(s.data, elems...)
	return s
}

func newSliceInt() *SliceInt {
	s := new(SliceInt)
	s.data = make([]int, 0)
	return s
}

// Append adds the elements to the end of SliceInt
func (s *SliceInt) Append(elems ...int) *SliceInt {
	if s == nil {
		return nil
	}
	s.data = append(s.data, elems...)
	return s
}

// Prepend adds the elements to the beginning of SliceInt
func (s *SliceInt) Prepend(elems ...int) *SliceInt {
	if s == nil {
		return nil
	}
	s.data = append(elems, s.data...)
	return s
}

// Insert inserts the value into the slice at the specified index.
func (s *SliceInt) Insert(index, value int) *SliceInt {
	// Grow the slice by one element.
	s.data = s.data[0 : len(s.data)+1]
	// Use copy to move the upper part of the slice out of the way and open a hole.
	copy(s.data[index+1:], s.data[index:])
	// Store the new value.
	s.data[index] = value
	// Return the result.
	return s
}

// Data returns the underlying slice of data.
func (s *SliceInt) Data() []int {
	if s == nil {
		return nil
	}
	return s.data
}

// Sort sorts the elements of SliceInt in increasing order.
func (s *SliceInt) Sort() *SliceInt {
	if s == nil {
		return nil
	}
	sort.Ints(s.data)
	return s
}

// Filter removes elements from SliceInt that do not satisfy the filter function.
func (s *SliceInt) Filter(fn func(elem int) bool) *SliceInt {
	if s == nil {
		return nil
	}
	data := s.data[:0]
	for _, elem := range s.data {
		if fn(elem) {
			data = append(data, elem)
		}
	}
	s.data = data
	return s
}

// Transform modifies each element of SliceInt according to the specified function.
func (s *SliceInt) Transform(fn func(elem int) int) *SliceInt {
	if s == nil {
		return nil
	}
	for i, elem := range s.data {
		s.data[i] = fn(elem)
	}
	return s
}

// Unique returns a SliceInt containing only unique elements.
func (s *SliceInt) Unique() *SliceInt {
	if s == nil {
		return nil
	}
	sort.Ints(s.data)
	j := 0
	for i := 1; i < len(s.data); i++ {
		if s.data[j] == s.data[i] {
			continue
		}
		j++
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	s.data = s.data[:j+1]

	return s
}

// Reverse returns a SliceInt with the elements in reverse order.
func (s *SliceInt) Reverse() *SliceInt {
	if s == nil {
		return nil
	}

	for i := len(s.data)/2 - 1; i >= 0; i-- {
		opp := len(s.data) - 1 - i
		s.data[i], s.data[opp] = s.data[opp], s.data[i]
	}

	return s
}

// Shuffle returns a SliceInt with the elements in random order.
func (s *SliceInt) Shuffle() *SliceInt {
	if s == nil {
		return nil
	}

	for i := len(s.data) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	return s
}

// Count returns the number of elements in SliceInt.
func (s *SliceInt) Count() int {
	return len(s.data)
}

// Min returns the smallest element in SliceInt (or an error if there aren't any elements).
func (s *SliceInt) Min() (int, error) {
	var min int
	if s.data == nil || len(s.data) == 0 {
		return min, errors.New("SliceInt does not contain any elements")
	}
	for _, elem := range s.data {
		if elem < min {
			min = elem
		}
	}
	return min, nil
}

// Max returns the largest element in SliceInt (or an error if there aren't any elements).
func (s *SliceInt) Max() (int, error) {
	var max int
	if s.data == nil || len(s.data) == 0 {
		return max, errors.New("SliceInt does not contain any elements")
	}
	for _, elem := range s.data {
		if elem > max {
			max = elem
		}
	}
	return max, nil
}

// At returns the element in SliceInt at the specified index (or an error if there aren't any elements).
func (s *SliceInt) At(index int) (int, error) {
	if s.data == nil || len(s.data) == 0 {
		return 0, errors.New("SliceInt does not contain any elements")
	}

	if index >= len(s.data) || index <= 0 {
		return 0, fmt.Errorf("index %d outside the range of SliceInt", index)
	}

	return s.data[index], nil
}

// Equal returns true if the SliceInt is equal to the specified SpliceInt.
func (s *SliceInt) Equal(x *SliceInt) bool {
	if s == nil || x == nil {
		return s == x
	}

	if len(s.data) != len(x.data) {
		return false
	}

	for i, elem := range s.data {
		if elem != x.data[i] {
			return false
		}
	}

	return true
}
