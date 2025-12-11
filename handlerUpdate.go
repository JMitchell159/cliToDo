package main

import (
	"fmt"
	"time"

	todolist "github.com/JMitchell159/CLIToDo/internal/toDoList"
)

func update(s *state, cmd command) error {
	if cmd.args != nil {
		fmt.Println("The update handler does not take any arguments, all supplied arguments will be ignored.")
	}

	for key, val := range s.unfinished {
		for i, v := range val {
			if time.Since(v.Created).Minutes() >= 5 {
				if i == 0 {
					s.unfinished[key] = val[i+1:]
				} else if i == len(val) - 1 {
					s.unfinished[key] = val[:i]
				} else {
					s.unfinished[key] = append(val[:i], val[i+1:]...)
				}
				val2, ok := s.warn[key]
				if !ok {
					list := todolist.CreateList()
					list = list.AddTask(v.Name)
					err := s.warn.AddList(key, list)
					if err != nil {
						return err
					}
				} else {
					s.warn[key] = val2.AddTask(v.Name)
				}
			}
		}
	}

	return nil
}

