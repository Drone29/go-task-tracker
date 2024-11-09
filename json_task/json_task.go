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
func Stringify(tsk Task) (string, error) {
	js_bytes, err := json.MarshalIndent(tsk, "", "    ")
	return string(js_bytes), err
}

func ToBytes(tsk Task) ([]byte, error) {
	return json.Marshal(tsk)
}

// Restore from json string
func Restore(encoded []byte) (Task, error) {
	var tsk Task
	err := json.Unmarshal(encoded, &tsk)
	return tsk, err
}

// Dump to json file
func Dump(filename string, tasks []Task) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	encoder := json.NewEncoder(f)
	return encoder.Encode(tasks)
}

// Read tasks array from file
func Read(filename string) (tasks []Task, err error) {
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
