package add

import (
	"testing"
)

type Test struct {
	op1 int
	op2 int
	exp int
}

var tests = []Test{
	{1, 1, 2}, // 1 + 1 = 2
	{2, 3, 5}, // 2 + 3 = 5
}

func TestAdd(t *testing.T) {
	for i, test := range tests {
		result := Add(test.op1, test.op2)
		if result != test.exp {
			t.Errorf("#%d: add(%d, %d)=%s; want %s", i, test.op1, test.op2, result, test.exp)
		}
	}
}
