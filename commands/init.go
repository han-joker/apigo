package commands

import (
	"flag"
	"fmt"
)

func init() {
	Register(NewCmdInit("init", initIntro))
}

type CmdInit struct {
	Cmd
	Flags struct {
	}
}

func (cmd *CmdInit) initFlagSet() {
	cmd.FlagSet = flag.NewFlagSet(cmd.Name, flag.ContinueOnError)
}

func NewCmdInit(name, intro string) *CmdInit {
	cmd := &CmdInit{}
	cmd.Name = name
	cmd.Intro = intro
	cmd.initFlagSet()

	return cmd
}

func (cmd *CmdInit) Run() {
	fmt.Println("init")
}
