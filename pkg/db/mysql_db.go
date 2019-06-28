package db

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"sync"
)

type MySQLDB struct {
	db *xorm.Engine
}

var inst *MySQLDB
var so sync.Once

func GetInstance() *MySQLDB {
	so.Do(func() {
		inst = NewMySQLDB()
	})

	return inst
}

func NewMySQLDB() (hostDB *MySQLDB) {
	o := new(MySQLDB)
	dbConn := "root:@tcp(127.0.0.1:3306)/wordpress?charset=utf8&parseTime=True"
	db, err := xorm.NewEngine("mysql", dbConn)
	if err != nil {
		panic(err)
	}

	o.db = db
	return o
}

func (o *MySQLDB) InsertStorage() error {
	if o.db == nil {
		return fmt.Errorf("db is nil")
	}

	//if _, err := o.db.Insert(obj); err != nil {
	//    return nil, err
	//}

	return nil
}
