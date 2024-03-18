package main

import (
	"fmt"
	"math"
)

const (
	a int8  = 1
	b int16 = 2
)

type tType string

const (
	aa = 1
	ss = "go" // string and tType
)

const (
	ia = iota
	ib
	ic
	id
)

const (
	ja = iota << 1
	jb
	jc
	jd
)

const (
	_ = iota
	kb
	_
	kd
)

func main() {
	fmt.Println(
		math.MinInt32,
	)
}
