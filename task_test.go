package main

import (
	"errors"
	"testing"
)

var tt = []task{
	newTask("test"),
	newTask("1"),
	newTask("2"),
}

var descriptions = []string{"0test", "01", "02"}

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

func TestRemoveTask(t *testing.T) {
	newtt, _ := removeTask(tt, 3)
	for i, tsk := range newtt {
		if tsk != tt[i] {
			t.Errorf("expected: %v, actual: %v", tt[i], tsk)
		}
	}

	othertt, err := removeTask(tt, len(tt)+1)
	if othertt != nil || err == nil {
		t.Errorf("expected tasks: %v, actual tasks: %v, expected error, got %v", nil, othertt, nil)
	}
}

func TestCheckIfTaskIsDone(t *testing.T) {
	cases := []struct {
		s       string
		expDone bool
		expErr  error
	}{
		{"0test", false, nil},
		{"1test", true, nil},
		{"test", false, errors.New("Cannot identify if the following task is done or not: test")},
	}

	for _, c := range cases {
		ok, err := checkIfTaskIsDone(c.s)
		if ok != c.expDone {
			t.Errorf("checkIfTaskIsDone(%v) expected bool: %v, actual bool: %v", c.s, c.expDone, ok)
		}

		if err != nil {
			if err == nil {
				t.Errorf("checkIfTaskIsDone(%v) expected err: %v, actual err: %v", c.s, c.expErr, err)
			}
		}
	}
}
