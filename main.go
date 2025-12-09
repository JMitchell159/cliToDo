package main

import (
	"log"
	"os"

	"github.com/JMitchell159/CLIToDo/internal/app"
)

type state struct {
	unfinished app.ListCollection
	warn app.ListCollection
	finished app.ListCollection
}

func main() {
	u, err := app.Load("unfinished")
	if err != nil {
		log.Fatalf("error while loading unfinished tasks: %v", err)
	}
	w, err := app.Load("warn")
	if err != nil {
		log.Fatalf("error while loading warn tasks: %v", err)
	}
	f, err := app.Load("finished")
	if err != nil {
		log.Fatalf("error while loading finished tasks: %v", err)
	}

	s := state{
		unfinished: u,
		warn: w,
		finished: f,
	}

	cmds := commands{
		handler: make(map[string]func(*state, command) error),
	}

	cmds.register("addTask", addTask)
	cmds.register("addList", addList)
	cmds.register("listTasks", listTasks)
	cmds.register("update", makeUrgent)
	inputs := os.Args
	if len(inputs) < 2 {
		log.Fatal("not enough arguments provided")
	}

	cmd := command{
		name: inputs[1],
		args: nil,
	}
	if len(inputs) > 2 {
		cmd.args = inputs[2:]
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatal(err)
	}

	s.unfinished.Save("unfinished")
	s.warn.Save("warn")
	s.finished.Save("finished")
}

