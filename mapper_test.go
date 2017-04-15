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

func TestCanIgnoreField(t *testing.T) {
	t.Log("Ignoring firstname (expecting an empty string)")

	gm := NewDefault()

	source := employeeSource{"John", "Doe", 1000}
	destination := employeeDestination{}

	gm.Add(source, destination, map[string]MapConfig{
		"FirstName": {true, "LastName"},
	})
	if gm.Map(source, &destination); destination.FirstName != "" {
		t.Errorf("Test failed, LastName should be empty.")
	}
}
