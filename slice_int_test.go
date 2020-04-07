package golist

import (
	"reflect"
	"testing"
)

func TestNewSliceInt(t *testing.T) {
	type args struct {
		elems []int
	}
	tests := []struct {
		name string
		args args
		want *SliceInt
	}{
		{
			name: "new nil SliceInt",
			args: args{elems: nil},
			want: &SliceInt{data: []int{}},
		},
		{
			name: "new empty SliceInt",
			args: args{elems: []int{}},
			want: &SliceInt{data: []int{}},
		},
		{
			name: "new populated SliceInt",
			args: args{elems: []int{2, 4, 3}},
			want: &SliceInt{data: []int{2, 4, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSliceInt(tt.args.elems...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSliceInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceInt_Append(t *testing.T) {
	type args struct {
		elems []int
	}
	tests := []struct {
		name  string
		start *SliceInt
		args  args
		want  *SliceInt
	}{
		{
			name:  "append one",
			start: NewSliceInt(2, 3, 4),
			args:  args{elems: []int{5}},
			want:  NewSliceInt(2, 3, 4, 5),
		},
		{
			name:  "append lots",
			start: NewSliceInt(2, 3, 4),
			args:  args{elems: []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}},
			want:  NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29),
		},
		{
			name:  "append to nil SliceInt",
			start: nil,
			args:  args{elems: []int{}},
			want:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.start.Append(tt.args.elems...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceInt.Append() = %v, want %v", got, tt.want)
			}
			if got != tt.start {
				t.Errorf("SliceInt.Append() pointer not the same: got = %v, start = %v", got, tt.start)
			}
		})
	}
}

func TestSliceInt_Prepend(t *testing.T) {
	type args struct {
		elems []int
	}
	tests := []struct {
		name  string
		start *SliceInt
		args  args
		want  *SliceInt
	}{
		{
			name:  "prepend one",
			start: NewSliceInt(2, 3, 4),
			args:  args{elems: []int{5}},
			want:  NewSliceInt(5, 2, 3, 4),
		},
		{
			name:  "prepend lots",
			start: NewSliceInt(2, 3, 4),
			args:  args{elems: []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}},
			want:  NewSliceInt(5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 2, 3, 4),
		},
		{
			name:  "prepend to nil SliceInt",
			start: nil,
			args:  args{elems: []int{}},
			want:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.start.Prepend(tt.args.elems...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceInt.Prepend() = %v, want %v", got, tt.want)
			}
			if got != tt.start {
				t.Errorf("SliceInt.Prepend() pointer not the same: got = %v, start = %v", got, tt.start)
			}
		})
	}
}

func TestSliceInt_At(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name      string
		start     *SliceInt
		args      args
		want      int
		wantPanic bool
	}{
		{
			name:      "At() called on nil SliceInt receiver",
			start:     nil,
			wantPanic: true,
		},
		{
			name:      "At() called on empty SliceInt receiver",
			start:     NewSliceInt(),
			wantPanic: true,
		},
		{
			name:      "At() called with negative index",
			start:     NewSliceInt(2, 3, 4),
			args:      args{index: -1},
			wantPanic: true,
		},
		{
			name:      "At() called with too large index",
			start:     NewSliceInt(2, 3, 4),
			args:      args{index: 3},
			wantPanic: true,
		},
		{
			name:  "At() first element",
			start: NewSliceInt(2, 3, 4),
			args:  args{index: 0},
			want:  2,
		},
		{
			name:  "At() middle element",
			start: NewSliceInt(2, 3, 4),
			args:  args{index: 1},
			want:  3,
		},
		{
			name:  "At() last element",
			start: NewSliceInt(2, 3, 4),
			args:  args{index: 2},
			want:  4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("SliceInt.At() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			got := tt.start.At(tt.args.index)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceInt.At() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceInt_Set(t *testing.T) {
	type args struct {
		index int
		elem  int
	}
	tests := []struct {
		name  string
		start *SliceInt
		args  args
		want  *SliceInt
	}{
		{
			name:  "Set on nil SliceInt receiver",
			start: nil,
			want:  nil,
		},
		{
			name:  "set first",
			start: NewSliceInt(2, 3, 4),
			args:  args{index: 0, elem: 5},
			want:  NewSliceInt(5, 3, 4),
		},
		{
			name:  "set middle",
			start: NewSliceInt(2, 3, 4),
			args:  args{index: 1, elem: 5},
			want:  NewSliceInt(2, 5, 4),
		},
		{
			name:  "set last",
			start: NewSliceInt(2, 3, 4),
			args:  args{index: 2, elem: 5},
			want:  NewSliceInt(2, 3, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.start.Set(tt.args.index, tt.args.elem)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceInt.Set() = %v, want %v", got, tt.want)
			}
			if got != tt.start {
				t.Errorf("SliceInt.Set() pointer not the same: got = %v, start = %v", got, tt.start)
			}
		})
	}
}

func TestSliceInt_Insert(t *testing.T) {
	type args struct {
		index int
		elems []int
	}
	tests := []struct {
		name  string
		start *SliceInt
		args  args
		want  *SliceInt
	}{
		{
			name:  "insert one at the beginning",
			start: NewSliceInt(2, 3, 4),
			args:  args{index: 0, elems: []int{5}},
			want:  NewSliceInt(5, 2, 3, 4),
		},
		{
			name:  "insert many at the beginning",
			start: NewSliceInt(2, 3, 4),
			args:  args{index: 0, elems: []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}},
			want:  NewSliceInt(5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 2, 3, 4),
		},
		{
			name:  "insert one in the middle",
			start: NewSliceInt(2, 3, 4),
			args:  args{index: 2, elems: []int{5}},
			want:  NewSliceInt(2, 3, 5, 4),
		},
		{
			name:  "insert many in the middle",
			start: NewSliceInt(2, 3, 4),
			args:  args{index: 2, elems: []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}},
			want:  NewSliceInt(2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 4),
		},
		{
			name:  "insert one at the end",
			start: NewSliceInt(2, 3, 4),
			args:  args{index: 0, elems: []int{5}},
			want:  NewSliceInt(5, 2, 3, 4),
		},
		{
			name:  "insert many at the end",
			start: NewSliceInt(2, 3, 4),
			args:  args{index: 3, elems: []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}},
			want:  NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29),
		},
		{
			name:  "insert into nil SliceInt",
			start: nil,
			args:  args{index: 0, elems: []int{}},
			want:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.start.Insert(tt.args.index, tt.args.elems...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceInt.Insert() = %v, want %v", got, tt.want)
			}
			if got != tt.start {
				t.Errorf("SliceInt.Insert() pointer not the same: got = %v, start = %v", got, tt.start)
			}
		})
	}
}

func TestSliceInt_Remove(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name  string
		start *SliceInt
		args  args
		want  *SliceInt
	}{
		{
			name:  "remove one at the beginning",
			start: NewSliceInt(2, 3, 4),
			args:  args{index: 0},
			want:  NewSliceInt(3, 4),
		},
		{
			name:  "remove one in the middle",
			start: NewSliceInt(2, 3, 4),
			args:  args{index: 1},
			want:  NewSliceInt(2, 4),
		},
		{
			name:  "remove one at the end",
			start: NewSliceInt(2, 3, 4),
			args:  args{index: 2},
			want:  NewSliceInt(2, 3),
		},
		{
			name:  "remove into nil SliceInt",
			start: nil,
			want:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.start.Remove(tt.args.index)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceInt.Remove() = %v, want %v", got, tt.want)
			}
			if got != tt.start {
				t.Errorf("SliceInt.Remove() pointer not the same: got = %v, start = %v", got, tt.start)
			}
		})
	}
}

func TestSliceInt_Filter(t *testing.T) {
	type args struct {
		fn func(elem int) bool
	}
	tests := []struct {
		name  string
		start *SliceInt
		args  args
		want  *SliceInt
	}{
		{
			name:  "filter nil SliceInt",
			start: nil,
			args:  args{fn: func(a int) bool { return a > 0 }},
			want:  nil,
		},
		{
			name:  "keep even numbers (sorted)",
			start: NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13),
			args:  args{fn: func(a int) bool { return a%2 == 0 }},
			want:  NewSliceInt(2, 4, 6, 8, 10, 12),
		},
		{
			name:  "keep numbers > 5 (sorted)",
			start: NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13),
			args:  args{fn: func(a int) bool { return a > 5 }},
			want:  NewSliceInt(6, 7, 8, 9, 10, 11, 12, 13),
		},
		{
			name:  "keep even numbers (not sorted)",
			start: NewSliceInt(13, 10, 5, 8, 2, 11, 6, 12, 4, 9, 3, 7),
			args:  args{fn: func(a int) bool { return a%2 == 0 }},
			want:  NewSliceInt(10, 8, 2, 6, 12, 4),
		},
		{
			name:  "keep numbers > 5 (not sorted)",
			start: NewSliceInt(13, 10, 5, 8, 2, 11, 6, 12, 4, 9, 3, 7),
			args:  args{fn: func(a int) bool { return a > 5 }},
			want:  NewSliceInt(13, 10, 8, 11, 6, 12, 9, 7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.start.Filter(tt.args.fn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceInt.Filter() = %v, want %v", got, tt.want)
			}
			if got != tt.start {
				t.Errorf("SliceInt.Filter() pointer not the same: got = %v, start = %v", got, tt.start)
			}
		})
	}
}

func TestSliceInt_Transform(t *testing.T) {
	type args struct {
		fn func(elem int) int
	}
	tests := []struct {
		name  string
		start *SliceInt
		args  args
		want  *SliceInt
	}{
		{
			name:  "transform nil SliceInt",
			start: nil,
			args:  args{fn: func(a int) int { return a * 2 }},
			want:  nil,
		},
		{
			name:  "add 3",
			start: NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13),
			args:  args{fn: func(a int) int { return a + 3 }},
			want:  NewSliceInt(5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16),
		},
		{
			name:  "mod 3",
			start: NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13),
			args:  args{fn: func(a int) int { return a % 3 }},
			want:  NewSliceInt(2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.start.Transform(tt.args.fn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceInt.Transform() = %v, want %v", got, tt.want)
			}
			if got != tt.start {
				t.Errorf("SliceInt.Transform() pointer not the same: got = %v, start = %v", got, tt.start)
			}
		})
	}
}

func TestSliceInt_Unique(t *testing.T) {
	tests := []struct {
		name  string
		start *SliceInt
		want  *SliceInt
	}{
		{
			name:  "unique nil SliceInt",
			start: nil,
			want:  nil,
		},
		{
			name:  "unique numbers (already unique, sorted)",
			start: NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13),
			want:  NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13),
		},
		{
			name:  "unique numbers (already unique, not sorted)",
			start: NewSliceInt(13, 10, 5, 8, 2, 11, 6, 12, 4, 9, 3, 7),
			want:  NewSliceInt(13, 10, 5, 8, 2, 11, 6, 12, 4, 9, 3, 7),
		},
		{
			name:  "unique numbers (sorted)",
			start: NewSliceInt(2, 3, 3, 3, 4, 4, 5, 6, 6, 6, 6, 7, 7, 8, 9, 10, 11, 11, 12, 13, 13, 13),
			want:  NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13),
		},
		{
			name:  "unique numbers (not sorted)",
			start: NewSliceInt(13, 13, 10, 5, 5, 5, 10, 8, 2, 11, 11, 2, 13, 5, 6, 6, 12, 4, 12, 11, 9, 12, 3, 7, 3, 7, 3, 7, 9, 9, 10, 11, 11),
			want:  NewSliceInt(13, 10, 5, 8, 2, 11, 6, 12, 4, 9, 3, 7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.start.Unique()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceInt.Unique() = %v, want %v", got, tt.want)
			}
			if got != tt.start {
				t.Errorf("SliceInt.Unique() pointer not the same: got = %v, start = %v", got, tt.start)
			}
		})
	}
}

func TestSliceInt_Sort(t *testing.T) {
	tests := []struct {
		name  string
		start *SliceInt
		want  *SliceInt
	}{
		{
			name:  "sort nil SliceInt",
			start: nil,
			want:  nil,
		},
		{
			name:  "sort numbers (already sorted, unique values",
			start: NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13),
			want:  NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13),
		},
		{
			name:  "sort numbers (unique values)",
			start: NewSliceInt(13, 10, 5, 8, 2, 11, 6, 12, 4, 9, 3, 7),
			want:  NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13),
		},
		{
			name:  "sort numbers (already sorted, with duplicates)",
			start: NewSliceInt(2, 3, 3, 3, 4, 4, 5, 6, 6, 6, 6, 7, 7, 8, 9, 10, 11, 11, 12, 13, 13, 13),
			want:  NewSliceInt(2, 3, 3, 3, 4, 4, 5, 6, 6, 6, 6, 7, 7, 8, 9, 10, 11, 11, 12, 13, 13, 13),
		},
		{
			name:  "sort numbers (with duplicates)",
			start: NewSliceInt(13, 13, 10, 5, 5, 5, 10, 8, 2, 11, 11, 2, 13, 5, 6, 6, 12, 4, 12, 11, 9, 12, 3, 7, 3, 7, 3, 7, 9, 9, 10, 11, 11),
			want:  NewSliceInt(2, 2, 3, 3, 3, 4, 5, 5, 5, 5, 6, 6, 7, 7, 7, 8, 9, 9, 9, 10, 10, 10, 11, 11, 11, 11, 11, 12, 12, 12, 13, 13, 13),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.start.Sort()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceInt.Sort() = %v, want %v", got, tt.want)
			}
			if got != tt.start {
				t.Errorf("SliceInt.Sort() pointer not the same: got = %v, start = %v", got, tt.start)
			}
		})
	}
}

func TestSliceInt_Reverse(t *testing.T) {
	tests := []struct {
		name  string
		start *SliceInt
		want  *SliceInt
	}{
		{
			name:  "reverse nil SliceInt",
			start: nil,
			want:  nil,
		},
		{
			name:  "reverse numbers (sorted, unique values",
			start: NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13),
			want:  NewSliceInt(13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2),
		},
		{
			name:  "reverse numbers (not sorted, unique values)",
			start: NewSliceInt(13, 10, 5, 8, 2, 11, 6, 12, 4, 9, 3, 7),
			want:  NewSliceInt(7, 3, 9, 4, 12, 6, 11, 2, 8, 5, 10, 13),
		},
		{
			name:  "reverse numbers (sorted, with duplicates)",
			start: NewSliceInt(2, 3, 3, 3, 4, 4, 5, 6, 6, 6, 6, 7, 7, 8, 9, 10, 11, 11, 12, 13, 13, 13),
			want:  NewSliceInt(13, 13, 13, 12, 11, 11, 10, 9, 8, 7, 7, 6, 6, 6, 6, 5, 4, 4, 3, 3, 3, 2),
		},
		{
			name:  "reverse numbers (not sorted, with duplicates)",
			start: NewSliceInt(13, 13, 10, 5, 5, 5, 10, 8, 2, 11, 11, 2, 13, 5, 6, 6, 12, 4, 12, 11, 9, 12, 3, 7, 3, 7, 3, 7, 9, 9, 10, 11, 11),
			want:  NewSliceInt(11, 11, 10, 9, 9, 7, 3, 7, 3, 7, 3, 12, 9, 11, 12, 4, 12, 6, 6, 5, 13, 2, 11, 11, 2, 8, 10, 5, 5, 5, 10, 13, 13),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.start.Reverse()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceInt.Reverse() = %v, want %v", got, tt.want)
			}
			if got != tt.start {
				t.Errorf("SliceInt.Reverse() pointer not the same: got = %v, start = %v", got, tt.start)
			}
		})
	}
}

func TestSliceInt_Shuffle(t *testing.T) {
	type args struct {
		seed int64
	}
	tests := []struct {
		name  string
		start *SliceInt
		args  args
		want  *SliceInt
	}{
		{
			name:  "shuffle nil SliceInt",
			start: nil,
			want:  nil,
		},
		{
			name:  "shuffle empty SliceInt",
			start: NewSliceInt(),
			args:  args{seed: 0}, // seed == 0 means use nanoseconds since time epoch
			want:  NewSliceInt(),
		},
		{
			name:  "shuffle one SliceInt",
			start: NewSliceInt(2),
			args:  args{seed: 0}, // seed == 0 means use nanoseconds since time epoch
			want:  NewSliceInt(2),
		},
		{
			name:  "shuffle numbers (sorted, unique values)",
			start: NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13),
			args:  args{seed: 1},
			want:  NewSliceInt(4, 3, 13, 11, 7, 2, 6, 10, 5, 8, 12, 9),
		},
		{
			name:  "shuffle numbers (sorted, unique values) -- same seed, same order",
			start: NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13),
			args:  args{seed: 1},
			want:  NewSliceInt(4, 3, 13, 11, 7, 2, 6, 10, 5, 8, 12, 9),
		},
		{
			name:  "shuffle numbers (sorted, unique values) -- different seed, different order",
			start: NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13),
			args:  args{seed: 2},
			want:  NewSliceInt(7, 9, 5, 11, 10, 12, 8, 6, 3, 2, 13, 4),
		},
		{
			name:  "shuffle numbers (not sorted, unique values)",
			start: NewSliceInt(13, 10, 5, 8, 2, 11, 6, 12, 4, 9, 3, 7),
			args:  args{seed: 3},
			want:  NewSliceInt(2, 13, 8, 7, 11, 5, 10, 3, 6, 9, 12, 4),
		},
		{
			name:  "shuffle numbers (sorted, with duplicates)",
			start: NewSliceInt(2, 3, 3, 3, 4, 4, 5, 6, 6, 6, 6, 7, 7, 8, 9, 10, 11, 11, 12, 13, 13, 13),
			args:  args{seed: 4},
			want:  NewSliceInt(3, 7, 7, 3, 13, 6, 11, 13, 6, 2, 11, 4, 9, 10, 6, 8, 6, 13, 12, 5, 3, 4),
		},
		{
			name:  "shuffle numbers (not sorted, with duplicates)",
			start: NewSliceInt(13, 13, 10, 5, 5, 5, 10, 8, 2, 11, 11, 2, 13, 5, 6, 6, 12, 4, 12, 11, 9, 12, 3, 7, 3, 7, 3, 7, 9, 9, 10, 11, 11),
			args:  args{seed: 5},
			want:  NewSliceInt(9, 12, 7, 9, 5, 5, 11, 6, 2, 12, 9, 3, 7, 11, 13, 10, 11, 5, 11, 10, 2, 6, 13, 7, 8, 3, 13, 11, 5, 4, 10, 12, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.start.Shuffle(tt.args.seed)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceInt.Shuffle() = %v, want %v", got, tt.want)
			}
			if got != tt.start {
				t.Errorf("SliceInt.Shuffle() pointer not the same: got = %v, start = %v", got, tt.start)
			}
		})
	}
}

func TestSliceInt_Data(t *testing.T) {
	tests := []struct {
		name  string
		start *SliceInt
		want  []int
	}{
		{
			name:  "data nil SliceInt",
			start: nil,
			want:  nil,
		},
		{
			name:  "data numbers (sorted, unique values",
			start: NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13),
			want:  []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
		},
		{
			name:  "data numbers (not sorted, unique values)",
			start: NewSliceInt(13, 10, 5, 8, 2, 11, 6, 12, 4, 9, 3, 7),
			want:  []int{13, 10, 5, 8, 2, 11, 6, 12, 4, 9, 3, 7},
		},
		{
			name:  "data numbers (sorted, with duplicates)",
			start: NewSliceInt(2, 3, 3, 3, 4, 4, 5, 6, 6, 6, 6, 7, 7, 8, 9, 10, 11, 11, 12, 13, 13, 13),
			want:  []int{2, 3, 3, 3, 4, 4, 5, 6, 6, 6, 6, 7, 7, 8, 9, 10, 11, 11, 12, 13, 13, 13},
		},
		{
			name:  "data numbers (not sorted, with duplicates)",
			start: NewSliceInt(13, 13, 10, 5, 5, 5, 10, 8, 2, 11, 11, 2, 13, 5, 6, 6, 12, 4, 12, 11, 9, 12, 3, 7, 3, 7, 3, 7, 9, 9, 10, 11, 11),
			want:  []int{13, 13, 10, 5, 5, 5, 10, 8, 2, 11, 11, 2, 13, 5, 6, 6, 12, 4, 12, 11, 9, 12, 3, 7, 3, 7, 3, 7, 9, 9, 10, 11, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.start.Data()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceInt.Data() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceInt_Len(t *testing.T) {
	tests := []struct {
		name  string
		start *SliceInt
		want  int
	}{
		{
			name:  "data numbers (sorted, unique values",
			start: NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13),
			want:  12,
		},
		{
			name:  "data numbers (not sorted, unique values)",
			start: NewSliceInt(13, 10, 5, 8, 2, 11, 6, 12, 4, 9, 3, 7),
			want:  12,
		},
		{
			name:  "data numbers (sorted, with duplicates)",
			start: NewSliceInt(2, 3, 3, 3, 4, 4, 5, 6, 6, 6, 6, 7, 7, 8, 9, 10, 11, 11, 12, 13, 13, 13),
			want:  22,
		},
		{
			name:  "data numbers (not sorted, with duplicates)",
			start: NewSliceInt(13, 13, 10, 5, 5, 5, 10, 8, 2, 11, 11, 2, 13, 5, 6, 6, 12, 4, 12, 11, 9, 12, 3, 7, 3, 7, 3, 7, 9, 9, 10, 11, 11),
			want:  33,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.start.Len()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceInt.Len() = %v, want %v", got, tt.want)
			}
			count := tt.start.Count()
			if !reflect.DeepEqual(got, count) {
				t.Errorf("SliceInt.Len() = %v, SliceInt.Count() = %v", got, count)
			}

		})
	}
}

func TestSliceInt_Min(t *testing.T) {
	tests := []struct {
		name      string
		start     *SliceInt
		want      int
		wantPanic bool
	}{
		{
			name:      "Min() called on nil SliceInt receiver",
			start:     nil,
			wantPanic: true,
		},
		{
			name:      "Min() called on empty SliceInt receiver",
			start:     NewSliceInt(),
			wantPanic: true,
		},
		{
			name:  "one number",
			start: NewSliceInt(13),
			want:  13,
		},
		{
			name:  "data numbers (sorted, unique values)",
			start: NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13),
			want:  2,
		},
		{
			name:  "data numbers (not sorted, unique values)",
			start: NewSliceInt(13, 10, 5, 8, 2, 11, 6, 12, 4, 9, 3, 7),
			want:  2,
		},
		{
			name:  "data numbers (sorted, with duplicates)",
			start: NewSliceInt(2, 3, 3, 3, 4, 4, 5, 6, 6, 6, 6, 7, 7, 8, 9, 10, 11, 11, 12, 13, 13, 13),
			want:  2,
		},
		{
			name:  "data numbers (not sorted, with duplicates)",
			start: NewSliceInt(13, 13, 10, 5, 5, 5, 10, 8, 2, 11, 11, 2, 13, 5, 6, 6, 12, 4, 12, 11, 9, 12, 3, 7, 3, 7, 3, 7, 9, 9, 10, 11, 11),
			want:  2,
		},
		{
			name:  "smallest negative number",
			start: NewSliceInt(-2, -3, -4, -5, -6, -7, -8, -9, -10, -11, -12, -13),
			want:  -13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("SliceInt.Min() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			got := tt.start.Min()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceInt.Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceInt_Max(t *testing.T) {
	tests := []struct {
		name      string
		start     *SliceInt
		want      int
		wantPanic bool
	}{
		{
			name:      "Max() called on nil SliceInt receiver",
			start:     nil,
			wantPanic: true,
		},
		{
			name:      "Max() called on empty SliceInt receiver",
			start:     NewSliceInt(),
			wantPanic: true,
		},
		{
			name:  "one number",
			start: NewSliceInt(13),
			want:  13,
		},
		{
			name:  "data numbers (sorted, unique values)",
			start: NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13),
			want:  13,
		},
		{
			name:  "data numbers (not sorted, unique values)",
			start: NewSliceInt(13, 10, 5, 8, 2, 11, 6, 12, 4, 9, 3, 7),
			want:  13,
		},
		{
			name:  "data numbers (sorted, with duplicates)",
			start: NewSliceInt(2, 3, 3, 3, 4, 4, 5, 6, 6, 6, 6, 7, 7, 8, 9, 10, 11, 11, 12, 13, 13, 13),
			want:  13,
		},
		{
			name:  "data numbers (not sorted, with duplicates)",
			start: NewSliceInt(13, 13, 10, 5, 5, 5, 10, 8, 2, 11, 11, 2, 13, 5, 6, 6, 12, 4, 12, 11, 9, 12, 3, 7, 3, 7, 3, 7, 9, 9, 10, 11, 11),
			want:  13,
		},
		{
			name:  "largest negative number",
			start: NewSliceInt(-2, -3, -4, -5, -6, -7, -8, -9, -10, -11, -12, -13),
			want:  -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("SliceInt.Max() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			got := tt.start.Max()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceInt.Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceInt_Equal(t *testing.T) {
	type args struct {
		s2 *SliceInt
	}
	tests := []struct {
		name  string
		start *SliceInt
		args  args
		want  bool
	}{
		{
			name:  "nil SliceInt receiver equal nil input SliceInt",
			start: nil,
			args:  args{s2: nil},
			want:  true,
		},
		{
			name:  "nil SliceInt receiver equal non-nil input SliceInt",
			start: nil,
			args:  args{s2: NewSliceInt(1)},
			want:  false,
		},
		{
			name:  "non-nil SliceInt receiver equal nil input SliceInt",
			start: NewSliceInt(1),
			args:  args{s2: nil},
			want:  false,
		},
		{
			name:  "SliceInts have different lengths",
			start: NewSliceInt(1, 2),
			args:  args{s2: NewSliceInt(1)},
			want:  false,
		},
		{
			name:  "SliceInts have different elements",
			start: NewSliceInt(1, 2),
			args:  args{s2: NewSliceInt(1, 3)},
			want:  false,
		},
		{
			name:  "SliceInts have the same elements in a different order",
			start: NewSliceInt(1, 2),
			args:  args{s2: NewSliceInt(2, 1)},
			want:  false,
		},
		{
			name:  "SliceInts have the same elements in the same order",
			start: NewSliceInt(1, 2, 3),
			args:  args{s2: NewSliceInt(1, 2, 3)},
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.start.Equal(tt.args.s2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceInt.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceInt_Clone(t *testing.T) {
	tests := []struct {
		name  string
		start *SliceInt
		want  *SliceInt
	}{
		{
			name:  "clone nil SliceInt",
			start: nil,
			want:  nil,
		},
		{
			name:  "clone empty SliceInt",
			start: NewSliceInt(),
			want:  NewSliceInt(),
		},
		{
			name:  "clone numbers",
			start: NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13),
			want:  NewSliceInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.start.Clone()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceInt.Clone = %v, want %v", got, tt.want)
			}
			if got != nil && got == tt.start {
				t.Errorf("SliceInt.Clone pointer is the same: got = %v, start = %v", got, tt.start)
			}
		})
	}
}
