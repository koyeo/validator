package validator

import (
	"encoding/json"
)

type Errors map[string]string

func (e Errors) Error() string {
	r, err := json.Marshal(e)
	if err != nil {
		return err.Error()
	}
	return string(r)
}

func NewValidator() *Validator {
	return &Validator{}
}

type Validator struct {
	errors Errors
}

func (p *Validator) Error() error {
	if len(p.errors) == 0 {
		p.errors = nil
	}
	return p.errors
}

func (p *Validator) AddError(field string, msg string) {
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

func (p *Validator) CheckError(field string) bool {
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
