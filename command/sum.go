package command

import (
	"github.com/angludi/tech-test-kb/receiver"
)

type SumCommand struct {
	Calc receiver.ICalc
}

func (c *SumCommand) Execute() {
	c.Calc.Sum()
}
