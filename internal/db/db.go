package db

import (
	// 引入mysql包
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

// Connection
type Connection struct {
	*dbr.Connection
	Name string
}

// NewSession
func (c *Connection) NewSession() *dbr.Session {
	return c.Connection.NewSession(nil)
}

// Open
func Open(dataSource string, providerName string) *Connection {
	conn, err := dbr.Open("mysql", dataSource, &dbr.NullEventReceiver{})
	if err != nil {
		panic(err)
	}
	return &Connection{Connection: conn, Name: providerName}
}
