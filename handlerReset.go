package main

import (
	"fmt"

	"github.com/JMitchell159/CLIToDo/internal/app"
)

func reset(s *state, cmd command) error {
	if cmd.args != nil {
		fmt.Println("The reset handler takes no arguments. All supplied arguments will be ignored.")
	}

	s.unfinished = app.CreateCollection()
	s.warn = app.CreateCollection()
	s.finished = app.CreateCollection()

	return nil
}

