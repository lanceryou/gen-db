package main

import (
	"flag"
	"os"
	"os/exec"

	"github.com/lanceryou/gen-db/internal/db"
	"github.com/lanceryou/gen-db/template"
)

const (
	emptyTable = "empty table"
)

var path = flag.String("path", "./config/service.yml", "db config path")
var tableName = flag.String("t", emptyTable, "table name")
var ormType = flag.String("o", "xorm", "orm type.")
var datasource = flag.String("d", "", "datasource")
var injectName = flag.String("i", "", "inject name")
var packageName = flag.String("p", "model", "package name")

func main() {
	flag.Parse()

	if *tableName == emptyTable {
		panic("miss table name please input -t=$tablename")
	}

	attr := db.QueryDBAttr(getConn(*path, *datasource, *injectName), *tableName, *packageName)
	code := template.Generate(attr, *ormType)
	mock := template.MockTemplate(attr, *ormType)
	writeFile(mock, *tableName+"_gen_test.go")
	writeFile(code, *tableName+"_gen.go")
	cmd := exec.Command("go", "fmt", *tableName+"_gen.go", *tableName+"_gen_test.go")
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

func getConn(path string, datasource string, injectName string) *db.Connection {
	if datasource != "" {
		return db.Open(datasource, injectName)
	}

	conf := db.GetDBConf(path)
	return db.Open(conf.Datasource(), conf.ProvideName)
}
