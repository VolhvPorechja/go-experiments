package main

import (
	"os"
	"awesomeProject/application"
)

func main() {
	app := application.New()
	if len(os.Args) < 2 {
		app.ShowHelp()
		return
	}
	command := os.Args[1]
	app.Process(command)
}
