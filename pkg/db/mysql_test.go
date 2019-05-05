package db

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	query := fmt.Sprintf("insert into account(name,phone,email,password,loginType,token) values('%s','%s','%s','%s','%s','%s')",
		"yongho", "010", "yongho1037@gmail.com", "testpassword", "zimzua", "$#%$%RFSDFDF")

	t.Log(query)
	id, err := Insert(query)
	if err != nil {
		t.Error(err)
	}

	t.Log(id)
}

func TestSelectQuery(t *testing.T) {
	query := fmt.Sprintf("select name,phone,email,password,loginType,token from account where email='%s'", "yongho1037@gmail.com")
	t.Log(query)

	var name, phone, email, password, loginType, token string
	db := GetInstnace().getDB()
	err := db.QueryRow(query).Scan(
		&name,
		&phone,
		&email,
		&password,
		&loginType,
		&token,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(name, phone, email, password, loginType, token)
}
