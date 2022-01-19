package main

import (
	"flag"
	"os"
	"fmt"
	"test3/Reminders/client"
)

var (
	backendURIFlag = flag.String("backend", "http://localhost:8080", "Backend URI")
	helpFlag = flag.Bool("help", false, "displays the help message to direct functionality")
)

func main() {
	flag.Parse()
	s := client.NewSwitch(*backendURIFlag)
	if *helpFlag || len(os.Args) == 1 {
		s.Help()
		return
	}

	err := s.Switch()
	if err != nil {
		fmt.Printf("command error: %v\n", err)
		os.Exit(2)
	}
}