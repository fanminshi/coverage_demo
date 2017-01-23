package op

import (
	"testing"
)

type Test struct {
	op      string
	oprand1 int
	oprand2 int
	exp     int
}

var tests = []Test{
	{"-", 2, 1, 1}, // 2 - 1 = 1?
}

func TestOp(t *testing.T) {
	for i, test := range tests {
		result := Op(test.op, test.oprand1, test.oprand2)
		if result != test.exp {
			t.Errorf("#%d: %d %s %d =%d; want %d", i, test.oprand1, test.op, test.oprand2, result, test.exp)
		}
	}
}
