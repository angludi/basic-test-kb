package command

import (
	"github.com/angludi/command-pattern/receiver"
)

type FibonacciCommand struct {
	Calc receiver.ICalc
}

func (c *FibonacciCommand) Execute() {
	c.Calc.Fibonacci()
}
