package main

import (
	"flag"
	"fmt"
	"os"
)

const datafile = "data/data.csv"

var tasks []task

func main() {
	parseFlags()
	printHeader()
	if lFlag {
		tasks = readTasksFromFile(datafile)
		listTasks(tasks)
		return
	}

	showHelp()
}

func printHeader() {
	fmt.Print(
		`
Tada - a CLI todo application
=============================

`)
}

func showHelp() {
	flag.PrintDefaults()
	os.Exit(0)
}
