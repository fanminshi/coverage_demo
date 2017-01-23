package op

import (
	"log"
	"playground/coverage_demo/add"
	"playground/coverage_demo/sub"
)

func Op(op string, op1 int, op2 int) int {
	switch {
	case op == "+":
		return add.Add(op1, op2)
	case op == "-":
		return sub.Sub(op1, op2)
	default:
		log.Panicf("Unknown Operation %v", op)
	}
	// never reach here
	return -1
}
