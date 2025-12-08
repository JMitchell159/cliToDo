package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	handler map[string]func(command) error
}

func (c *commands) run(cmd command) error {
	handle, ok := c.handler[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}

	err := handle(cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(command) error) {
	if c.handler == nil {
		c.handler = make(map[string]func(command) error)
	}

	c.handler[name] = f
}

