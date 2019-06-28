package account

import (
	"context"
	"github.com/YonghoChoi/zimzua/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"sync"
)

// MongoDB 정의
type MongoDB struct {
	db *db.MongoDB
}

var mongoInst *MongoDB
var mongoOnce sync.Once

func GetInstance() *MongoDB {
	mongoOnce.Do(func() {
		mongoInst = NewMongoDB()
	})

	return mongoInst
}

// NewMongoDB 새로운 MongoDB 인스턴스 생성
func NewMongoDB() *MongoDB {
	o := new(MongoDB)
	db, err := db.NewMongoDB(&db.Config{Host: "localhost", Port: 27018, Database: "zimzua", Username: "root", Password: "root"})
	if err != nil {
		panic(err)
	}
	o.db = db

	return o
}

func (o *MongoDB) InsertAccount(obj *Account) (interface{}, error) {
	id, err := GetInstance().db.InsertOne(context.Background(), "account", obj)
	if err != nil {
		return nil, err
	}

	return id, err
}

func MapToFilter(param map[string]string) bson.M {
	filter := bson.M{}
	for k, v := range param {
		filter[k] = v
	}

	return filter
}

func (o *MongoDB) GetAccount(param map[string]string) (*Account, error) {
	filter := MapToFilter(param)
	result, err := GetInstance().db.FindOne(context.Background(),
		"account",
		filter,
		func(d bson.Raw) (interface{}, error) {
			obj := new(Account)
			if err := bson.Unmarshal(d, obj); err != nil {
				return nil, err
			}

			return obj, nil
		})
	return result.(*Account), err
}
