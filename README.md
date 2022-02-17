# go-superconf

## 使用

```
安装包
go get github.com/yprocfly/go-superconf

代码实现
import "github.com/yprocfly/go-superconf/superconfig"

superconfig.RegisterNode("/wk_risk/mysql", func(data []byte) {
    _ = json.Unmarshal(data, &MysqlConfig)
})
```
