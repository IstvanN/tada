package main

import "testing"

func TestNewTask(t *testing.T) {
	expected := "testing"
	task := newTask(expected)

	if task.name != expected {
		t.Errorf("expected: %v, actual: %v", expected, task.name)
	}
	if task.done {
		t.Errorf("expected: %v, actual: %v", expected, task.done)
	}
}
