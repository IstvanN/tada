package main

import (
	"flag"
	"fmt"
	"os"
)

const datafile = "data/data.csv"

func main() {
	parseFlags()
	if lFlag {
		tasks := readTasksFromFile(datafile)
		listTasks(tasks)
		return
	}
}

func printHeader() {
	fmt.Print(
		`
Tada - a CLI todo application
=============================

`)
}

func showHelp() {
	printHeader()
	flag.PrintDefaults()
	os.Exit(0)
}
