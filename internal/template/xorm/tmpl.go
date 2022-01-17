package xorm

var tmpl = `// Code generated by gen-db. DO NOT EDIT.

package {{.Package}}

import (
	"github.com/lanceryou/micro/db/xormcluster"
	"time"
)

type {{.LowTable}}Model interface {
	BatchInsert(entity ...*{{.Table}}Entity) error
}

// {{.Table}}Entity 数据库表entity
type {{.Table}}Entity struct {
	{{range .TableFields}}{{.CamelField}} {{.FieldType}} {{$.Inject}}xorm:"{{.Field}}"{{$.Inject}}
	{{end}}
	updateInfo map[string]interface{}
}

const (
	table{{.Table}} = "{{.TableName}}"
)

var (
	_ {{.LowTable}}Model = &{{.LowTable}}Impl{}
)
type {{.LowTable}}Impl struct {
	*xormcluster.DB {{.Inject}}inject:"{{.ProviderName}}"{{.Inject}}
}

func (e *{{.Table}}Entity) set(k string, v interface{}) {
	if e.updateInfo == nil {
		e.updateInfo = map[string]interface{}{}
	}

	e.updateInfo[k] = v
}

{{range .TableFields}}
// Set{{.CamelField}} 设置{{.CamelField}}
func (e *{{$.Table}}Entity) Set{{.CamelField}}(v interface{}) *{{$.Table}}Entity{
	e.set("{{.Field}}", v)
	return e
}
{{end}}
// 生成rows
var columns{{.Table}}Fields = []string {
{{range .TableFields}} 	"{{.Field}}",
{{end}}
}

type columns{{.Table}}Type struct {
{{range .TableFields}} 	{{.CamelField}} string
{{end}}
}

var columns{{.Table}} = columns{{.Table}}Type{
{{range .TableFields}} 	{{.CamelField}} : "{{.Field}}",
{{end}}
}
// BatchInsert 批量插入
func (e *{{.LowTable}}Impl) BatchInsert(entity ...*{{.Table}}Entity) error {
	_, err := e.Table(table{{.Table}}).Insert(entity)
	return err
}
`