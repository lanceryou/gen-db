package xorm

// Xorm
type Xorm struct {
}

// String xorm 实现
func (x *Xorm) String() string {
	return "xorm"
}

// ReplaceTemplate xorm 模板
func (x *Xorm) ReplaceTemplate() string {
	return tmpl
}

// MockTemplate mock 模板
func (x *Xorm) MockTemplate() string {
	return mockTmpl
}
