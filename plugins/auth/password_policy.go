package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Parms = map[string]interface{}{"name": "required", "min": "required"}

func PasswordPolicy(m map[string]interface{}) map[string]interface{} {

	return map[string]interface{}{"h": "Hellow World..."}
}

func Validation(c map[string]interface{}) bool {
	validate := validator.New()
	errs := validate.ValidateMap(c, Parms)
	fmt.Println(errs)

	return true
}
