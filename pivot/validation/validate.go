package validation

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/pivot-g/pivot/pivot/config"
	"github.com/pivot-g/pivot/pivot/plugin"
	"github.com/pivot-g/pivot/pivot/utility"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func ConfigVal(conf *config.Config, Plugins *plugin.Plugins) {
	valid := validator.New()
	for name, plugs := range Plugins.PluginMap {
		for version, plug := range plugs {
			err := valid.Struct(plug)
			if err != nil {
				for _, err := range err.(validator.ValidationErrors) {
					fmt.Println(err.Field(), "is", err.Tag(), "on plugin", name, "version", version)
					os.Exit(2)
				}
			}
		}

	}
	for _, c := range conf.ReadConfig() {
		BlockVal(c, conf, Plugins, "")

	}

}

func BlockVal(block map[string]interface{}, conf *config.Config, Plugins *plugin.Plugins, parentPlug string) {
	for k, v := range block {
		// fmt.Println("k, v")
		// fmt.Println("parentPlug")
		// fmt.Println(parentPlug)
		// fmt.Println(k, v)
		d := utility.ConvMapInterface(v.(map[interface{}]interface{}))
		version := ""
		if parentPlug == "" {
			version = Plugins.GetLatestPlugin(k)

		} else {
			version = Plugins.GetLatestPlugin(parentPlug)
			version = Plugins.PluginMap[parentPlug][version].Dependency[k].Version
		}

		// fmt.Println("version")
		// fmt.Println(version)
		errs := Plugins.PluginMap[k][version].Validation(d)
		if len(errs) > 0 {
			fmt.Println(errs)
			os.Exit(2)
		}

		refType, incType := config.GetPlugingMentioned(Plugins.PluginMap[k][version].Dependency, d)
		fmt.Println("refType", refType, "incType", incType)

		for _, depend := range incType {
			subBlock := map[string]interface{}{depend: d[depend]}
			fmt.Println("subBlock", subBlock)
			BlockVal(subBlock, conf, Plugins, k)
		}
		for _, depend := range refType {
			fmt.Println("d[depend]", d[depend])
			fmt.Println("maps.Keys(conf.ConfigMap[depend])", maps.Keys(conf.ConfigMap[depend]))

			if slices.Contains(maps.Keys(conf.ConfigMap[depend]), d[depend].(string)) {
				fmt.Println(depend, "present")

			} else {
				fmt.Println(depend, "not present")
			}

		}
	}

}
