package command

import (
	"fmt"
)

const MainCmdName = "ApiGo"
const MainCmd = "apigo"

const commandNotExists = MainCmdName + " %s: unknown command\n"

var mainHelpMessage = MainCmdName + " is a scaffold for build APIs based on go frameworks." + "\n" +
	"Usage:" + "\n" +
	"\t" + MainCmd + " <command> [arguments]" + "\n" +
	"Examples:" + "\n" +
	"\t" + MainCmd + " build:api -conf api.json" + "\n"

func MainHelpMessage() string {
	if len(CmdSet) > 0 {
		mainHelpMessage += "The commands are:" + "\n"
	}
	for name, cmd := range CmdSet {
		intro := cmd.Elem().FieldByName("Intro").String()
		mainHelpMessage += fmt.Sprintf("\t%s\t\t%s\n", name, intro)
	}
	return mainHelpMessage
}

// version messages
const versionIntro = "print " + MainCmdName + " version"
const versionFlagLUsage = "lists the available versions"
const versionOutputTemplate = MainCmdName + " version: %s\n"
