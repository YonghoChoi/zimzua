package main

import (
	"net/http"
	"zimzua/internal/packet"
	"zimzua/internal/code"
	"fmt"
	"log"
	"zimzua/internal/typedef"
	"strconv"
	"zimzua/pkg/db"
)

func regUser(w http.ResponseWriter, r *http.Request) {
	res := packet.Res{Code: code.ResultOK}
	defer func() {
		log.Println(res.ToString())
		fmt.Fprintf(w, res.ToJson())
	}()

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//r.ParseForm()
	loginType := r.Form.Get("loginType")
	name := r.Form.Get("name")
	phone := r.Form.Get("phone")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	token := r.Form.Get("token")
	log.Println(fmt.Sprintf("Request regUser. loginType : %s, name : %s, phone : %s, email : %s, password : %s, token : %s",
		loginType, name, phone, email, password, token))

	accountInfo := new(typedef.AccountInfo)
	accountInfo.LoginType = loginType
	accountInfo.Name = name
	accountInfo.Phone = phone
	accountInfo.Email = email
	accountInfo.Password = password
	accountInfo.Token = token

	if err := accountInfo.ValidReg(); err != nil {
		res.Code = code.ResultInternalServerError
		res.Message = err.Error()
		log.Println("reg fail. err : ", err.Error())
		return
	}

	if err := accountInfo.Insert(); err != nil {
		res.Code = code.ResultInternalServerError
		res.Message = "가입에 실패하였습니다."
		log.Println("insert fail. err : ", err.Error())
		return
	}

	res.Message = "가입 되었습니다."
	log.Println(fmt.Sprintf("registed user. %v", accountInfo))
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	res := packet.Res{Code: code.ResultOK}
	defer func() {
		log.Println(res.ToString())
		fmt.Fprintf(w, res.ToJson())
	}()

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")
	token := r.Form.Get("token")

	accountInfo := new(typedef.AccountInfo)
	if err := accountInfo.Select(email); err != nil {
		res.Code = code.ResultInternalServerError
		res.Message = err.Error()
		log.Println("select account fail. err : ", err.Error())
		return
	}

	if err := accountInfo.ValidLogin(password, token); err != nil {
		res.Code = code.ResultInternalServerError
		res.Message = err.Error()
		log.Println("invalid login. err : ", err.Error())
		return
	}

	res.Message = "로그인 되었습니다."
}

func getStorageList(w http.ResponseWriter, r *http.Request) {
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

	query := fmt.Sprintf(`SELECT *, ST_DISTANCE_SPHERE(POINT(%f, %f), location) AS dist FROM storage`, lon, lat)
	rows, err := db.SelectQuery(query)
	if err != nil {
		res.Code = code.ResultInternalServerError
		res.Message = "오류가 발생하였습니다."
		log.Printf("fail query. query : %s, err : %s\n", query, err.Error())
		return
	}
}