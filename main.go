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

func main() {
	printHeader()
	parseFlags()
	tasks = readTasksFromFile(datafile)
	if lFlag {
		listTasks(tasks)
		return
	}
	if aFlag != "" {
		t = newTask(aFlag)
		tasks = append(tasks, t)
		writeTasksToFile(tasks, datafile)
		fmt.Printf("'%v' has been added to your todos!", t.descr)
		return
	}
	if rFlag != 0 {
		var err error
		tasks, err = removeTask(tasks, rFlag)
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
