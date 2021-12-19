package models

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Employee struct {
	Name string
}

func (e Employee) Validate() error {
	return validation.ValidateStruct(&e, e.ValidateFieldRules()...)
}

func (e *Employee) ValidateFieldRules() []*validation.FieldRules {
	return []*validation.FieldRules{
		validation.Field(&e.Name, validation.Required),
	}
}

type Manager struct {
	Employee
	Level int
}

func (m Manager) Validate() error {
	return validation.ValidateStruct(&m, m.ValidateFieldRules()...)
}

func (m *Manager) ValidateFieldRules() []*validation.FieldRules {
	return []*validation.FieldRules{
		validation.Field(&m.Employee),
		validation.Field(&m.Level, validation.Required),
	}
}
