package plugin

import (
	"encoding/json"
	"fmt"
	gplugin "plugin"
	"sort"

	"github.com/go-playground/validator/v10"
	"github.com/pivot-g/pivot/pivot/utility"
	"golang.org/x/exp/maps"
)

type PluginConfig struct {
	Name       string                `yaml:"Name,omitempty" json:"Name,omitempty" validate:"required"`
	Plugin     string                `yaml:"Plugin,omitempty" json:"Plugin,omitempty" validate:"required"`
	Main       string                `yaml:"Main,omitempty" json:"Main,omitempty" validate:"required"`
	Type       string                `yaml:"Type,omitempty" json:"Type,omitempty" validate:"required"`
	Version    string                `yaml:"Version,omitempty" json:"Version,omitempty" validate:"required,semver"`
	Dependency map[string]Dependency `yaml:"Dependency,omitempty" json:"Dependency,omitempty"`
}

type Dependency struct {
	Type    string `yaml:"type,omitempty" json:"type,omitempty" validate:"required"`
	Version string `yaml:"version,omitempty" json:"version,omitempty" validate:"required,semver"`
}

type PluginMap struct {
	Func          func(interface{}) interface{}
	Validation    func(map[string]interface{}) map[string]interface{}
	Dependency    map[string]Dependency
	Name          string
	Type          string
	Version       string
	DependencyMap map[string]DependencyMap
}

type DependencyMap struct {
	Func          func(interface{}) interface{}
	DependencyMap map[string]DependencyMap
}

type Plugin struct {
	Dependency map[string]func(interface{}) interface{}
	Config     map[string]interface{}
}

type Plugins struct {
	PluginDir string
	PluginMap map[string]map[string]PluginMap
}

func (p *Plugins) LoadPlugins() {
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
		if p.PluginMap[plug.Name] == nil {
			p.PluginMap[plug.Name] = map[string]PluginMap{plug.Version: PluginMap{}}
		}
		p.PluginMap[plug.Name][plug.Version] = PluginMap{Func: GetFunc(plug.Main, lplug),
			Validation:    GetValidationFunc("Validation", lplug),
			Name:          plug.Name,
			Type:          plug.Type,
			Version:       plug.Version,
			Dependency:    plug.Dependency,
			DependencyMap: map[string]DependencyMap{},
		}
		fmt.Println(p.PluginMap)
	}
}

func (p *Plugins) GetLatestPlugin(plugName string) string {
	versions := maps.Keys(p.PluginMap[plugName])
	sort.Strings(versions)
	return versions[0]

}

func (p *Plugins) ReadPlugin() []PluginConfig {
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

func (p *Plugins) LoadDependencyMap() {
	for name, plugs := range p.PluginMap {
		for versions, _ := range plugs {
			bb := p.GetDependency(name, versions)
			if len(bb) == 0 {
				continue
			}
			fmt.Println(bb)
			pp := maps.Keys(bb)[0]
			p.PluginMap[name][versions].DependencyMap[pp] = DependencyMap{
				Func:          bb[pp].Func,
				DependencyMap: bb[pp].DependencyMap,
			}
		}

	}

}

func (p *Plugins) MakeDependencyMap(dep map[string]Dependency, name string, versions string) {
	for depend, value := range dep {
		p.PluginMap[name][versions].DependencyMap[depend] = DependencyMap{
			Func: p.PluginMap[depend][value.Version].Func,
		}
		if len(p.PluginMap[depend][value.Version].Dependency) != 0 {
			p.MakeDependencyMap(p.PluginMap[depend][value.Version].Dependency, depend, value.Version)
		}

	}
}

func (p *Plugins) GetDependency(plugName string, version string) map[string]DependencyMap {
	out := make(map[string]DependencyMap)
	for name, depend := range p.PluginMap[plugName][version].Dependency {
		out[name] = DependencyMap{
			Func:          p.PluginMap[name][depend.Version].Func,
			DependencyMap: p.GetDependency(name, depend.Version),
		}
	}
	return out

}

func GetPARAM(p *gplugin.Plugin) *map[string]interface{} {
	s, _ := p.Lookup("Parms")
	return s.(*map[string]interface{})
}

func GetFunc(n string, p *gplugin.Plugin) func(interface{}) interface{} {
	Func, err := p.Lookup(n)
	fmt.Println(err)
	from_parent := Func.(func(interface{}) interface{})
	return from_parent
}

func GetValidationFunc(n string, p *gplugin.Plugin) func(map[string]interface{}) map[string]interface{} {
	Func, _ := p.Lookup(n)
	from_parent := Func.(func(map[string]interface{}) map[string]interface{})
	return from_parent
}
