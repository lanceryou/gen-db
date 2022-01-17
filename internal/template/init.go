package template

import (
	"github.com/lanceryou/gen-db/internal/template/xorm"
)

func init() {
	Register(&xorm.Xorm{})
}
