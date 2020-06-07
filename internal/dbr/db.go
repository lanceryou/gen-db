package dbr

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

type Connection struct {
	*dbr.Connection
	Name string
}

func (c *Connection) NewSession() *dbr.Session {
	return c.Connection.NewSession(nil)
}

func Open(option *DBConf) *Connection {
	if option.Host == "" {
		option.Host = "localhost"
	}

	if option.Port == 0 {
		option.Port = 3306
	}

	option.DataSource = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		option.UserName,
		option.Password,
		fmt.Sprintf("%s:%d", option.Host, option.Port),
		option.DBName,
	)

	fmt.Printf("data%v\n", option.DataSource)
	if option.Driver == "" {
		option.Driver = "mysql"
	}
	conn, err := dbr.Open(option.Driver, option.DataSource, &sqlEventReceiver{})
	if err != nil {
		panic(err)
	}
	conn.Dialect = &mysql{}
	conn.SetMaxIdleConns(option.MaxIdleConns)
	conn.SetMaxOpenConns(option.MaxOpenConns)
	return &Connection{Connection: conn, Name: option.ProvideName}
}
