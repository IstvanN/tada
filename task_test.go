package main

import "testing"

func TestNewTask(t *testing.T) {
	expected := "testing"
	task := newTask(expected)

	if task.descr != expected {
		t.Errorf("expected: %v, actual: %v", expected, task.descr)
	}
	if task.done {
		t.Errorf("expected: %v, actual: %v", expected, task.done)
	}
}
