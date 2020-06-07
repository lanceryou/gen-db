package dbr

import (
	"time"

	"github.com/gocraft/dbr/dialect"
)

type mysql struct{}

const (
	timeFormat = "2006-01-02 15:04:05.000000"
)

func (d mysql) QuoteIdent(s string) string {
	return dialect.MySQL.QuoteIdent(s)
}

func (d mysql) EncodeString(s string) string {
	return dialect.MySQL.EncodeString(s)
}

func (d mysql) EncodeBool(b bool) string {
	return dialect.MySQL.EncodeBool(b)
}

func (d mysql) EncodeTime(t time.Time) string {
	return `'` + t.Format(timeFormat) + `'`
}

func (d mysql) EncodeBytes(b []byte) string {
	return dialect.MySQL.EncodeBytes(b)
}

func (d mysql) Placeholder(n int) string {
	return dialect.MySQL.Placeholder(n)
}
