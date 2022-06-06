package bigint_test

import (
	"bigint/bigint"
	"fmt"
	"testing"
)

// func TestAddBasic(t *testing.T) {
// 	a, err := bigint.NewInt("123")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	b, err := bigint.NewInt("456")
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	want, err := bigint.NewInt("579")
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	ans := bigint.Add(a, b)
// 	if ans != want {
// 		t.Errorf("got %v, want %v", ans, want)
// 	}

// }

func TestAddTD(t *testing.T) {
	tests := []struct {
		lable string
		a, b  string
		want  string
		err   error
	}{
		{"Test 1", "123", "456", "579", nil},
	}

	for _, tt := range tests {
		t.Run(tt.lable, func(t *testing.T) {
			var err error
			var a, b bigint.BigInt

			a, err = bigint.NewInt(tt.a)

			b, err = bigint.NewInt(tt.b)

			testCaseTextOut := fmt.Sprintf("Add(%s,%s)", a.Value(), b.Value())

			if err != nil {
				if tt.err == nil {
					t.Errorf("Test: %s | Result: %s | Expected: %s", testCaseTextOut, err.Error(), "nil")
				} else if err.Error() != tt.err.Error() {
					t.Errorf("Test: %s | Result: %s | Expected: %s", testCaseTextOut, err.Error(), tt.err.Error())
				}
			} else {
				ans := bigint.Add(a, b)
				if ans.Value() != tt.want {
					t.Errorf("Test: %s | Result: %s | Expected: %s", testCaseTextOut, ans.Value(), tt.want)
				}
			}

		})
	}
}
