# Task Tracker

## Description

This tool lets you add, update, delete and change status for tasks  
Tasks' properties and status are stored in a `tasks.json` file  
located in the same directory as the executable

## Usage

### Add Task
```sh
./task-tracker add "Task description"
# Output: Task added successfully (ID: 1)
```
### Update Task Description
```sh
./task-tracker update 1 "New description"
# Output: Task updated successfully (ID: 1)
```
### Delete Task
```sh
./task-tracker delete 1
# Output: Task deleted successfully (ID: 1)
```
### Update Task Status
```sh
./task-tracker mark-in-progress 1
# Output: Task marked as mark-in-progress successfully (ID: 1)
./task-tracker mark-done 1
# Output: Task marked as done successfully (ID: 1)
```

### List Tasks
#### List All
```sh
./task-tracker list
# Output:
# [
# {
#     "id": 1,
#     "description": "New description",
#     "status": "done",
#     "created-at": "2024-11-10T10:22:32.080947938+01:00",
#     "updated-at": "2024-11-10T10:22:32.080947938+01:00"
# }
# ]
```
#### List With Filter By Status
```sh
./task-tracker list todo
# Output:
# []
```
```sh
./task-tracker list in-progress
# Output:
# []
```
```sh
./task-tracker list done
# Output:
# [
# {
#     "id": 1,
#     "description": "New description",
#     "status": "done",
#     "created-at": "2024-11-10T10:22:32.080947938+01:00",
#     "updated-at": "2024-11-10T10:22:32.080947938+01:00"
# }
# ]
```

## Build And Install

To build and install, use `go build` and `go install` respectively, from the project's root directory 
```sh
go build
```
```sh
go install
```

## Testing

To test, run `go test` with recursive search from the root directory
```sh
go test ./...
```

## Roadmap link
https://roadmap.sh/projects/task-tracker