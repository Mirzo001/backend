package main

import (
	"bigint/bigint"
	"fmt"
)

func main() {

	a, err := bigint.NewInt("11231231231231231231231231233")
	if err != nil {
		panic(err)
	}
	b, err := bigint.NewInt("1")
	if err != nil {
		panic(err)
	}

	c := bigint.Add(a, b)
	fmt.Println(c)
}
