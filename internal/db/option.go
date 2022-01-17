package db

import (
	"fmt"
	"strings"

	"go.uber.org/config"
)

// DBConf
type DBConf struct {
	Driver     string `default:"mysql"`
	DataSource string
	DBName     string
	UserName   string
	Password   string
	Host       string
	Port       int

	MaxIdleConns int
	MaxOpenConns int

	ProvideName string
}

// Datasource
func (c *DBConf) Datasource() string {
	if c.DataSource != "" {
		return c.DataSource
	}

	if c.Host == "" {
		c.Host = "localhost"
	}

	if c.Port == 0 {
		c.Port = 3306
	}

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		c.UserName,
		c.Password,
		fmt.Sprintf("%s:%d", c.Host, c.Port),
		c.DBName,
	)
}

// GetDBConf
func GetDBConf(path string) *DBConf {
	idx := strings.LastIndex(path, ".")
	if idx == -1 {
		panic(fmt.Sprintf("file error miss . path:%v", path))
	}

	switch path[idx+1:] {
	case "yaml", "yml":
		return yamlParse(path)
	default:
		panic(fmt.Sprintf("path error, path:%v", path))
	}
}

func yamlParse(path string) *DBConf {
	conf, err := config.NewYAML(config.File(path))
	if err != nil {
		panic(err)
	}

	var cv config.Value
	if cv = conf.Get("db"); !cv.HasValue() {
		return nil
	}

	var opts map[string]*DBConf
	if err := cv.Populate(&opts); err != nil {
		panic(err)
	}

	// TODO 暂时支持一个库读取
	if len(opts) != 1 {
		panic("db config not eq 1")
	}

	var cf *DBConf
	// TODO 实现有点丑 以后优化
	for k, v := range opts {
		cf = v
		cf.ProvideName = "db." + k
		return cf
	}

	if cf == nil {
		panic("impossiable cf empty")
	}
	return cf
}
