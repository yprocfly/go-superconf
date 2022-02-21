# go-superconf

## 使用

```
1、在项目根目录新建一个superconf.json文件，文件内容是zk的配置，格式如下：
{
  "env": {
    "name": "env_name",
    "zookeeper": {
      "host": "127.0.0.1",
      "port": 2181,
      "auth_data": {
        "scheme": "digest",
        "auth": "user:password"
      }
    },
    "deploy": "dev",
    "group": "group_name"
  }
}
这样配置的路径前缀是 /superconf/group_name/env_name


2、安装包
go get github.com/yprocfly/go-superconf

代码实现
import "github.com/yprocfly/go-superconf/superconfig"

MysqlConfig := make(map[string]string)

superconfig.RegisterNode("/wk_risk/mysql", func(data []byte) {
    // 这里可以按照实际的mysql配置，来定义MysqlConfig数据结构
    _ = json.Unmarshal(data, &MysqlConfig)
    initMysqlConnect(MysqlConfig)
})
```
