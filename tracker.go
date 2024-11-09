// https://roadmap.sh/projects/task-tracker
package main

import (
	"fmt"
	"task-tracker/json_task"
	"task-tracker/parser"
)

type Task = json_task.Task

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

	tsks := []Task{
		{"qwe", 123},
		{"rty", 456},
	}
	err := json_task.Dump("test.json", tsks)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(json_task.Read("test.json"))

	parser.AddCmd("add", AddTask)
	parser.AddCmd("update", UpdateTask)
	parser.AddCmd("delete", DeleteTask)
	parser.Parse()

}
