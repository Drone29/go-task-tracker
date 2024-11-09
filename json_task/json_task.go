package json_task

import (
	"encoding/json"
	"os"
)

type Task struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// Convert to json string
func Stringify(tsk Task) ([]byte, error) {
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
