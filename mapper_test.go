package gomap

import (
	"testing"
)

type employeeSource struct {
	FirstName string
	LastName  string
	Salary    int
}

type employeeDestination struct {
	FirstName string
	LastName  string
}

func TestCanIgnore(t *testing.T) {
	t.Log("Ignoring firstname (expecting an empty string)")
	gm := New(Options{
		IgnoreFields: []string{"FirstName"},
	})

	source := employeeSource{"John", "Doe", 1000}
	destination := employeeDestination{}

	if gm.Map(source, &destination); destination.FirstName != "" {
		t.Errorf("Test failed, LastName should be empty.")
	}
}
