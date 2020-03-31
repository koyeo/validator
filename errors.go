package validator

import (
	"encoding/json"
)

func NewErrors() Errors {
	return make(map[string]string)
}

type Errors map[string]string


func (p Errors) Error() string {
	if p != nil {
		r, err := json.Marshal(p)
		if err != nil {
			return err.Error()
		}
		return string(r)
	}
	return ""
}

func (p Errors) Add(field, msg string) {
	p[field] = msg
}

func (p Errors) Get(field string) string {
	if err, ok := p[field]; ok {
		return err
	}
	return ""
}
