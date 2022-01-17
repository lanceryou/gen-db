package conf

import (
	"fmt"
	"strings"

	"go.uber.org/config"
)

// NodeConf 数据库节点配置
type NodeConf struct {
	DataSource string
	DBName     string
	UserName   string
	Password   string
	Host       string
	Port       int

	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifeTime int64

	Slaves map[string]*DBConf
}

// DBConf db 配置
type DBConf struct {
	DataSource string
	DBName     string
	UserName   string
	Password   string
	Host       string
	Port       int

	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifeTime int64
}

// GetDBConf
func GetDBConf(path string, key string) *NodeConf {
	idx := strings.LastIndex(path, ".")
	if idx == -1 {
		panic(fmt.Sprintf("file error miss . path:%v", path))
	}

	switch path[idx+1:] {
	case "yaml", "yml":
		return yamlParse(path, key)
	default:
		panic(fmt.Sprintf("path error, path:%v", path))
	}
}

func yamlParse(path string, key string) *NodeConf {
	cf, err := config.NewYAML(config.File(path))
	if err != nil {
		panic(err)
	}

	var cv config.Value
	if cv = cf.Get(key); !cv.HasValue() {
		return nil
	}

	var nc NodeConf
	if err := cv.Populate(&nc); err != nil {
		panic(err)
	}

	return &nc
}
