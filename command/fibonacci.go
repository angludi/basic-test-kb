package command

import (
	"command-pattern/receiver"
)

type FibonacciCommand struct {
	Calc receiver.ICalc
}

func (c *FibonacciCommand) Execute() {
	c.Calc.Fibonacci()
}
