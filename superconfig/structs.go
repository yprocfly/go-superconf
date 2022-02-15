package superconfig

type AuthData struct {
	Scheme string `json:"scheme"`
	Auth   string `json:"auth"`
}

type ZooKeeper struct {
	Host     string   `json:"host"`
	Port     int      `json:"port"`
	AuthData AuthData `json:"auth_data"`
}

type Env struct {
	Name      string    `json:"name"`
	Zookeeper ZooKeeper `json:"zookeeper"`
	Deploy    string    `json:"deploy"`
	Group     string    `json:"group"`
}

type Config struct {
	Env Env `json:"env"`
}

type RegNode struct {
	NodePath   string
	HandleFunc func(data []byte)
}
