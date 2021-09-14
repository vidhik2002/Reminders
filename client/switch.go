package client

import (
	"fmt"
	"os"
)

type BackendHTTPClient interface {
}

func NewSwitch(uri string) Switch {
	httpClient := NewHTTPClient(uri)
	s := Switch{
		client:        httpClient,
		backendAPIURL: uri,
	}
	s.commands = map[string]func() func(string) error{
		"create": s.create,
		"edit":   s.edit,
		"fetch":  s.fetch,
		"delete": s.delete,
		"health": s.health,
	}
	return s
}
type Switch struct {
	client        BackendHTTPClient
	backendAPIURL string
	commands      map[string]func() func(string) error
}
func (s Switch) Switch() error{
	cmdNameofCommand := os.Args[1]
	cmd, ok := s.commands[cmdNameofCommand]
	if !ok{
		return fmt.Errorf("invalid command - '%s'", cmdNameofCommand)
	}
	return cmd()(cmdNameofCommand) 
}

func (s Switch) create() func(string) error {
	return func(cmd string) error {
		fmt.Println("create reminder")
		return nil
	}
}

func (s Switch) edit() func(string) error {
	return func(cmd string) error {
		fmt.Println("edit reminder")
		return nil
	}
}
func (s Switch) fetch() func(string) error {
	return func(cmd string) error {
		fmt.Println("fetch reminder")
		return nil
	}
}
func (s Switch) delete() func(string) error {
	return func(cmd string) error {
		fmt.Println("delete reminder")
		return nil
	}
}
func (s Switch) health() func(string) error {
	return func(cmd string) error {
		fmt.Println("health reminder")
		return nil
	}
}
