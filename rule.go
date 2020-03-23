package validator

import (
	"fmt"
	"regexp"
)

const (
	username   = `username`
	password   = `password`
	email      = `email`
	url        = `url`
	integer    = `integer`
	float      = `float`
	chinese    = `chinese`
	length     = "length"
	maxLength  = "maxLength"
	minLength  = "minLength"
	max        = "max"
	min        = "min"
	between    = "between"
	in         = "in"
	notIn      = "notIn"
	equal      = "equal"
	great      = "great"
	less       = "less"
	greatEqual = "greatEqual"
	lessEqual  = "lessEqual"
	notEqual   = "notEqual"
)

var rules = map[string]string{
	username: `test`,
	password: `test`,
}

var messages = map[string]string{
	username: "用户名格式不正确",
	password: "密码格式不正确",
}

type Flow struct {
	validator *Validator
	field     string
	label     string
	values    []interface{}
}

func (p *Flow) getMessage(rule string, msg []string) (message string) {

	if len(msg) > 0 {
		return msg[0]
	}

	return messages[rule]
}

func (p *Flow) Rule(check func(validator *Validator, flow *Flow)) {
	check(p.validator, p)
}

func (p *Flow) Username(msg ...string) *Flow {

	if p.validator.hasError(p.label) {
		return p
	}

	reg := regexp.MustCompile(rules[username])
	for _, v := range p.values {
		if !reg.MatchString(fmt.Sprintf("%+v", v)) {
			p.validator.addError(p.field, p.getMessage(username, msg))
			return p
		}
	}

	return p
}

func (p *Flow) Password(msg ...string) *Flow {

	if p.validator.hasError(p.label) {
		return p
	}

	reg := regexp.MustCompile(rules[password])
	for _, v := range p.values {
		if !reg.MatchString(fmt.Sprintf("%+v", v)) {
			p.validator.addError(p.field, p.getMessage(password, msg))
			return p
		}
	}

	return p
}
