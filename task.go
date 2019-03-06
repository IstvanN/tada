package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type task struct {
	name string
	done bool
}

func newTask(name string) *task {
	t := &task{
		name: name,
		done: false,
	}
	return t
}

func readTasksFromFile(filename string) []task {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var tasks []task
	s := strings.Split(string(b), ",")
	for _, expr := range s {
		t := newTask(expr)
		tasks = append(tasks, *t)
	}
	return tasks
}

func listTasks(tasks []task) {
	for i, t := range tasks {
		fmt.Printf("%v -  %v\n", i+1, t.name)
	}
}
