package main

import (
	"flag"
	"fmt"
	"github.com/han-joker/apigo/commands"
	"os"
	"strings"
)

var mainFlagSet = flag.NewFlagSet(commands.MainCmdName, flag.ContinueOnError)

func init() {
	mainFlagSet.Usage = func() {
		fmt.Println()
		fmt.Print(commands.MainHelpMessage())
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Print(commands.MainHelpMessage())
		return
	}

	if err := mainFlagSet.Parse(os.Args[1:]); err != nil {
		return
	}

	commands.Run(strings.ToLower(os.Args[1]))
}
