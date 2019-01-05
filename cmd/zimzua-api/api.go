package main

import (
	"net/http"
	"zimzua/internal/packet"
	"zimzua/internal/code"
	"fmt"
	"log"
	"zimzua/internal/typedef"
)

func regUser(w http.ResponseWriter, r *http.Request) {
	res := packet.Res{Code: code.ResultOK}
	defer func() {
		log.Println(res.ToString())
		fmt.Fprintf(w, res.ToJson())
	}()

	accountInfo := new(typedef.AccountInfo)
	accountInfo.LoginType = r.FormValue("loginType")
	accountInfo.Name = r.FormValue("name")
	accountInfo.Phone = r.FormValue("phone")
	accountInfo.Email = r.FormValue("email")
	accountInfo.Password = r.FormValue("password")
	accountInfo.Token = r.FormValue("token")

	if err := accountInfo.ValidReg(); err != nil {
		res.Code = code.ResultInternalServerError
		res.Message = err.Error()
		return
	}

	if err := accountInfo.Insert(); err != nil {
		res.Code = code.ResultInternalServerError
		res.Message = "가입에 실패하였습니다."
		return
	}

	res.Message = "가입 되었습니다."
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	res := packet.Res{Code: code.ResultOK}
	defer func() {
		log.Println(res.ToString())
		fmt.Fprintf(w, res.ToJson())
	}()

	email := r.FormValue("email")
	password := r.FormValue("password")
	token := r.FormValue("token")

	accountInfo := new(typedef.AccountInfo)
	if err := accountInfo.Select(email); err != nil {
		res.Code = code.ResultInternalServerError
		res.Message = err.Error()
		return
	}

	if err := accountInfo.ValidLogin(password, token); err != nil {
		res.Code = code.ResultInternalServerError
		res.Message = err.Error()
		return
	}

	res.Message = "로그인 되었습니다."
}