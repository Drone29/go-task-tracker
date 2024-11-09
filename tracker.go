// https://roadmap.sh/projects/task-tracker
package main

import (
	"fmt"
	"task-tracker/json_task"
	"task-tracker/parser"
)

type Task = json_task.Task

const dump_file = "dump.json"

var task_map = map[string]Task{}

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
	// read tasks stored in file
	tasks, _ := json_task.Read(dump_file)
	// populate map from array
	for _, tsk := range tasks {
		task_map[tsk.Description] = tsk
	}

	parser.AddCmd("add", AddTask)
	parser.AddCmd("update", UpdateTask)
	parser.AddCmd("delete", DeleteTask)
	parser.Parse()

}
