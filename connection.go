package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

type Connection struct {
	Driver string
	Dsn    string
}

func (c *Connection) Connect() (err error) {
	db, err = gorm.Open(c.Driver, c.Dsn)
	return
}

func (c *Connection) Close() {
	db.Close()
}
