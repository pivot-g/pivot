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
		PluginMap: make(map[string]plugin.PluginMap),
	}

	Plugins.LoadPlugins()
	Plugins.LoadDependencyMap()
	// fmt.Println(Plugins.PluginMap)
	conf := &config.Config{
		Dir:       "/Users/natsaiso/go/src/pivot-g/pivot/example/config",
		PluginMap: Plugins.PluginMap,
		Api:       make(map[string]map[string]map[string]interface{}),
		Auth:      make(map[string]map[string]map[string]interface{}),
		Scheduler: make(map[string]map[string]map[string]interface{}),
		Resources: make(map[string]map[string]map[string]interface{}),
		Policy:    make(map[string]map[string]map[string]interface{}),
		ConfigMap: make(map[string]map[string]map[string]interface{}),
	}
	conf.LoadConfig()

	fmt.Println(Plugins.PluginMap)
	validation.ConfigVal(conf, Plugins)

	// d := map[string]func(plugin.Plugin) map[string]interface{}{
	// 	"renew": Plugins.PluginMap["renew"].Dependency
	// }
	// Plugins.PluginMap["renew"].Func(plugin.Plugin{
	// 	Dependency: Plugins.PluginMap["renew"].Dependency,
	// })

}

// func ConfigVal(conf *config.Config, Plugins *plugin.Plugin) {
// 	for _, c := range conf.ReadConfig() {
// 		BlockVal(c, conf, Plugins)

// 	}
// }

// func BlockVal(block map[string]interface{}, conf *config.Config, Plugins *plugin.Plugin) {
// 	for k, v := range block {
// 		fmt.Println("k, v")
// 		fmt.Println(k, v)

// 		d := utility.ConvMapInterface(v.(map[interface{}]interface{}))
// 		Plugins.PluginMap[k].Validation(d)

// 		refType, incType := config.GetPlugingMentioned(Plugins.PluginMap[k].Dependency, d)
// 		fmt.Println("refType", refType, "incType", incType)

// 		for _, depend := range incType {
// 			//subBlock := utility.ConvMapInterface(d[depend].(map[interface{}]interface{}))
// 			subBlock := map[string]interface{}{depend: d[depend]}
// 			fmt.Println("subBlock", subBlock)
// 			//Plugins.PluginMap[depend].Validation(subBlock)
// 			BlockVal(subBlock, conf, Plugins)
// 		}
// 		for _, depend := range refType {
// 			if d[depend] == maps.Keys(conf.ConfigMap[depend])[0] {
// 				fmt.Println(depend, "present")

// 			} else {
// 				fmt.Println(depend, "not present")
// 			}

// 		}
// 	}

// }
