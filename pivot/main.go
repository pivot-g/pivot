package main

import (
	"fmt"

	"github.com/pivot-g/pivot/pivot/config"
	"github.com/pivot-g/pivot/pivot/plugin"
	"github.com/pivot-g/pivot/pivot/validation"
)

func main() {
	Plugins := &plugin.Plugins{
		PluginDir: "/Users/natsaiso/go/src/pivot-g/pivot/plugins/auth/",
		PluginMap: make(map[string]map[string]plugin.PluginMap),
	}

	// a := plugin.PluginMap{}
	// a.DependencyMap = make(map[string]func(map[string]interface{}) map[string]interface{})

	Plugins.LoadPlugins()
	// fmt.Println(Plugins.PluginMap)
	conf := &config.Config{
		Dir:       "/Users/natsaiso/go/src/pivot-g/pivot/example/config",
		PluginMap: Plugins.PluginMap,
		ConfigMap: make(map[string]map[string]map[string]interface{}),
	}
	conf.LoadConfig()

	fmt.Println(Plugins.PluginMap)
	validation.ConfigVal(conf, Plugins)
	Plugins.LoadDependencyMap()
	fmt.Println(Plugins.PluginMap)
	fmt.Println(Plugins.GetLatestPlugin("renew"))

}
