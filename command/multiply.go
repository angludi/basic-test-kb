package command

import (
	"github.com/angludi/command-pattern/receiver"
)

type MultiplyCommand struct {
	Calc receiver.ICalc
}

func (c *MultiplyCommand) Execute() {
	c.Calc.Multiply()
}
