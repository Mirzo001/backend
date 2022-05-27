package bigint

import (
	"errors"
	"fmt"
	"math/big"
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
	i := new(big.Int)
	i, ok := i.SetString(a.value, 10)
	j := new(big.Int)
	j, ok1 := j.SetString(b.value, 10)
	if !ok && !ok1 {
		fmt.Println("SetString: error")
	}
	i.Add(i, j)
	bigstr := i.String()

	return BigInt{bigstr}
}
