package utils

import (
	"encoding/json"
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

// SuperConfig 保存zk的相关配置信息
var SuperConfig Config

// PrefixPath zk节点路径前缀
var PrefixPath string

// 初始化路径前缀
func initPrefixPath(superConfig Config) {
	PrefixPath = "/superconf/" + superConfig.Env.Name + "/" + superConfig.Env.Group
}

func init() {
	// 获取 superconf 配置的路径，设置在项目的根目录下，文件名为 superconf.json
	rootPath, _ := os.Getwd()
	filePath := path.Join(rootPath, "superconf.json")
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		panic("读取superconf配置失败：" + err.Error())
	}
	err = json.Unmarshal(fileContent, &SuperConfig)
	if err != nil {
		panic("解析superconf配置失败：" + err.Error())
	}

	initPrefixPath(SuperConfig)
}
