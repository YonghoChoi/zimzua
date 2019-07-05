package storage

import (
	"context"
	"github.com/YonghoChoi/zimzua/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"sync"
	"github.com/YonghoChoi/zimzua/cmd/zimzua-api/config"
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
	c := config.GetInstance().Mongo
	db, err := db.NewMongoDB(&db.Config{Host: c.Host,
		Port: c.Port,
		Database: c.Database,
		Username: c.Username,
		Password: c.Password})
	if err != nil {
		panic(err)
	}
	o.db = db

	return o
}

func (o *MongoDB) InsertStorage(storage *Storage) (interface{}, error) {
	id, err := GetInstance().db.InsertOne(context.Background(), "storage", storage)
	if err != nil {
		return nil, err
	}

	return id, err
}

func (o *MongoDB) GetNearStorage(loc *Location, minDist, maxDist uint32) ([]Storage, error) {
	var results []Storage
	_, err := GetInstance().db.FindMany(
		context.Background(),
		"storage",
		bson.M{
			"location": bson.M{
				"$nearSphere": bson.M{
					"$geometry": bson.M{
						"type":        loc.Type,
						"coordinates": loc.Coordinates,
					},
					"$minDistance": minDist,
					"$maxDistance": maxDist,
				},
			},
		},
		nil,
		func(d bson.Raw) (interface{}, error) {
			var obj Storage
			if err := bson.Unmarshal(d, &obj); err != nil {
				return nil, err
			}

			results = append(results, obj)
			return obj, nil
		})

	if err != nil {
		return nil, err
	}

	return results, err
}
