package main

import (
	"errors"
	"fmt"

	todolist "github.com/JMitchell159/CLIToDo/internal/toDoList"
)

func addTask(s *state, cmd command) error {
	switch(len(cmd.args)) {
	case 0:
		return errors.New("the addTask handler takes at least 1 argument, the name of the task")
	case 1:
		switch(len(s.unfinished)) {
		case 0:
			list := todolist.CreateList()
			list = list.AddTask(cmd.args[0])
			s.unfinished.AddList("list1", list)
		case 1:
			for key, val := range s.unfinished {
				s.unfinished[key] = val.AddTask(cmd.args[0])
			}
		default:
			return errors.New("the addTask handler must be supplied with a 2nd argument, the name of the list, if more than one list already exists")
		}
	case 2:
		val, ok := s.unfinished[cmd.args[1]]
		if !ok {
			return errors.New("supplied list does not exist")
		}
		s.unfinished[cmd.args[1]] = val.AddTask(cmd.args[0])
	default:
		fmt.Println("The addTask handler takes 2 arguments, the name of the task & the name of the list (optional). All other arguments will be ignored.")
		val, ok := s.unfinished[cmd.args[1]]
		if !ok {
			return errors.New("supplied list does not exist")
		}
		s.unfinished[cmd.args[1]] = val.AddTask(cmd.args[0])
	}

	return nil
}

