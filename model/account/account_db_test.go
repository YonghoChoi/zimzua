package account

import (
	"fmt"
	"github.com/satori/go.uuid"
	"testing"
)

func TestMongoDB_InsertOne(t *testing.T) {
	id, err := GetInstance().InsertAccount(&Account{
		ID:    uuid.NewV4().String(),
		Name:  "test1",
		Email: "yongho1037@gmail.com",
		Token: "abcd",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	account, err := GetInstance().GetAccount(map[string]string{
		"_id": id.(string),
	})

	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(fmt.Sprintf("%v", account))
}
