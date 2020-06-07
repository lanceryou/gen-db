package template

import (
	"bytes"
	"text/template"
)

// 根据template 生成代码
func Generate(src interface{}) string {
	return replaceTemplate(src, tmpl)
}

func GenerateTmpl(src interface{}, tmpl string) string {
	return replaceTemplate(src, tmpl)
}

func replaceTemplate(src interface{}, tmpl string) string {
	tml := template.New("model")
	var err error
	tml, err = tml.Parse(tmpl)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = tml.Execute(&buf, src)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
