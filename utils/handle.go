package utils

import "fmt"

type Handle struct{}

func (handle *Handle) GetConfig(path string) {
	fullPath := PrefixPath + path
	data, _, err := ZkConn.Get(fullPath)
	if err != nil {
		panic("节点【%s】不存在")
	}
	fmt.Printf("%s 的值为 %s\n", fullPath, string(data))
}
