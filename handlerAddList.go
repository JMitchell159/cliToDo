package main

import (
	"errors"
	"fmt"

	todolist "github.com/JMitchell159/CLIToDo/internal/toDoList"
)

func addList(s *state, cmd command) error {
	if cmd.args == nil {
		return errors.New("the addList handler takes 1 argument, the name of the list")
	}
	if len(cmd.args) > 1 {
		fmt.Println("The addList handler only takes 1 argument, the name of the list. All other arguments will be ignored.")
	}

	list := todolist.CreateList()
	s.unfinished[cmd.args[0]] = list

	return nil
}

