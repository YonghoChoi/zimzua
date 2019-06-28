package main

import (
	"fmt"
	"github.com/YonghoChoi/zimzua/pkg/db"
	uuid "github.com/satori/go.uuid"
	"strconv"
)

func main() {
	storages, err := SelectAll()
	if err != nil {
		panic(err.Error())
	}

	for _, v := range storages {
		v.BusinessDays = []BusinessDay{}
		days := []string{"Monday", "TuesDay", "Wednedsday", "Thursday", "Friday", "Saturday", "Sunday"}
		for _, d := range days {
			v.BusinessDays = append(v.BusinessDays, BusinessDay{
				Day:       d,
				StartTime: "10:00:00",
				EndTime:   "18:00:00",
			})
		}

		id, err := GetInstance().InsertStorage(v)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(id)
	}
}

func SelectAll() ([]*Storage, error) {
	query := "select name,phone,address,lat,lon from storage"

	rows, err := db.GetInstnace().SelectQuery(query)
	if err != nil {
		return nil, err
	}

	var storages []*Storage
	for rows.Next() {
		var name, phone, address, lat, lon string
		err := rows.Scan(
			&name,
			&phone,
			&address,
			&lat,
			&lon,
		)

		if err != nil {
			return nil, err
		}

		latFloat, err := strconv.ParseFloat(lat, 64)
		if err != nil {
			return nil, err
		}

		lonFloat, err := strconv.ParseFloat(lon, 64)
		if err != nil {
			return nil, err
		}

		storages = append(storages, &Storage{
			ID:      uuid.NewV4().String(),
			Name:    name,
			Phone:   phone,
			Address: address,
			Location: Location{
				Type:        "Point",
				Coordinates: []float64{lonFloat, latFloat},
			},
		})
	}

	return storages, nil
}
