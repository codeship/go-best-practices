package fib

import (
	"fmt"
	"testing"
)

type testCase struct {
	in       int // input
	expected int // expected result
}

func newCase(in, expected int) testCase { return testCase{in, expected} }

func (c testCase) Name() string {
	return fmt.Sprintf("Fib(%d) should equal %d", c.in, c.expected)
}

func (c testCase) Run(t *testing.T) {
	actual := Fib(c.in)
	if actual != c.expected {
		t.Errorf("Fib(%d): expected %d, actual %d", c.in, c.expected, actual)
	}
}

func TestFib(t *testing.T) {
	for _, testCase := range []testCase{
		newCase(1, 1),
		newCase(2, 1),
		newCase(3, 2),
		newCase(4, 3),
		newCase(5, 5),
		newCase(6, 8),
		newCase(7, 13),
	} {
		t.Run(testCase.Name(), testCase.Run)
	}
}
