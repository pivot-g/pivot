package plugin

import (
	"encoding/json"
	"fmt"
	gplugin "plugin"

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
	Func          func(map[string]interface{}) map[string]interface{}
	Validation    func(map[string]interface{}) map[string]interface{}
	Dependency    []string
	Name          string
	Type          string
	Version       string
	DependencyMap map[string]func(map[string]interface{}) map[string]interface{}
}

type Plugin struct {
	Dependency map[string]func(map[string]interface{}) map[string]interface{}
	Input      map[string]interface{}
}

type Plugins struct {
	PluginDir string
	PluginMap map[string]PluginMap
}

func (p Plugins) LoadDependencyMap() {
	for name, plug := range p.PluginMap {
		//plug.DependencyMap = make(map[string]func(map[string]interface{}) map[string]interface{})
		for _, depend := range plug.Dependency {
			p.PluginMap[name].DependencyMap[depend] = p.PluginMap[depend].Func
		}
	}

}

func (p Plugins) LoadPlugins() {
	o := p.ReadPlugin()
	fmt.Println("p.ReadPlugin()")
	fmt.Println(o)
	for _, plug := range o {
		path := p.PluginDir + "/" + plug.Name + ".so"
		lplug, err := gplugin.Open(path)
		if err != nil {
			fmt.Println("Error")
			fmt.Println(err)
		}
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

func (p Plugins) ReadPlugin() []PluginConfig {
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

func GetPARAM(p *gplugin.Plugin) *map[string]interface{} {
	s, _ := p.Lookup("Parms")
	return s.(*map[string]interface{})
}

func GetFunc(n string, p *gplugin.Plugin) func(map[string]interface{}) map[string]interface{} {
	Func, err := p.Lookup(n)
	fmt.Println(err)
	from_parent := Func.(func(map[string]interface{}) map[string]interface{})
	return from_parent
}

func GetValidationFunc(n string, p *gplugin.Plugin) func(map[string]interface{}) map[string]interface{} {
	Func, _ := p.Lookup(n)
	from_parent := Func.(func(map[string]interface{}) map[string]interface{})
	return from_parent
}

func (p Plugins) GetPluginType(name string) string {
	return p.PluginMap[name].Type

}
