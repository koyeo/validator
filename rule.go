package validator

import (
	"fmt"
	"regexp"
)

const (
	Username   = `Username`
	Password   = `Password`
	Email      = `Email`
	Url        = `Url`
	Integer    = `Integer`
	Float      = `Float`
	Chinese    = `Chinese`
	Length     = "Length"
	MaxLength  = "MaxLength"
	MinLength  = "MinLength"
	Max        = "Max"
	Min        = "Min"
	Between    = "Between"
	In         = "In"
	NotIn      = "NotIn"
	Equal      = "Equal"
	Great      = "Great"
	Less       = "Less"
	GreatEqual = "GreatEqual"
	LessEqual  = "LessEqual"
	NotEqual   = "NotEqual"
)

var rules = map[string]string{
	Username: `test`,
	Password: `test`,
}

var messages = map[string]string{
	Username: "用户名格式不正确",
	Password: "密码格式不正确",
}

type Flow struct {
	validator *Validator
	field     string
	label     string
	values    []interface{}
}

func (p *Flow) GetLabel() string {
	return p.label
}

func (p *Flow) GetMessage(rule string, msg []string) (message string) {

	if len(msg) > 0 {
		return msg[0]
	}

	return messages[rule]
}

func (p *Flow) Rule(check func(validator *Validator, flow *Flow)) {
	check(p.validator, p)
}

func (p *Flow) Username(msg ...string) *Flow {

	if p.validator.CheckError(p.label) {
		return p
	}

	reg := regexp.MustCompile(rules[Username])
	for _, v := range p.values {
		if !reg.MatchString(fmt.Sprintf("%+v", v)) {
			p.validator.AddError(p.field, p.GetMessage(Username, msg))
			return p
		}
	}

	return p
}

func (p *Flow) Password(msg ...string) *Flow {

	if p.validator.CheckError(p.label) {
		return p
	}

	reg := regexp.MustCompile(rules[Password])
	for _, v := range p.values {
		if !reg.MatchString(fmt.Sprintf("%+v", v)) {
			p.validator.AddError(p.field, p.GetMessage(Password, msg))
			return p
		}
	}

	return p
}
