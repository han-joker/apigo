package main

import (
	"flag"
	"fmt"
	"github.com/han-joker/apigo/command"
	"os"
	"strings"
)

var apigoFlagSet = flag.NewFlagSet(command.MainCmdName, flag.ContinueOnError)

func init() {
	apigoFlagSet.Usage = func() {
		fmt.Println()
		fmt.Print(command.MainHelpMessage())
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Print(command.MainHelpMessage())
		return
	}

	if err := apigoFlagSet.Parse(os.Args[1:]); err != nil {
		return
	}

	command.Run(strings.ToLower(os.Args[1]))
}
