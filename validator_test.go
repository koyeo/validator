package validator

import (
	"fmt"
	"testing"
)

type OK struct {
	Name string `json:"name"`
}

type A struct {
	Code int64
	Err  interface{}
}

func TestNewValidator(t *testing.T) {

	var a = &OK{Name: "1te123st2123"}

	v := NewValidator()
	v.Validate("Username", "用户名", a.Name).Username()
	v.Validate("Password", "密码", a.Name).Password()
	//v.Validate("age", "年龄", 8).Rule(func(validator *Validator, flow *Flow) {
	//
	//	if validator.CheckError(flow.field) {
	//		return
	//	}
	//	for _, v := range flow.values {
	//		v, err := strconv.ParseFloat(fmt.Sprintf("%+v", v), 64)
	//		if err != nil {
	//			validator.AddError(flow.field, "年龄格式不正确")
	//			return
	//		}
	//
	//		if v < 10 {
	//			validator.AddError(flow.field, "年龄不能大于10岁")
	//			return
	//		}
	//	}
	//})

	if v.HasError() {
		fmt.Println(v.Error())
	}

}
