package superconfig

import (
	"encoding/json"
	"os"
	"path"
)

// SuperConfig 保存zk的相关配置信息
var SuperConfig Config

// PrefixPath zk节点路径前缀
var PrefixPath string

// 获取 superconf 配置
// 配置文件保存在项目的根目录下，文件名为 superconf.json
func getSuperConfig() {
	// 获取 superconf 配置的路径
	rootPath, _ := os.Getwd()
	filePath := path.Join(rootPath, "superconf.json")

	// 读取文件内容
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		panic("读取superconf配置失败：" + err.Error())
	}

	// json序列化
	err = json.Unmarshal(fileContent, &SuperConfig)
	if err != nil {
		panic("解析superconf配置失败：" + err.Error())
	}
}

// 初始化路径前缀
func initPrefixPath() {
	PrefixPath = "/superconf/" + SuperConfig.Env.Name + "/" + SuperConfig.Env.Group
}

func init() {
	getSuperConfig()
	initPrefixPath()
}
