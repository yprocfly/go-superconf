package superconf

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type AuthData struct {
	Scheme string
	Auth   string
}

type ZooKeeper struct {
	Host     string
	Port     int
	AuthData AuthData
}

type Env struct {
	Name      string
	Zookeeper ZooKeeper
	Deploy    string
	Group     string
}

type Config struct {
	Env Env
}

func init() {
	var SuperConfig Config

	rootPath, _ := os.Getwd()
	filePath := path.Join(rootPath, "superconf.json")
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取superconf配置失败", err)
		return
	}
	err = json.Unmarshal(fileContent, &SuperConfig)
	if err != nil {
		fmt.Println("解析superconf配置失败", err)
		return
	}
}
