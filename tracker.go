// https://roadmap.sh/projects/task-tracker
package main

import (
	"fmt"
	"task-tracker/parser"
)

func AddTask(args []string) {
	fmt.Println("Add", args[0])
}

func UpdateTask(args []string) {
	fmt.Println("Update", args[0], args[1])
}

func DeleteTask(args []string) {
	fmt.Println("Delete", args[0])
}

func main() {

	parser.AddCmd("add", AddTask)
	parser.AddCmd("update", UpdateTask)
	parser.AddCmd("delete", DeleteTask)
	parser.Parse()

}
