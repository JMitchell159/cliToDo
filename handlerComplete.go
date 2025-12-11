package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	todolist "github.com/JMitchell159/CLIToDo/internal/toDoList"
)

func complete(s *state, cmd command) error {
	switch(len(cmd.args)) {
	case 0:
		return errors.New("the complete handler takes 3 arguments, the priority of the task, the name of the list, and whether or not it should ignore urgent tasks. This is the structure: (program name) complete <priority> <list_name> [ignore_urgent]")
	case 1:
		return errors.New("the complete handler takes 3 arguments, the priority of the task, the name of the list, and whether or not it should ignore urgent tasks. This is the structure: (program name) complete <priority> <list_name> [ignore_urgent]")
	case 2:
		val, ok := s.warn[cmd.args[1]]
		if !ok {
			return errors.New("provided list does not exist")
		}

		priority, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("something went wrong while trying to convert priority argument to number: %v", err)
		}
		if priority >= len(val) {
			return errors.New("priority out of range")
		}

		val2, ok := s.finished[cmd.args[1]]
		if !ok {
			list := todolist.CreateList()
			list = list.AddTask(val[priority].Name)
			err := s.finished.AddList(cmd.args[1], list)
			if err != nil {
				return err
			}
		} else {
			s.finished[cmd.args[1]] = val2.AddTask(val[priority].Name)
		}

		switch(priority) {
		case 0:
			s.warn[cmd.args[1]] = val[1:]
		case len(val)-1:
			s.warn[cmd.args[1]] = val[:len(val)-1]
		default:
			s.warn[cmd.args[1]] = append(val[:priority], val[priority+1:]...)
		}
	case 3:
		if strings.ToLower(cmd.args[2]) == "true" {
			val, ok := s.unfinished[cmd.args[1]]
			if !ok {
				return errors.New("provided list does not exist")
			}

			priority, err := strconv.Atoi(cmd.args[0])
			if err != nil {
				return fmt.Errorf("something went wrong while trying to convert priority argument to number: %v", err)
			}
			if priority >= len(val) {
				return errors.New("priority out of range")
			}

			val2, ok := s.finished[cmd.args[1]]
			if !ok {
				list := todolist.CreateList()
				list = list.AddTask(val[priority].Name)
				err := s.finished.AddList(cmd.args[1], list)
				if err != nil {
					return err
				}
			} else {
				s.finished[cmd.args[1]] = val2.AddTask(val[priority].Name)
			}

			switch(priority) {
			case 0:
				s.unfinished[cmd.args[1]] = val[1:]
			case len(val)-1:
				s.unfinished[cmd.args[1]] = val[:len(val)-1]
			default:
				s.unfinished[cmd.args[1]] = append(val[:priority], val[priority+1:]...)
			}
		} else {
			fmt.Println("If the [ignore_urgent] argument is anything other than true (case insensitive), it will be ignored.")
			val, ok := s.warn[cmd.args[1]]
			if !ok {
				return errors.New("provided list does not exist")
			}

			priority, err := strconv.Atoi(cmd.args[0])
			if err != nil {
				return fmt.Errorf("something went wrong while trying to convert priority argument to number: %v", err)
			}
			if priority >= len(val) {
				return errors.New("priority out of range")
			}

			val2, ok := s.finished[cmd.args[1]]
			if !ok {
				list := todolist.CreateList()
				list = list.AddTask(val[priority].Name)
				err := s.finished.AddList(cmd.args[1], list)
				if err != nil {
					return err
				}
			} else {
				s.finished[cmd.args[1]] = val2.AddTask(val[priority].Name)
			}

			switch(priority) {
			case 0:
				s.warn[cmd.args[1]] = val[1:]
			case len(val)-1:
				s.warn[cmd.args[1]] = val[:len(val)-1]
			default:
				s.warn[cmd.args[1]] = append(val[:priority], val[priority+1:]...)
			}
		}
	default:
		fmt.Println("The complete handler only takes 3 arguments. All other arguments will be ignored.")
		if strings.ToLower(cmd.args[2]) == "true" {
			val, ok := s.unfinished[cmd.args[1]]
			if !ok {
				return errors.New("provided list does not exist")
			}

			priority, err := strconv.Atoi(cmd.args[0])
			if err != nil {
				return fmt.Errorf("something went wrong while trying to convert priority argument to number: %v", err)
			}
			if priority >= len(val) {
				return errors.New("priority out of range")
			}

			val2, ok := s.finished[cmd.args[1]]
			if !ok {
				list := todolist.CreateList()
				list = list.AddTask(val[priority].Name)
				err := s.finished.AddList(cmd.args[1], list)
				if err != nil {
					return err
				}
			} else {
				s.finished[cmd.args[1]] = val2.AddTask(val[priority].Name)
			}

			switch(priority) {
			case 0:
				s.unfinished[cmd.args[1]] = val[1:]
			case len(val)-1:
				s.unfinished[cmd.args[1]] = val[:len(val)-1]
			default:
				s.unfinished[cmd.args[1]] = append(val[:priority], val[priority+1:]...)
			}
		} else {
			fmt.Println("If the [ignore_urgent] argument is anything other than true (case insensitive), it will be ignored.")
			val, ok := s.warn[cmd.args[1]]
			if !ok {
				return errors.New("provided list does not exist")
			}

			priority, err := strconv.Atoi(cmd.args[0])
			if err != nil {
				return fmt.Errorf("something went wrong while trying to convert priority argument to number: %v", err)
			}
			if priority >= len(val) {
				return errors.New("priority out of range")
			}

			val2, ok := s.finished[cmd.args[1]]
			if !ok {
				list := todolist.CreateList()
				list = list.AddTask(val[priority].Name)
				err := s.finished.AddList(cmd.args[1], list)
				if err != nil {
					return err
				}
			} else {
				s.finished[cmd.args[1]] = val2.AddTask(val[priority].Name)
			}

			switch(priority) {
			case 0:
				s.warn[cmd.args[1]] = val[1:]
			case len(val)-1:
				s.warn[cmd.args[1]] = val[:len(val)-1]
			default:
				s.warn[cmd.args[1]] = append(val[:priority], val[priority+1:]...)
			}
		}
	}

	return nil
}

