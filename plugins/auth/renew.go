package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/pivot-g/pivot/pivot/plugin"
)

var Parms = map[string]interface{}{"name": "required", "schedule": "required", "password_policy": "required"}

func Renew(m interface{}) interface{} {
	fmt.Println("p.Dependency")
	m.(plugin.Plugin).Dependency["aws_secret"](map[string]string{"from": "reniv"})

	return map[string]interface{}{"h": "Hellow World..."}
}

func Validation(c map[string]interface{}) map[string]interface{} {
	validate := validator.New()
	errs := validate.ValidateMap(c, Parms)
	return errs
}
