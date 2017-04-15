package gomap

import (
	"testing"
)

type employee struct {
	FirstName string
	LastName  string
	Salary    int
}

type employeeViewModel struct {
	FirstName string
	LastName  string
}

func TestCanIgnoreField(t *testing.T) {
	t.Log("Ignoring FirstName (expecting: empty string)")

	gm := New()

	source := employee{"John", "Doe", 1000}
	destination := employeeViewModel{}

	gm.Add(source, destination, map[string]FieldConfig{
		"FirstName": {
			Ignore: true,
		},
	})
	if gm.Map(source, &destination); destination.FirstName != "" {
		t.Errorf("Test failed, LastName should be empty.")
	}
}

func TestCanChangeFieldSource(t *testing.T) {
	t.Log("Changing FirstName source to LastName (expecting: Doe)")
	gm := New()

	source := employee{"John", "Doe", 1000}
	destination := employeeViewModel{}

	gm.Add(source, destination, map[string]FieldConfig{
		"FirstName": {
			Source: "LastName",
		},
	})

	gm.Map(employee{"John", "Doe", 1000}, &destination)

	if destination.FirstName != "Doe" || destination.LastName != "Doe" {
		t.Errorf("Test failed, FirstName should be equal to 'Doe'.")
	}
}
