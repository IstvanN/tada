package main

import "testing"

var tt = []task{
	*newTask("test"),
	*newTask("1"),
	*newTask("2"),
}

func TestNewTask(t *testing.T) {
	expected := []string{"test", "1", "2"}
	for i, tsk := range tt {
		if tsk.descr != expected[i] && tsk.done != false {
			t.Errorf("expected descr: %v expected done: %v, actual descr: %v, actual done: %v",
				expected[i], false, tsk.descr, tsk.done)
		}
	}
}
