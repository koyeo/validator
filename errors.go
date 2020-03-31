package validator

import "encoding/json"

func NewErrors() *Errors {
	return &Errors{
		errors: make(map[string]string),
	}
}

type Errors struct {
	errors map[string]string
}

func (p *Errors) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.errors)
}

func (p *Errors) Error() string {
	if p != nil && p.errors != nil {
		r, err := json.Marshal(p.errors)
		if err != nil {
			return err.Error()
		}
		return string(r)
	}
	return ""
}

func (p *Errors) Add(field, msg string) {
	p.errors[field] = msg
}

func (p *Errors) Get(field string) string {
	if err, ok := p.errors[field]; ok {
		return err
	}
	return ""
}
