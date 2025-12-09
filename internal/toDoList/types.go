package todolist

import (
	"errors"
	"fmt"
	"time"
)

type Task struct {
	Name string `json:"name"`
	Created time.Time `json:"created"`
}

type TaskList []Task

func CreateList() TaskList {
	return make(TaskList, 0, 8)
}

func (list TaskList) AddTask(task string) TaskList {
	return append(list, Task{Name: task, Created: time.Now()})
}

func (list TaskList) promoteTask(priority int) (TaskList, error) {
	if priority < 0 {
		return list, errors.New("cannot take a negative priority")
	}
	if priority == 0 {
		return list, errors.New("task is already at highest priority")
	}
	list[priority], list[priority-1] = list[priority-1], list[priority]
	return list, nil
}

func (list TaskList) demoteTask(priority int) (TaskList, error) {
	if priority > len(list)-1 {
		return list, errors.New("priority does not exist")
	}
	if priority == len(list)-1 {
		return list, errors.New("task is already at lowest priority")
	}
	list[priority], list[priority+1] = list[priority+1], list[priority]
	return list, nil
}

func (list TaskList) MoveTask(old, new int) (TaskList, error) {
	if new > len(list)-1 {
		return list, errors.New("priority does not exist")
	}
	if new < 0 {
		return list, errors.New("cannot take a negative priority")
	}
	current := old
	for current < new {
		list, err := list.demoteTask(current)
		if err != nil {
			return list, fmt.Errorf("error while demoting task: %v", err)
		}
		current++
	}
	for current > new {
		list, err := list.promoteTask(current)
		if err != nil {
			return list, fmt.Errorf("error while promoting task: %v", err)
		}
		current--
	}
	return list, nil
}

