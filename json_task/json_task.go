package json_task

import (
	"encoding/json"
	"os"
	"time"
)

type TaskStatus = string
type TaskID = int
type TaskTime = time.Time

const (
	ToDo       TaskStatus = "todo"
	InProgress TaskStatus = "in-progress"
	Done       TaskStatus = "done"
	None       TaskStatus = ""
)

type Task struct {
	ID          TaskID     `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   TaskTime   `json:"created-at"`
	UpdatedAt   TaskTime   `json:"updated-at"`
}

// Convert to json string
func Stringify(tasks []Task) (string, error) {
	js_bytes, err := json.MarshalIndent(tasks, "", "    ")
	return string(js_bytes), err
}

// Writeto json file
func WriteToFile(filename string, tasks []Task) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "    ")
	return encoder.Encode(tasks)
}

// Read tasks array from file
func ReadFile(filename string) (tasks []Task, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	js_bytes, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(js_bytes, &tasks)
	return
}
