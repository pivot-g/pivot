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

func ConfigVal(conf *config.Config, Plugins *plugin.Plugin) {
	for name, plug := range Plugins.PluginMap {
		valid := validator.New()
		err := valid.Struct(plug)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println(err.Field(), "is", err.Tag(), "on plugin", name)
				os.Exit(2)
			}
		}

	}
	for _, c := range conf.ReadConfig() {
		BlockVal(c, conf, Plugins)

	}

}

func BlockVal(block map[string]interface{}, conf *config.Config, Plugins *plugin.Plugin) {
	for k, v := range block {
		fmt.Println("k, v")
		fmt.Println(k, v)

		d := utility.ConvMapInterface(v.(map[interface{}]interface{}))
		errs := Plugins.PluginMap[k].Validation(d)
		if len(errs) > 0 {
			fmt.Println(errs)
			os.Exit(2)
		}

		refType, incType := config.GetPlugingMentioned(Plugins.PluginMap[k].Dependency, d)
		fmt.Println("refType", refType, "incType", incType)

		for _, depend := range incType {
			//subBlock := utility.ConvMapInterface(d[depend].(map[interface{}]interface{}))
			subBlock := map[string]interface{}{depend: d[depend]}
			fmt.Println("subBlock", subBlock)
			//Plugins.PluginMap[depend].Validation(subBlock)
			BlockVal(subBlock, conf, Plugins)
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
