package main

import (
	"errors"
	"fmt"
)

func listTasks(s *state, cmd command) error {
	if len(s.warn) + len(s.unfinished) == 0 {
		fmt.Println("You have no tasks recorded. Use the addTask handler to add some tasks to your to do list.")
	}
	if cmd.args == nil {
		if len(s.warn) > 0 {
			fmt.Println("=====URGENT=====")
			for key, val := range s.warn {
				fmt.Printf("-----%s-----\n", key)
				for i, v := range val {
					fmt.Printf("%d. %s @ %v\n", i + 1, v.Name, v.Created)
				}
			}
			fmt.Println()
		}
		if len(s.unfinished) > 0 {
			fmt.Println("=====Unfinished=====")
			for key, val := range s.unfinished {
				fmt.Printf("-----%s-----\n", key)
				for i, v := range val {
					fmt.Printf("%d. %s @ %v\n", i + 1, v.Name, v.Created)
				}
			}
		}
		return nil
	}

	if len(cmd.args) > 1 {
		fmt.Println("the listTasks handler takes 1 optional argument, a list name. All other arguments will be ignored.")
	}

	val, ok := s.warn[cmd.args[0]]
	val2, ok2 := s.unfinished[cmd.args[0]]
	if !ok && !ok2 {
		return errors.New("supplied list does not exist")
	}

	if ok {
		fmt.Println("=====URGENT=====")
		fmt.Printf("-----%s-----\n", cmd.args[0])
		for i, v := range val {
			fmt.Printf("%d. %s @ %v\n", i + 1, v.Name, v.Created)
		}
		fmt.Println()
	}
	if ok2 {
		fmt.Println("=====Unfinished=====")
		fmt.Printf("-----%s-----\n", cmd.args[0])
		for i, v := range val2 {
			fmt.Printf("%d. %s @ %v\n", i+1, v.Name, v.Created)
		}
	}

	return nil
}

