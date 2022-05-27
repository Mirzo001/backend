package main

import (
	"bigint/bigint"
	"fmt"
)

func main() {
	// a, err := bigint.NewInt("988847123412385995937737458959")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(a)
	a, err := bigint.NewInt("13")
	if err != nil {
		panic(err)
	}
	b, err := bigint.NewInt("12212121222232322222")
	if err != nil {
		panic(err)
	}
	// err = b.Set("1") // b = "1"
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(a)
	// a.Set("321")
	// fmt.Println(a)
	// fmt.Println(b)
	c := bigint.Add(a, b)
	fmt.Println(c)
}
