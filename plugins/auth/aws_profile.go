package main

import (
	"github.com/go-playground/validator/v10"
)

var Parms = map[string]interface{}{"name": "required", "region": "required", "assume_role": ""}

func AwsProfile(m map[string]interface{}) map[string]interface{} {

	return map[string]interface{}{"h": "Hellow World..."}
}

func Validation(c map[string]interface{}) map[string]interface{} {
	validate := validator.New()
	errs := validate.ValidateMap(c, Parms)

	return errs
}
