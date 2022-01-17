package dbr

// Xorm
type Dbr struct {
}

// String xorm 实现
func (x *Dbr) String() string {
	return "dbr"
}

// ReplaceTemplate xorm 模板
func (x *Dbr) ReplaceTemplate() string {
	return tmpl
}

// MockTemplate mock 模板
func (x *Dbr) MockTemplate() string {
	return mockTmpl
}
