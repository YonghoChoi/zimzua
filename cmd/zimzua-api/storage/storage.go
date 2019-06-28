package storage

import (
	"fmt"
	"github.com/YonghoChoi/zimzua/model/storage"
	"github.com/YonghoChoi/zimzua/pkg/code"
	"github.com/YonghoChoi/zimzua/pkg/packet"
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

	storages, err := storage.GetInstance().GetNearStorage(
		&storage.Location{
			Type:        "Point",
			Coordinates: []float64{lon, lat},
		},
		1000,
		12000, // 12000 = 1.2km
	)

	if err != nil {
		res.Code = code.ResultInternalServerError
		res.Message = "오류가 발생하였습니다."
		log.Printf("fail get near storage. err : %s\n", err.Error())
		return
	}

	res.AddData("storageList", storages)
}
