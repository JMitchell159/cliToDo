package main

import "fmt"

func listComplete(s *state, cmd command) error {
	if len(s.finished) == 0 {
		fmt.Println("You have no completed tasks. Complete some tasks first to see them here.")
	} else {
		fmt.Println("=====Completed=====")
		for key, val := range s.finished {
			fmt.Printf("-----%s-----\n", key)
			for i, v := range val {
				fmt.Printf("%d. %s @ %v\n", i+1, v.Name, v.Created)
			}
		}
	}

	return nil
}

