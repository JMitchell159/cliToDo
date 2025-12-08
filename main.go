package main

import (
	"fmt"
	"log"

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

	fmt.Println(s)
}

