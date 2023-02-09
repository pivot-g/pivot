package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Parms = map[string]interface{}{"name": "required", "region": "required", "assume_role": ""}

func AwsSecret(m interface{}) interface{} {
	fmt.Println("hellow swcret")
	fmt.Println(m)

	// fmt.Println(p["d"].(plugin.Plugin).Dependency["password_policy"](map[string]interface{}{"a": "s"}))

	return map[string]interface{}{"h": "Hellow World..."}
}

func Validation(c map[string]interface{}) map[string]interface{} {
	validate := validator.New()
	errs := validate.ValidateMap(c, Parms)

	return errs
}

func Get() {

}

func Put() {

}
