package template

import (
	"github.com/lanceryou/gen-db/internal/template/dbr"
	"github.com/lanceryou/gen-db/internal/template/xorm"
	"github.com/lanceryou/gen-db/template"
)

func init() {
	template.Register(&xorm.Xorm{})
	template.Register(&dbr.Dbr{})
}
