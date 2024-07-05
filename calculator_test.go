package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestCalc(t *testing.T) {
	var tests = []struct {
		a, b    float64
		op      string
		want    float64
		calcErr error
	}{
		{3, 2, "+", 5, nil},
		{-9, -2, "+", -11, nil},
		{-2, 6, "*", -12, nil},
		{1.23, 6.5, "*", 7.995, nil},
		{5, 5, "+", 10, nil},
		{-8, 3, "+", -5, nil},
		{4, 9, "*", 36, nil},
		{-2.5, 5, "*", -12.5, nil},
		{8, 0, "/", 0, errors.New("division by zero")},
		{10, 2, "/", 5, nil},
		{15, 3, "/", 5, nil},
		{7.5, 1.5, "/", 5, nil},
		{666, 13, "+-", 0, errors.New("unknown operator")},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%v %s %v", test.a, test.op, test.b)
		t.Run(testName, func(t *testing.T) {
			res, err := calc(test.a, test.b, test.op)
			if res != test.want {
				t.Errorf("got res:%v err:%s, want res:%v err:%s", res, err, test.want, test.calcErr)
			}
		})
	}
}
