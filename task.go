package main

type task struct {
	name string
}

func newTask(name string) *task {
	t := &task{
		name: name,
	}

	return t
}
