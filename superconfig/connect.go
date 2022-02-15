package superconfig

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

// ZkConn 定义zk连接对象
var ZkConn = &zk.Conn{}

// GetConn 获取zk连接
func GetConn() *zk.Conn {
	hosts := []string{
		fmt.Sprintf("%s:%d", SuperConfig.Env.Zookeeper.Host, SuperConfig.Env.Zookeeper.Port),
	}
	// 连接zk
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		panic(err)
	}

	// 设置权限
	authData := SuperConfig.Env.Zookeeper.AuthData
	err = conn.AddAuth(authData.Scheme, []byte(authData.Auth))
	if err != nil {
		panic(err)
	}

	return conn
}

func init() {
	ZkConn = GetConn()
}
