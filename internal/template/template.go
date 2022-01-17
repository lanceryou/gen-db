package template

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

var (
	templateMap = make(map[string]Template)
)

// Register 注册模板
func Register(t Template) {
	templateMap[t.String()] = t
}

// DRegister 反注册模板
func DRegister(t string) {
	delete(templateMap, t)
}

// Template orm template
type Template interface {
	ReplaceTemplate() string
	MockTemplate() string
	String() string
}

// 根据template 生成代码
func Generate(src interface{}, orm string) string {
	tmpl, ok := templateMap[orm]
	if !ok {
		panic(fmt.Errorf("orm not register %v", orm))
	}
	return replaceTemplate(src, tmpl.ReplaceTemplate())
}

func MockTemplate(src interface{}, orm string) string {
	tmpl, ok := templateMap[orm]
	if !ok {
		panic(fmt.Errorf("orm not register %v", orm))
	}

	mock := replaceTemplate(src, tmpl.MockTemplate())
	mock = mockPlaceholder(mock, "// holder start", "// holder end", "placeholder")
	mock = mockPlaceholder(mock, "// fields start", "// fields end", "fields-holder")
	return mock
}

func mockPlaceholder(mock string, start, end string, old string) string {
	startIdx := strings.Index(mock, start)
	endIdx := strings.Index(mock, end)
	placeholder := mock[startIdx:endIdx]
	mock = strings.Replace(mock, placeholder+end, "", -1)
	return strings.Replace(mock, old, placeholder, -1)
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
