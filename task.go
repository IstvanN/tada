package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type task struct {
	descr string
	done  bool
}

func newTask(descr string) *task {
	t := &task{
		descr: descr,
		done:  false,
	}
	return t
}

func convertTasksToCSV(tasks []task) string {
	var s []string
	for _, t := range tasks {
		s = append(s, t.descr)
	}
	return strings.Join(s, ",")
}

func writeTasksToFile(tasks []task, filename string) {
	csv := convertTasksToCSV(tasks)
	err := ioutil.WriteFile(filename, []byte(csv), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func convertBytesToTasks(b []byte) []task {
	var tasks []task
	s := strings.Split(string(b), ",")
	for _, expr := range s {
		t := newTask(expr)
		tasks = append(tasks, *t)
	}
	return tasks
}

func readTasksFromFile(filename string) []task {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return convertBytesToTasks(b)
}

func listTasks(tasks []task) {
	if len(tasks) > 0 && tasks[0].descr != "" {
		for i, t := range tasks {
			fmt.Printf("%v -  %v\n", i+1, t.descr)
		}
	} else {
		fmt.Println("You have no tasks!")
	}
}
