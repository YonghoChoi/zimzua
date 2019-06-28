package storage

import (
	"fmt"
	"github.com/YonghoChoi/zimzua/pkg/code"
	"github.com/YonghoChoi/zimzua/pkg/db"
	"github.com/YonghoChoi/zimzua/pkg/packet"
	"github.com/YonghoChoi/zimzua/pkg/typedef"
	"log"
	"net/http"
	"strconv"
)

// Sample Data : 디캠프 좌표 (lon = 127.043695, lat = 37.5084632)
// Sample Data : 종로3가 좌표 (lon = 126.9895646721, lat = 37.5702449756)
// call GetStorageList(POINT(126.9895646721,37.5702449756))
func GetStorageList(w http.ResponseWriter, r *http.Request) {
	res := packet.Res{Code: code.ResultOK}
	defer func() {
		log.Println(res.ToString())
		fmt.Fprintf(w, res.ToJson())
	}()

	latStr := r.FormValue("lat")
	lonStr := r.FormValue("lon")
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		res.Code = code.ResultInternalServerError
		res.Message = "잘못 입력하셨습니다."
		log.Println("invalid param(lat). err : ", err.Error())
		return
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		res.Code = code.ResultInternalServerError
		res.Message = "잘못 입력하셨습니다."
		log.Println("invalid parma(lon). err : ", err.Error())
		return
	}

	query := fmt.Sprintf(`call GetStorageList(POINT(%f,%f))`, lon, lat)
	rows, err := db.SelectQuery(query)
	if err != nil {
		res.Code = code.ResultInternalServerError
		res.Message = "오류가 발생하였습니다."
		log.Printf("fail query. query : %s, err : %s\n", query, err.Error())
		return
	}
	defer rows.Close()

	var storageList []*typedef.Storage
	for rows.Next() {
		var location, lon, lat string
		storage := new(typedef.Storage)
		err := rows.Scan(
			&storage.Id,
			&storage.Name,
			&storage.Phone,
			&storage.Address,
			&location,
			&lon,
			&lat,
			&storage.Created,
			&storage.Updated,
			&storage.Dist)
		storage.Location.Lon, err = strconv.ParseFloat(lon, 64)
		if err != nil {
			log.Println(err)
		}
		storage.Location.Lat, err = strconv.ParseFloat(lat, 64)
		if err != nil {
			log.Println(err)
		}

		storage.Print()
		storageList = append(storageList, storage)
	}

	res.AddData("storageList", storageList)
}
