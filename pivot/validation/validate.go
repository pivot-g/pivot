package validation

import (
	"fmt"

	"github.com/pivot-g/pivot/pivot/config"
)

func Validate() {
	conf := &config.ConfigDir{
		Dir: "/Users/natsaiso/go/src/pivot-g/pivot/example/config",
	}
	fmt.Println(conf.LoadConfig())

}
