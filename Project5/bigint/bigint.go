package bigint

import (
	"errors"
	"strconv"
	"strings"
)

type BigInt struct {
	value string
}

var BadInput = errors.New("Bad Input")

func validate(num *string) error {
	allowed := "123456789"
	str := *num
	for i := 0; i < len(str); i++ {
		if !strings.Contains(allowed, string(str[i])) {
			return BadInput
		}
	}

	return nil
}

func NewInt(num string) (BigInt, error) {
	if err := validate(&num); err != nil {
		return BigInt{}, err
	}
	return BigInt{value: num}, nil
}

func (z *BigInt) Set(num string) error {
	if err := validate(&num); err != nil {
		return err
	}

	z.value = num
	return nil
}

func Add(a, b BigInt) BigInt {
	// val := a.value + b.value
	a1, err := strconv.Atoi(a.value)
	b1, err := strconv.Atoi(b.value)
	if err != nil {
		panic(err)
	}
	val := a1 + b1
	str := strconv.Itoa(val)
	return BigInt{value: str}
	// return BigInt{
}
