package config

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/pivot-g/pivot/pivot/log"
	"github.com/pivot-g/pivot/pivot/plugin"
	"github.com/pivot-g/pivot/pivot/utility"
	"golang.org/x/exp/maps"
	"gopkg.in/yaml.v2"
)

type ConfigMap struct {
	Config map[string]map[string]map[string]interface{}
}
type Config struct {
	Dir       string
	Api       map[string]map[string]map[string]interface{}
	Auth      map[string]map[string]map[string]interface{}
	Scheduler map[string]map[string]map[string]interface{}
	Resources map[string]map[string]map[string]interface{}
	Policy    map[string]map[string]map[string]interface{}
	ConfigMap map[string]map[string]map[string]interface{}
	PluginMap map[string]plugin.PluginMap
}

func (c *Config) ReadConfig() []map[string]interface{} {
	configOut := []map[string]interface{}{}
	for file, content := range *utility.ReadYamlandJsonFile(&c.Dir) {
		config := []map[string]interface{}{}
		split := strings.Split(file, ".")
		fileType := strings.ToLower(split[len(split)-1])
		log.Debug(content, "byte data from file ", file)
		if utility.In([]string{"yaml", "yml"}, fileType) {
			err := yaml.Unmarshal(content, &config)
			if err != nil {

				log.Fatal(err)
			}
		}
		if utility.In([]string{"json"}, fileType) {
			err := json.Unmarshal(content, &config)
			if err != nil {

				log.Fatal(err)
			}
		}
		configOut = append(configOut, config...)
	}
	return configOut
}

func (c *Config) LoadConfig() {
	for _, block := range c.ReadConfig() {
		for k, v := range block {
			// fmt.Println("k, v ")
			// fmt.Println(k, v)
			blockConf := utility.ConvMapInterface(v.(map[interface{}]interface{}))
			name := blockConf["name"].(string)
			bb := map[string]map[string]interface{}{name: blockConf}
			if c.ConfigMap[k] == nil {
				c.ConfigMap[k] = bb
			} else {
				maps.Copy(c.ConfigMap[k], bb)
			}

			// fmt.Println("name")
			// fmt.Println(name)

			switch Type := c.PluginMap[k].Type; Type {
			case "auth":
				if c.Auth[k] == nil {
					c.Auth[k] = bb
				} else {
					maps.Copy(c.Auth[k], bb)
				}

			case "scheduler":
				if c.Scheduler[k] == nil {
					c.Scheduler[k] = bb
				} else {
					maps.Copy(c.Scheduler[k], bb)
				}
			case "api":
				if c.Api[k] == nil {
					c.Api[k] = bb
				} else {
					maps.Copy(c.Api[k], bb)
				}
			case "resources":
				if c.Resources[k] == nil {
					c.Resources[k] = bb
				} else {
					maps.Copy(c.Resources[k], bb)
				}
			case "policy":
				if c.Policy[k] == nil {
					c.Policy[k] = bb
				} else {
					maps.Copy(c.Policy[k], bb)
				}
			default:
				fmt.Printf("unknown plugin refered in config")
			}

		}
	}
}

func Validate(p map[string]interface{}, c map[string]interface{}) {
	validate := validator.New()
	errs := validate.ValidateMap(c, p)

	if len(errs) > 0 {
		fmt.Println(errs)
		// The user is invalid
	}

}
