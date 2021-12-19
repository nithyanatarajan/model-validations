package models_test

import (
	"fmt"
	"github.com/nithyanatarajan/validations/pkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestManager_Validate(t *testing.T) {
	m := models.Manager{}
	err := m.Validate()

	fmt.Println(err)
	// Output: (if Employee is embedded type)
	// Level: cannot be blank; Name: cannot be blank.
	// Output (if Employee is not embedded type)
	// Employee: (Name: cannot be blank.); Level: cannot be blank.

	// Output (validation.Field(&m.Employee) is not given)
	// Level: cannot be blank.
	assert.EqualError(t, err, "Level: cannot be blank; Name: cannot be blank.")
}

func TestEmployee_Validate(t *testing.T) {
	employee := models.Employee{}

	err := employee.Validate()

	assert.EqualError(t, err, "Name: cannot be blank.")
}
