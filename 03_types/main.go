package main

import (
	"fmt"
	"math"
)

type Human1 struct {
	Age    uint32
	Name   string
	Weight *Params
}

type Human2 struct {
	Age    uint32
	Name   string
	Weight Params
}

type Human3 struct {
	Age  uint32
	Name string
	Params
}

type Human4 struct {
	Age    uint32
	Name   string
	Weight float32
	Params
}

type Human5 struct {
	Age  uint32
	Name string
	Params
	Weight float32
}

type Human6 struct {
	Age  uint32
	Name string
	Params
	Weight float64
}

type Params struct {
	Weight float32
}

func main() {
	var h3 Human3
	h3.Weight = 1
	h3.Params.Weight = 2

	fmt.Println(h3.Weight, h3.Params.Weight)

	var h4 Human4
	h4.Weight = 1
	h4.Params.Weight = 2

	fmt.Println(h4.Weight, h4.Params.Weight)

	var h5 Human5
	h5.Weight = 1
	h5.Params.Weight = 2

	fmt.Println(h5.Weight, h5.Params.Weight)

	var h6 Human6
	h6.Weight = math.MaxFloat64
	h6.Params.Weight = math.MaxFloat64

	fmt.Println(h5.Weight, h5.Params.Weight)
}
