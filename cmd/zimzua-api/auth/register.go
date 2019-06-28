package auth

import (
	"fmt"
	"github.com/YonghoChoi/zimzua/pkg/code"
	"github.com/YonghoChoi/zimzua/pkg/packet"
	"github.com/YonghoChoi/zimzua/pkg/typedef"
	"log"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
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
