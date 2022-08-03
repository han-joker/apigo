package commands

import (
	"flag"
	"fmt"
	"os"
	"reflect"
)

const nameFieldName = "Name"
const flagSetFieldName = "FlagSet"
const runMethodName = "Run"

type Cmder interface {
	Run()
}

type Cmd struct {
	Name    string
	Intro   string
	FlagSet *flag.FlagSet
}

// CmdSet 命令集合
var CmdSet = make(map[string]reflect.Value)

func Register(cmder Cmder) {
	v := reflect.ValueOf(cmder)
	name := v.Elem().FieldByName(nameFieldName).String()
	CmdSet[name] = v
}

func Run(name string) {
	cmd, exists := CmdSet[name]
	if !exists {
		fmt.Printf(commandNotExists, name)
		fmt.Println()
		fmt.Print(MainHelpMessage())
		return
	}

	// parse flags
	flagSetField := cmd.Elem().FieldByName(flagSetFieldName)
	if !flagSetField.IsZero() {
		args := reflect.ValueOf(os.Args[2:])
		parseMethod := flagSetField.MethodByName("Parse")
		parseMethod.Call([]reflect.Value{args})
		//fmt.Println(cmd.Elem().FieldByName("Flags"))
	}

	cmd.MethodByName(runMethodName).Call(nil)
}
