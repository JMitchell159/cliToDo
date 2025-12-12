package main

import (
	"log"
	"os"

	"github.com/JMitchell159/CLIToDo/internal/app"
	"github.com/JMitchell159/CLIToDo/internal/config"
)

type state struct {
	unfinished app.ListCollection
	warn app.ListCollection
	finished app.ListCollection
	cfg *config.Config
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
	c, err := config.Load()
	if err != nil {
		log.Fatalf("error while loading config: %v", err)
	}

	s := state{
		unfinished: u,
		warn: w,
		finished: f,
		cfg: c,
	}

	cmds := commands{
		handler: make(map[string]func(*state, command) error),
	}

	cmds.register("addTask", addTask)
	cmds.register("addList", addList)
	cmds.register("clearComplete", clearComplete)
	cmds.register("complete", complete)
	cmds.register("listComplete", listComplete)
	cmds.register("listTasks", listTasks)
	cmds.register("reset", reset)
	cmds.register("update", update)
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

