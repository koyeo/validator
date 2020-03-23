package validator

import (
	"fmt"
	"strconv"
	"testing"
)

type OK struct {
	Name string `json:"name"`
}

func TestNewValidator(t *testing.T) {

	var a = &OK{Name: "123"}

	v := NewValidator()
	v.Validate("username", "用户名", a.Name).Username()
	v.Validate("password", "密码", a.Name).Username()
	v.Validate("age", "年龄", 8).Rule(func(validator *Validator, flow *Flow) {

		if validator.hasError(flow.field) {
			return
		}
		for _, v := range flow.values {
			v, err := strconv.ParseFloat(fmt.Sprintf("%+v", v), 64)
			if err != nil {
				validator.addError(flow.field, "年龄格式不正确")
				return
			}

			if v < 10 {
				validator.addError(flow.field, "年龄不能大于10岁")
				return
			}
		}
	})

	if v.HasError() {
		fmt.Println(v.Error())
	}
}
