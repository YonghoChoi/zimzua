package storage

import (
	"fmt"
	"github.com/satori/go.uuid"
	"testing"
)

func TestMongoDB_InsertOne(t *testing.T) {
	storages := []Storage{
		Storage{ID: uuid.NewV4().String(),
			Name:     "제주 커피 박물관",
			Address:  "제주 커피 박물관",
			Location: Location{Type: "Point", Coordinates: []float64{126.8990639, 33.4397954}},
			Phone:    "010-000-0000",
		},
		Storage{ID: uuid.NewV4().String(),
			Name:     "신산리 마을카페",
			Address:  "신산리 마을카페",
			Location: Location{Type: "Point", Coordinates: []float64{126.8759347, 33.3765812}},
			Phone:    "010-000-0000",
		},
	}

	for _, v := range storages {
		id, err := GetInstance().InsertStorage(&v)
		if err != nil {
			t.Fatal(err.Error())
		}

		fmt.Println(id)
	}

	//find
	results, err := GetInstance().GetNearStorage(&Location{Type: "Point", Coordinates: []float64{126.941131, 33.459216}}, 1000, 12000)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(fmt.Sprintf("%v", results))
}
