// https://roadmap.sh/projects/task-tracker
package main

import (
	"fmt"
	"strconv"
	"task-tracker/json_task"
	"task-tracker/parser"
	"time"
)

type Task = json_task.Task
type TaskID = json_task.TaskID

const tasks_file = "tasks.json"

var (
	task_map = map[TaskID]Task{}
	last_id  TaskID
)

func find_task(str_id string) (*Task, int) {
	id, err := strconv.Atoi(str_id)
	if err != nil {
		fmt.Printf("Conversion error: [%v] %v\n", id, err)
		return nil, -1
	}
	tsk, ok := task_map[id]
	if !ok {
		fmt.Println("No task with id", id)
		return nil, -1
	}
	return &tsk, id
}

func require_args(args []string, min int) bool {
	if len(args) < min {
		fmt.Println("Not enough arguments!")
		return false
	}
	return true
}

func AddTask(args []string) {
	if !require_args(args, 1) {
		return
	}
	last_id++
	task_map[last_id] = Task{
		ID:          last_id,
		Description: args[0],
		Status:      json_task.ToDo,
		CreatedAt:   time.Now(),
	}
	fmt.Printf("Task added successfully (ID: %d)\n", last_id)
}

func UpdateTask(args []string) {
	if !require_args(args, 2) {
		return
	}
	tsk, id := find_task(args[0])
	if id < 0 {
		return
	}
	tsk.Description = args[1]
	tsk.UpdatedAt = time.Now()
	task_map[id] = *tsk
	fmt.Printf("Task updated successfully (ID: %d)\n", id)
}

func DeleteTask(args []string) {
	if !require_args(args, 1) {
		return
	}
	_, id := find_task(args[0])
	if id < 0 {
		return
	}
	delete(task_map, id)
	fmt.Printf("Task deleted successfully (ID: %d)\n", id)
}

func UpdateTaskStatus(args []string, status json_task.TaskStatus) {
	if !require_args(args, 1) {
		return
	}
	tsk, id := find_task(args[0])
	if id < 0 {
		return
	}
	tsk.Status = status
	tsk.UpdatedAt = time.Now()
	task_map[id] = *tsk
	fmt.Printf("Task marked as %v successfully (ID: %d)\n", status, id)
}

func MarkInProgress(args []string) {
	UpdateTaskStatus(args, json_task.InProgress)
}

func MarkDone(args []string) {
	UpdateTaskStatus(args, json_task.Done)
}

func List(args []string) {

	status_filter := json_task.None

	if len(args) > 0 {
		switch args[0] {
		case "done":
			status_filter = json_task.Done
		case "todo":
			status_filter = json_task.ToDo
		case "in-progress":
			status_filter = json_task.InProgress
		}
	}

	tasks := make([]Task, 0, len(task_map))

	for _, tsk := range task_map {
		if status_filter == json_task.None || status_filter == tsk.Status {
			tasks = append(tasks, tsk)
		}
	}
	task_str, err := json_task.Stringify(tasks)
	if err != nil {
		fmt.Printf("Error stringifying tasks %v\n", err)
	}
	fmt.Println(task_str)
}

func loadTasks() {
	tasks, _ := json_task.ReadFile(tasks_file)
	for _, tsk := range tasks {
		task_map[tsk.ID] = tsk
		last_id = max(last_id, tsk.ID)
	}
}

func saveTasks() {
	tasks := make([]Task, 0, len(task_map))
	for _, tsk := range task_map {
		tasks = append(tasks, tsk)
	}
	json_task.WriteToFile(tasks_file, tasks)
}

func main() {

	loadTasks()

	parser.AddCmd("add", AddTask)
	parser.AddCmd("update", UpdateTask)
	parser.AddCmd("delete", DeleteTask)
	parser.AddCmd("mark-in-progress", MarkInProgress)
	parser.AddCmd("mark-done", MarkDone)
	parser.AddCmd("list", List)
	parser.Parse()

	saveTasks()
}
