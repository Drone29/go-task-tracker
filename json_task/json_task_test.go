package json_task

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

var sample_task = Task{
	ID:          12,
	Description: "Test",
	Status:      ToDo,
	CreatedAt:   time.Date(2022, 7, 25, 14, 30, 0, 0, time.UTC),
}

func TestStringify(t *testing.T) {
	taskstr, err := Stringify([]Task{sample_task})
	if err != nil {
		t.Errorf("Error stringifying task [%v] %v", sample_task, err)
	}
	var new_tasks []Task
	if err = json.Unmarshal([]byte(taskstr), &new_tasks); err != nil {
		t.Errorf("Error unmarshaling string %v %v", taskstr, err)
	}
	if new_tasks[0] != sample_task {
		t.Errorf("Results do not match! old: %v new: %v", sample_task, new_tasks[0])
	}
}

func TestWriteRead(t *testing.T) {
	if err := WriteToFile("test.json", []Task{sample_task}); err != nil {
		t.Errorf("Error writing to file %v", sample_task)
	}
	new_tasks, err := ReadFile("test.json")
	if err != nil {
		t.Errorf("Error reading from file %v", err)
	}
	if new_tasks[0] != sample_task {
		t.Errorf("Results do not match! old: %v new: %v", sample_task, new_tasks[0])
	}
	os.Remove("test.json")
}
