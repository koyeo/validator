package validator

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	Require    = "Require"
	Username   = `Username`
	Password   = `Password`
	Email      = `Email`
	Url        = `Url`
	Integer    = `Integer`
	Float      = `Float`
	Chinese    = `Chinese`
	Ip         = "Ip"
	ChineseId  = "ChineseId"
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
	Username:  `^[a-z0-9_-]{3,16}$`,
	Password:  `^[0-9a-zA-Z@.]{6,30}$`,
	Email:     `^([a-z0-9_\.-]+)@([\da-z\.-]+)\.([a-z\.]{2,6})$`,
	Url:       `^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$/`,
	Integer:   `-?\d+`,
	Float:     `(-?\d+)(\.\d+)?`,
	Ip:        `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/`,
	Chinese:   ` [\u4e00-\u9fa5]`,
	ChineseId: `\d{15}|\d{18}`,
}

var messages = map[string]string{
	Require:   "字段不能为空",
	Username:  "用户名格式不正确",
	Password:  "密码格式不正确",
	Email:     "邮箱格式不正确",
	Url:       "URL 格式不正确",
	Integer:   "请输入整数",
	Float:     "小数格式不正确",
	Ip:        "IP 地址格式不正确",
	Chinese:   "请输入中文",
	ChineseId: "身份证号码格式不正确",
}

type Flow struct {
	validator *Validator
	field     string
	label     string
	values    []interface{}
}

func (p *Flow) GetField() string {
	return p.field
}

func (p *Flow) GetLabel() string {
	return p.label
}

func (p *Flow) GetValue(index int) interface{} {
	if index < len(p.values) {
		return p.values[index]
	}
	return nil
}

func (p *Flow) GetValues() []interface{} {
	return p.values
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

func (p *Flow) Require(msg ...string) *Flow {

	if p.validator.CheckError(p.label) {
		return p
	}

	for _, v := range p.values {
		if strings.TrimSpace(fmt.Sprintf("%+v", v)) == "" {
			p.validator.AddError(p.field, p.GetMessage(Require, msg))
			return p
		}
	}

	return p
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
