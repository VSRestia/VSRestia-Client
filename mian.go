package main

import (
	"fmt"
	"github.com/VSRestia/VSRestia-Client/config"
	"github.com/VSRestia/VSRestia-Client/core"
)

func main() {
	core.Hello()
	cfg, err := config.CheckConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	core.Run(cfg)
}
