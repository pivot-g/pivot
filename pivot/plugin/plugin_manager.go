package plugin

import (
	"encoding/json"
	"fmt"
	"plugin"

	"github.com/go-playground/validator/v10"
	"github.com/pivot-g/pivot/pivot/utility"
)

type PluginConfig struct {
	Name       string   `yaml:"Name,omitempty" json:"Name,omitempty" validate:"required"`
	Plugin     string   `yaml:"Plugin,omitempty" json:"Plugin,omitempty" validate:"required"`
	Main       string   `yaml:"Main,omitempty" json:"Main,omitempty" validate:"required"`
	Type       string   `yaml:"Type,omitempty" json:"Type,omitempty" validate:"required"`
	Version    string   `yaml:"Version,omitempty" json:"Version,omitempty" validate:"required,semver"`
	Dependency []string `yaml:"Dependency,omitempty" json:"Dependency,omitempty"`
}

type PluginMap struct {
	Func       func(map[string]interface{}, map[string]PluginMap) map[string]interface{}
	Validation func(map[string]interface{}) map[string]interface{}
	Dependency []string
	Name       string
	Type       string
	Version    string
}

type Plugin struct {
	PluginDir string
	PluginMap map[string]PluginMap
}

func (p Plugin) LoadPlugins() {
	o := p.ReadPlugin()
	fmt.Println("p.ReadPlugin()")
	fmt.Println(o)
	for _, plug := range o {
		dir := p.PluginDir + "/" + plug.Name + ".so"
		lplug, _ := plugin.Open(dir)

		p.PluginMap[plug.Name] = PluginMap{Func: GetFunc(plug.Main, lplug),
			Validation: GetValidationFunc("Validation", lplug),
			Dependency: plug.Dependency,
			Name:       plug.Name,
			Type:       plug.Type,
			Version:    plug.Version,
		}
		fmt.Println(p.PluginMap)
	}
}

func (p Plugin) ReadPlugin() []PluginConfig {
	configOut := []PluginConfig{}
	valid := validator.New()
	for file, content := range *utility.ReadYamlandJsonFile(&p.PluginDir) {
		localConf := PluginConfig{}
		json.Unmarshal(content, &localConf)
		err := valid.Struct(localConf)
		if err == nil {
			configOut = append(configOut, localConf)
		} else {
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println(err.Field(), "is", err.Tag(), "on plugin", file)
			}
		}
	}
	return configOut

}

func GetPARAM(p *plugin.Plugin) *map[string]interface{} {
	s, _ := p.Lookup("Parms")
	return s.(*map[string]interface{})
}

func GetFunc(n string, p *plugin.Plugin) func(map[string]interface{}) map[string]interface{} {
	Func, _ := p.Lookup(n)
	from_parent := Func.(func(map[string]interface{}) map[string]interface{})
	return from_parent
}

func GetValidationFunc(n string, p *plugin.Plugin) func(map[string]interface{}) map[string]interface{} {
	Func, _ := p.Lookup(n)
	from_parent := Func.(func(map[string]interface{}) map[string]interface{})
	return from_parent
}

func (p Plugin) GetPluginType(name string) string {
	return p.PluginMap[name].Type

}
