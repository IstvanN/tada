package main

import "testing"

var tt = []task{
	newTask("test"),
	newTask("1"),
	newTask("2"),
}

var descriptions = []string{"test", "1", "2"}

func TestNewTask(t *testing.T) {
	for i, tsk := range tt {
		if tsk.descr != descriptions[i] && tsk.done != false {
			t.Errorf("expected descr: %v, expected done: %v, actual descr: %v, actual done: %v",
				descriptions[i], false, tsk.descr, tsk.done)
		}
	}
}

func TestConvertTasksToCSV(t *testing.T) {
	actual := convertTasksToCSV(tt)

	var expected string
	for i, d := range descriptions {
		if i != len(descriptions)-1 {
			expected += d + ","
		} else {
			expected += d
		}
	}

	if actual != expected {
		t.Errorf("expected: %v but actual: %v", expected, actual)
	}
}

func TestConvertBytesToTasks(t *testing.T) {
	var csv string
	for i, d := range descriptions {
		if i != len(descriptions)-1 {
			csv += d + ","
		} else {
			csv += d
		}
	}

	actual := convertBytesToTasks([]byte(csv))
	for i, tsk := range actual {
		if tt[i] != tsk {
			t.Errorf("expected: %v, actual:%v", tt[i], tsk)
		}
	}
}
