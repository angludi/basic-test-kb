package command

import (
	"command-pattern/receiver"
)

type SumCommand struct {
	Calc receiver.ICalc
}

func (c *SumCommand) Execute() {
	c.Calc.Sum()
}
