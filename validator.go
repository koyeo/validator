package validator

import (
	"encoding/json"
)

func NewValidator() *Validator {
	return &Validator{}
}

type Validator struct {
	errors map[string]string
}

func (p *Validator) Error() string {
	r, err := json.Marshal(p.errors)
	if err != nil {
		return err.Error()
	}
	return string(r)
}

func (p *Validator) addError(field string, msg string) {
	if p.errors == nil {
		p.errors = make(map[string]string)
	}
	p.errors[field] = msg
}

func (p *Validator) getError(field string) string {
	if p.errors == nil {
		return ""
	}
	if err, ok := p.errors[field]; ok {
		return err
	}
	return ""
}

func (p *Validator) hasError(field string) bool {
	if p.getError(field) != "" {
		return true
	}
	return false
}

func (p *Validator) HasError() bool {
	return p.errors != nil
}

func (p *Validator) Validate(field, label string, value ...interface{}) *Flow {
	flow := new(Flow)
	flow.label = label
	flow.field = field
	flow.values = value
	flow.validator = p
	return flow
}
