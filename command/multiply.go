package command

import (
	"github.com/angludi/tech-test-kb/receiver"
)

type MultiplyCommand struct {
	Calc receiver.ICalc
}

func (c *MultiplyCommand) Execute() {
	c.Calc.Multiply()
}
