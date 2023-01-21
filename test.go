package main

import (
	"fmt"
	"plugin"

	"golang.org/x/exp/maps"
)

func main() {
	maps.Keys(map[string]string{"sfd": "w", "k": "ed"})
	p, _ := plugin.Open("/Users/natsaiso/go/src/pivot-g/pivot/plugins/auth/renew.so")
	fmt.Println(p.Lookup("Renew"))
}
