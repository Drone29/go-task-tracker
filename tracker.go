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

const dump_file = "dump.json"

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

func AddTask(args []string) {
	if len(args) < 1 {
		fmt.Println("Not enough arguments!")
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
	if len(args) < 2 {
		fmt.Println("Not enough arguments!")
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
	if len(args) < 1 {
		fmt.Println("Not enough arguments!")
		return
	}
	_, id := find_task(args[0])
	if id < 0 {
		return
	}
	delete(task_map, id)
	fmt.Printf("Task deleted successfully (ID: %d)\n", id)
}

func MarkInProgress(args []string) {
	if len(args) < 1 {
		fmt.Println("Not enough arguments!")
		return
	}
	tsk, id := find_task(args[0])
	if id < 0 {
		return
	}
	tsk.Status = json_task.InProgress
	tsk.UpdatedAt = time.Now()
	task_map[id] = *tsk
	fmt.Printf("Task marked as in progress successfully (ID: %d)\n", id)
}

func MarkDone(args []string) {
	if len(args) < 1 {
		fmt.Println("Not enough arguments!")
		return
	}
	tsk, id := find_task(args[0])
	if id < 0 {
		return
	}
	tsk.Status = json_task.Done
	tsk.UpdatedAt = time.Now()
	task_map[id] = *tsk
	fmt.Printf("Task marked as done successfully (ID: %d)\n", id)
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
		if status_filter != json_task.None && status_filter != tsk.Status {
			continue
		}
		tasks = append(tasks, tsk)
	}
	task_str, _ := json_task.Stringify(tasks)
	fmt.Println(task_str)
}

func loadTasks() {
	tasks, _ := json_task.Read(dump_file)
	for _, tsk := range tasks {
		task_map[tsk.ID] = tsk
		if tsk.ID > last_id {
			last_id = tsk.ID
		}
	}
}

func saveTasks() {
	tasks := make([]Task, 0, len(task_map))
	for _, tsk := range task_map {
		tasks = append(tasks, tsk)
	}
	json_task.Dump(dump_file, tasks)
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
