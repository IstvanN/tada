package main

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
