// https://roadmap.sh/projects/task-tracker
package main

import (
	"fmt"
	"task-tracker/json_task"
	"task-tracker/parser"
)

type Task = json_task.Task
type TaskID = json_task.TaskID

const dump_file = "dump.json"

var task_map = map[TaskID]Task{}
var last_id TaskID

func AddTask(args []string) {
	fmt.Println("Add", args[0])
}

func UpdateTask(args []string) {
	fmt.Println("Update", args[0], args[1])
}

func DeleteTask(args []string) {
	fmt.Println("Delete", args[0])
}

func MarkInProgress(args []string) {

}

func MarkDone(args []string) {

}

func List(args []string) {

}

func main() {
	// read tasks stored in file
	tasks, _ := json_task.Read(dump_file)
	// populate map from array, update last id
	for _, tsk := range tasks {
		task_map[tsk.ID] = tsk
		last_id = max(last_id, tsk.ID)
	}

	parser.AddCmd("add", AddTask)
	parser.AddCmd("update", UpdateTask)
	parser.AddCmd("delete", DeleteTask)
	parser.AddCmd("mark-in-progress", MarkInProgress)
	parser.AddCmd("mark-done", MarkDone)
	parser.AddCmd("list", List)
	parser.Parse()

}
