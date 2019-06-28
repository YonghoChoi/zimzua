package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

type MySQLWrap struct {
	Ip       string
	Port     int
	DBName   string
	User     string
	Password string
}

func (s *MySQLWrap) getDB() *sql.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", s.User, s.Password, s.Ip, s.Port, s.DBName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (s *MySQLWrap) insert(query string) (int64, error) {
	db := s.getDB()
	defer db.Close()
	res, err := db.Exec(query)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (s *MySQLWrap) update(query string) (int64, error) {
	db := s.getDB()
	defer db.Close()
	res, err := db.Exec(query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (s *MySQLWrap) delete(query string) (int64, error) {
	db := s.getDB()
	defer db.Close()
	res, err := db.Exec(query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (s *MySQLWrap) selectQuery(query string) (*sql.Rows, error) {
	db := s.getDB()
	return db.Query(query)
}

var (
	instance *MySQLWrap
	once     sync.Once
)

func GetInstnace() *MySQLWrap {
	// TODO : yaml 파일로 추출 예정
	once.Do(func() {
		instance = new(MySQLWrap)
		instance.Ip = "localhost"
		instance.Port = 13306
		instance.User = "root"
		instance.Password = "root"
		instance.DBName = "zimzua"
	})

	return instance
}

func Insert(query string) (int64, error) {
	return GetInstnace().insert(query)
}

func Update(query string) (int64, error) {
	return GetInstnace().update(query)
}

func Delete(query string) (int64, error) {
	return GetInstnace().delete(query)
}

func SelectQuery(query string) (*sql.Rows, error) {
	return GetInstnace().selectQuery(query)
}

func GetDB() *sql.DB {
	return GetInstnace().getDB()
}
