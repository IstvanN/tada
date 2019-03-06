package main

import (
	"flag"
)

var lFlag bool
var aFlag string

func parseFlags() {
	flag.Usage = showHelp
	flag.BoolVar(&lFlag, "l", false, "lists all the tasks")
	flag.StringVar(&aFlag, "a", "", "adds a new task")
	flag.Parse()
}
