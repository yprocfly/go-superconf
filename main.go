package main

import (
	"encoding/json"
	"fmt"
	"github.com/yprocfly/go-superconf/superconfig"
)

type MySqlConfig struct {
	Default struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	}
}

func main() {
	var handle = new(superconfig.Handle)
	handle.GetConfig("/wk_risk/mysql")
	var MySql MySqlConfig
	handle.RegisterNode("/wk_risk/mysql", func(data []byte) {
		_ = json.Unmarshal(data, &MySql)
		fmt.Println(MySql, MySql.Default)
	})

	fmt.Println(superconfig.RegNodeList)
}
