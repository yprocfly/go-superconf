package utils

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

// ZkConn 定义zk连接对象
var ZkConn = &zk.Conn{}

func init() {
	hosts := []string{
		SuperConfig.Env.Zookeeper.Host,
	}
	// 连接zk
	conn, _, err := zk.Connect(hosts, time.Second*5)
	// defer conn.Close()
	if err != nil {
		panic(err)
	}

	ZkConn = conn
}
