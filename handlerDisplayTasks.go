package main

import (
	"errors"
	"fmt"
)

func listTasks(s *state, cmd command) error {
	if cmd.args == nil {
		for key, val := range s.unfinished {
			fmt.Printf("=====%s=====\n", key)
			for i, v := range val {
				fmt.Printf("%d. %s @ %v\n", i + 1, v.Name, v.Created)
			}
		}
		return nil
	}

	if len(cmd.args) > 1 {
		fmt.Println("the listTasks handler takes 1 optional argument, a list name. All other arguments will be ignored.")
	}
	
	val, ok := s.unfinished[cmd.args[0]]
	if !ok {
		return errors.New("supplied list does not exist")
	}

	fmt.Printf("=====%s=====\n", cmd.args[0])
	for i, v := range val {
		fmt.Printf("%d. %s @ %v\n", i + 1, v.Name, v.Created)
	}

	return nil
}

