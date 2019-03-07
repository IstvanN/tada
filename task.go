package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type task struct {
	descr string
	done  bool
}

func newTask(descr string) task {
	t := task{
		descr: descr,
		done:  false,
	}
	return t
}

func convertTasksToCSV(tasks []task) string {
	var s []string
	for _, t := range tasks {
		if t.done {
			s = append(s, "1"+t.descr)
		} else {
			s = append(s, "0"+t.descr)
		}
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
		if isFirstCharZero(expr) {
			expr = string(expr[1:])
			t := newTask(expr)
			tasks = append(tasks, t)
		} else {
			expr = string(expr[1:])
			t := newTask(expr)
			t.done = true
			tasks = append(tasks, t)
		}
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
			if t.done {
				fmt.Printf("%v -  [X] %v\n", i+1, t.descr)
			} else {
				fmt.Printf("%v -  [ ] %v\n", i+1, t.descr)
			}
		}
	} else {
		fmt.Println("You have no tasks!")
	}
}

func addTask(tasks []task, tdescr string) []task {
	fmt.Printf("'%v' has been added to tasks\n", tdescr)
	return append(tasks, newTask(tdescr))
}

func removeTask(tasks []task, sernum int) ([]task, error) {
	if sernum > len(tasks) {
		err := errors.New("Can't delete task: index is out of bound")
		return nil, err
	}
	i := sernum - 1
	for j := range tasks {
		if i == j {
			removed := tasks[i]
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("'%v' has been removed from the tasks\n", removed.descr)
		}
	}
	return tasks, nil
}

func isFirstCharZero(s string) bool {
	return string(s[0]) == "0"
}
