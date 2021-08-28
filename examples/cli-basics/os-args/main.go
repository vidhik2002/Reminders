package main

import (
	"fmt"
	"os"
	"strings"
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
		msg := "Reminder"
		if len(os.Args) > 2{
			f := strings.Split(os.Args[2], "=")
			if len(f) == 2 && f[0] == "--msg" {
				msg = f[1]
			}
			fmt.Printf("the message is - %s", msg)
		}
	case "help":
		fmt.Println("some help msg")
	default:
		fmt.Printf("command not found - %s", cmd)
	}
}