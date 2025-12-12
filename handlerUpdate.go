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

	fmt.Println("Making Urgent...")

	for key, val := range s.unfinished {
		for i, v := range val {
			if time.Since(v.Created).Hours() >= float64(s.cfg.MakeUrgent) {
				fmt.Println(v.Name)
				switch(i) {
				case 0:
					s.unfinished[key] = val[1:]
				case len(val)-1:
					s.unfinished[key] = val[:len(val)-1]
				default:
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

	fmt.Println("...Removing...")

	for key, val := range s.finished {
		for i, v := range val {
			if time.Since(v.Created).Hours() >= float64(s.cfg.RemoveFromComplete) {
				fmt.Println(v.Name)
				switch(i) {
				case 0:
					s.finished[key] = val[1:]
				case len(val)-1:
					s.finished[key] = val[:len(val)-1]
				default:
					s.finished[key] = append(val[:i], val[i+1:]...)
				}
			}
		}
	}

	fmt.Println("...Finished")

	return nil
}

