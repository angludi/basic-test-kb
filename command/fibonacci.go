package command

import (
	"github.com/angludi/tech-test-kb/receiver"
)

type FibonacciCommand struct {
	Calc receiver.ICalc
}

func (c *FibonacciCommand) Execute() {
	c.Calc.Fibonacci()
}
