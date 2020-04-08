//go:generate go run cmd/golist/main.go -type=bool -output=SliceBool -sortable=false
//go:generate go run cmd/golist/main.go -type=byte -output=SliceByte
//go:generate go run cmd/golist/main.go -type=complex128 -output=SliceComplex128 -sortable=false
//go:generate go run cmd/golist/main.go -type=complex64 -output=SliceComplex64 -sortable=false
//go:generate go run cmd/golist/main.go -type=error -output=SliceError -sortable=false
//go:generate go run cmd/golist/main.go -type=float32 -output=SliceFloat32
//go:generate go run cmd/golist/main.go -type=float64 -output=SliceFloat64
//go:generate go run cmd/golist/main.go -type=int -output=SliceInt
//go:generate go run cmd/golist/main.go -type=int16 -output=SliceInt16
//go:generate go run cmd/golist/main.go -type=int32 -output=SliceInt32
//go:generate go run cmd/golist/main.go -type=int64 -output=SliceInt64
//go:generate go run cmd/golist/main.go -type=int8 -output=SliceInt8
//go:generate go run cmd/golist/main.go -type=rune -output=SliceRune
//go:generate go run cmd/golist/main.go -type=string -output=SliceString
//go:generate go run cmd/golist/main.go -type=uint -output=SliceUint
//go:generate go run cmd/golist/main.go -type=uint16 -output=SliceUint16
//go:generate go run cmd/golist/main.go -type=uint32 -output=SliceUint32
//go:generate go run cmd/golist/main.go -type=uint64 -output=SliceUint64
//go:generate go run cmd/golist/main.go -type=uint8 -output=SliceUint8
//go:generate go run cmd/golist/main.go -type=uintptr -output=SliceUintptr

package golist
