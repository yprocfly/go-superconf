package superconfig

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"path"
)

var RegNodeMap = make(map[string]func(data []byte))

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

	isExists := IsPathExists(fullPath)
	if !isExists {
		panic(fmt.Sprintf("zk节点不【%s】存在", fullPath))
	}

	// 处理方法与节点绑定
	RegNodeMap[fullPath] = handleFunc

	// 获取zk配置，执行被装饰方法
	zkData := GetConfigAndWatch(fullPath)
	handleFunc(zkData)
}

// IsPathExists 判断zk节点路径是否存在
// @param fullPath zk节点完整路径
func IsPathExists(fullPath string) bool {
	exists, _, err := ZkConn.Exists(fullPath)
	if err != nil {
		panic(fmt.Sprintf("判断zk节点【%s】是否存在时出错：%s", fullPath, err))
	}

	return exists
}

// GetConfigAndWatch 获取配置，并监听节点
// @param fullPath zk节点路径，去掉前缀部分
func GetConfigAndWatch(fullPath string) []byte {
	data, _, event, err := ZkConn.GetW(fullPath)
	if err != nil {
		panic(fmt.Sprintf("获取节点【%s】出错：%s", fullPath, err.Error()))
	}

	// 协程调用监听事件
	go watchPathContentChange(event)

	return data
}

// 监听节点内容变更
func watchPathContentChange(e <-chan zk.Event) {
	event := <-e
	data := GetConfigAndWatch(event.Path)

	// 调用注册的方法 - 热更
	RegNodeMap[event.Path](data)
}
