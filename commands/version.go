package commands

import (
	"flag"
	"fmt"
)

func init() {
	Register(NewCmdVersion("version", versionIntro))
}

type CmdVersion struct {
	Cmd
	Flags struct {
		l bool
	}
}

func (cmd *CmdVersion) initFlagSet() {
	cmd.FlagSet = flag.NewFlagSet(cmd.Name, flag.ContinueOnError)
	cmd.FlagSet.BoolVar(&cmd.Flags.l, "l", false, versionFlagLUsage)
}

func NewCmdVersion(name, intro string) *CmdVersion {
	cmd := &CmdVersion{}
	cmd.Name = name
	cmd.Intro = intro
	cmd.initFlagSet()

	return cmd
}

func (cmd *CmdVersion) Run() {
	fmt.Printf(versionOutputTemplate, "0.1.0")
}
