package main

import (
	"flag"
)

var lFlag bool
var aFlag string
var rFlag int
var cFlag int

func parseFlags() {
	flag.Usage = showHelp
	flag.BoolVar(&lFlag, "l", false, "lists all the tasks")
	flag.StringVar(&aFlag, "a", "", "adds a new task")
	flag.IntVar(&rFlag, "r", 0, "removes the task of the number provided")
	flag.IntVar(&cFlag, "c", 0, "checks the task of the number probided")
	flag.Parse()
}
