package main

import (
	"fmt"
	"os"
)

type TCommand interface {
	Run(arg ...string)
}

type HelpCommand struct {
}

func (*HelpCommand) Run(arg ...string) {
	fmt.Println(os.Args[0] + " command")
}

func main() {
	Commands := map[string]TCommand{
		"envadd":    &EnvaddCommand{},
		"compile":   &CompileCommand{},
		"uninstall": &UninstallCommand{},
	}

	if len(os.Args) < 2 {
		panic("Please supply a command!")
	} else {
		if command, ok := Commands[os.Args[1]]; ok {
			//fmt.Println("running " + os.Args[1])
			command.Run(os.Args[1:]...)
		} else {
			panic(os.Args[1] + " command is not defined!")
		}
	}

}
