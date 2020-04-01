package command

import (
	"command-pattern/receiver"
)

type MultiplyCommand struct {
	Calc receiver.ICalc
}

func (c *MultiplyCommand) Execute() {
	c.Calc.Multiply()
}
