package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Parms = map[string]interface{}{"name": "required", "schedule": "required", "password_policy": "required"}

func Renew(m map[string]interface{}) map[string]interface{} {
	fmt.Println("p.Dependency")

	return map[string]interface{}{"h": "Hellow World..."}
}

func Validation(c map[string]interface{}) map[string]interface{} {
	validate := validator.New()
	errs := validate.ValidateMap(c, Parms)
	return errs
}
