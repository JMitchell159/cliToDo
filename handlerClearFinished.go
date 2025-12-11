package main

import "github.com/JMitchell159/CLIToDo/internal/app"

func clearComplete(s *state, cmd command) error {
	s.finished = app.CreateCollection()
	return nil
}

