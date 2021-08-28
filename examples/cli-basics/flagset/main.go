package main

import (
	"fmt"
	"os"
	"flag"
	"log"
)

func main() {
	// fmt.Println("hi")
	// fmt.Println(os.Args)
	if len(os.Args) < 2 {
		fmt.Println("no command found")
		os.Exit(2)
	}
	cmd := os.Args[1]
	switch cmd {
	case "greet":
		greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
		msgFlag := greetCmd.String("msg", "Reminder", "msg being used")
		err := greetCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("the message is - %s", *msgFlag)
	case "help":
		fmt.Println("some help msg")
	default:
		fmt.Printf("command not found - %s", cmd)
	}
}