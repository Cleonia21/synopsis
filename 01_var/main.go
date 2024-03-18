package main

import "fmt"

var (
	a int8  = 1
	b int16 = 2
	c int32 = 3
	d int64 = 4

	ua uint8  = 1
	ub uint16 = 2
	uc uint32 = 3
	ud uint64 = 4

	f32 float32 = 23.3
	f64 float64 = 46.3

	bb byte = 8 // uint8

	r rune = 'M' // uint32

	s string = "golang"

	isBool bool = true

	co64  complex64  = 1 + 2i
	co128 complex128 = 1 + 3i
)

func main() {
	{ // ptr
		goVersion := "1.20"

		var ptr *string
		ptr = &goVersion
		_ = ptr
	}

	{ // interface
		var unknownType interface{} = 2009
		fmt.Println(unknownType)
		unknownType = "2009"
		fmt.Println(unknownType)
	}

}
