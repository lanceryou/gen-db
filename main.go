package main

import (
	"flag"
	"github.com/lanceryou/gen-db/internal/dbr"
	"github.com/lanceryou/gen-db/internal/template"
	"os"
	"os/exec"
)

const (
	emptyTable = "empty table"
)

var cfg = flag.String("p", "./config/service.yml", "db config path")
var tableName = flag.String("t", emptyTable, "table name")
var ormType = flag.String("o", "craft", "orm type.")
var dbName = flag.String("d", "", "inject name or db name")

// flag 读取文件路径 文件内容
// 解析数据库配置
// 初始化数据库连接
// 读取数据表 列 字段 属性
// 模板生成代码到指定目录
func main() {
	flag.Parse()

	if *tableName == emptyTable {
		panic("miss table name please input -t=$tablename")
	}

	conf := dbr.GetDBConf(*cfg)
	attr := dbr.QueryDBAttr(conf, *tableName)
	code := template.Generate(attr)

	writeFile(code, *tableName+"_gen.go")
	cmd := exec.Command("go", "fmt", *tableName+"_gen.go")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func writeFile(content string, path string) {
	fileObj, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer fileObj.Close()
	contents := []byte(content)
	if _, err := fileObj.Write(contents); err != nil {
		panic(err)
	}
}
