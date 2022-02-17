package superconfig

import (
	"fmt"
	"path"
)

var RegNodeList []RegNode

// GetFullPath 获取zk节点完整路径
// @param zkPath zk节点路径，去掉前缀部分
func GetFullPath(zkPath string) string {
	return path.Join(PrefixZkPath, zkPath)
}

// RegisterNode 注册节点【装饰器模式】
// @param zkPath zk节点路径，去掉前缀部分
// @param handleFunc 被装饰的方法，一般用于把zk配置赋值
func RegisterNode(zkPath string, handleFunc func(data []byte)) {
	fullPath := GetFullPath(zkPath)

	isExists := IsPathExists(zkPath)
	if !isExists {
		panic(fmt.Sprintf("zk节点不【%s】存在", fullPath))
	}

	// 获取zk配置，执行被装饰方法
	zkData := GetConfig(zkPath)
	handleFunc(zkData)

	// TODO 节点监听

	//RegNodeList = append(RegNodeList, RegNode{
	//	fullPath,
	//	handleFunc,
	//})
}

// IsPathExists 判断zk节点路径是否存在
// @param zkPath zk节点路径，去掉前缀部分
func IsPathExists(zkPath string) bool {
	fullPath := GetFullPath(zkPath)

	exists, _, err := ZkConn.Exists(fullPath)
	if err != nil {
		panic(fmt.Sprintf("判断zk节点【%s】是否存在时出错：%s", fullPath, err))
	}

	return exists
}

// GetConfig 获取配置
// @param zkPath zk节点路径，去掉前缀部分
func GetConfig(zkPath string) []byte {
	fullPath := GetFullPath(zkPath)

	data, _, err := ZkConn.Get(fullPath)
	if err != nil {
		panic(fmt.Sprintf("获取节点【%s】出错：%s", fullPath, err.Error()))
	}

	return data
}
