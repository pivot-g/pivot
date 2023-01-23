package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Parms = map[string]interface{}{"name": "required", "region": "required", "assume_role": ""}

var Doc string = `
this is example
`

type ReturnValue struct {
}

func AwsProfile(m interface{}) interface{} {
	fmt.Println("")

	return map[string]interface{}{"h": "Hellow World..."}
}

func Validation(c map[string]interface{}) map[string]interface{} {
	validate := validator.New()
	errs := validate.ValidateMap(c, Parms)

	return errs
}
