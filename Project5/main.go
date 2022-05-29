package main

import (
	"bigint/bigint"
	"fmt"
)

func main() {

	a, err := bigint.NewInt("123456789101112131415")
	if err != nil {
		panic(err)
	}
	b, err := bigint.NewInt("123456789101112131414")
	if err != nil {
		panic(err)
	}

	c := bigint.Add(a, b)
	_ = c
	d := bigint.Sub(a, b)
	_ = d
	e := bigint.Multiply(a, b)

	fmt.Println(e)
}
