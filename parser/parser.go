package parser

import (
	"fmt"
	"os"
)

var commands = map[string]func([]string){}

func AddCmd(name string, action func([]string)) {
	commands[name] = action
}

func Parse() {

	if len(os.Args) > 1 {
		name := os.Args[1]
		cmd, ok := commands[name]
		if ok {
			cmd(os.Args[2:])
		} else {
			fmt.Println("Unknown command", name)
			os.Exit(1)
		}
	}
}
