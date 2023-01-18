package main

import (
	"fmt"

	"github.com/pivot-g/pivot/pivot/config"
	"github.com/pivot-g/pivot/pivot/plugin"
	"github.com/pivot-g/pivot/pivot/utility"
	"golang.org/x/exp/maps"
)

func main() {
	Plugins := &plugin.Plugin{
		PluginDir: "/Users/natsaiso/go/src/pivot-g/pivot/plugins/auth/",
		PluginMap: make(map[string]plugin.PluginMap),
	}

	Plugins.LoadPlugins()
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
	fmt.Println(conf)

	for _, c := range conf.ReadConfig() {
		for k, v := range c {
			fmt.Println("k, v")
			fmt.Println(k, v)

			d := utility.ConvMapInterface(v.(map[interface{}]interface{}))
			Plugins.PluginMap[k].Validation(d)

			refType, incType := GetPlugingMentioned(Plugins.PluginMap[k].Dependency, d)
			fmt.Println("refType", refType, "incType", incType)
			for _, depend := range incType {
				x := utility.ConvMapInterface(d[depend].(map[interface{}]interface{}))
				Plugins.PluginMap[depend].Validation(x)
			}
			for _, depend := range refType {
				if d[depend] == maps.Keys(conf.ConfigMap[depend])[0] {
					fmt.Println(depend, "present")

				} else {
					fmt.Println(depend, "not present")
				}

			}
		}

	}

}

func GetPlugingMentioned(dependency []string, configBlock map[string]interface{}) ([]string, []string) {
	refType := []string{}
	incType := []string{}
	for k, _ := range configBlock {
		fmt.Println("k")
		fmt.Println(k)
		if utility.In(dependency, "ref:"+k) {
			refType = append(refType, k)
		}
		if utility.In(dependency, "inc:"+k) {
			incType = append(incType, k)
		}
	}
	return refType, incType
}

// func Val(str, p map[string]plugin.PluginMap) {
// 	p[k].Validation(d)

// }
