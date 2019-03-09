package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const datafile = "data/data.csv"

var tasks []task
var t task
var err error

func main() {
	printHeader()
	parseFlags()
	tasks = readTasksFromFile(datafile)
	if lFlag {
		listTasks(tasks)
		return
	}
	if aFlag != "" {
		tasks = addTask(tasks, aFlag)
		writeTasksToFile(tasks, datafile)
		return
	}
	if rFlag != 0 {
		tasks, err = removeTask(tasks, rFlag)
		if err != nil {
			log.Fatal(err)
		}
		writeTasksToFile(tasks, datafile)
		return
	}
	if cFlag != 0 {
		tasks, err = markTaskDone(tasks, cFlag)
		if err != nil {
			log.Fatal(err)
		}
		writeTasksToFile(tasks, datafile)
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
