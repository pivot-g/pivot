package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/pivot-g/pivot/pivot/plugin"
)

var Parms = map[string]interface{}{"name": "required", "schedule": "required", "password_policy": "required"}

func Renew(m map[string]interface{}, p map[string]plugin.PluginMap) map[string]interface{} {
	p["jhh"].Func.(func(map[string]interface{}, map[string]plugin.PluginMap) map[string]interface{})(m, p)

	return map[string]interface{}{"h": "Hellow World..."}
}

func Validation(c map[string]interface{}) map[string]interface{} {
	validate := validator.New()
	errs := validate.ValidateMap(c, Parms)
	return errs
}
