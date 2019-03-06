package main

import (
	"flag"
)

var lFlag bool

func parseFlags() {
	flag.Usage = showHelp
	flag.BoolVar(&lFlag, "l", false, "lists all the tasks")
	flag.Parse()
}
