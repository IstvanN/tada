package main

import (
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
		tsk := newTask(expr)
		tasks = append(tasks, *tsk)
	}
	return tasks
}
