package auth

import (
	"fmt"
	"github.com/YonghoChoi/zimzua/model/account"
	"github.com/YonghoChoi/zimzua/pkg/code"
	"github.com/YonghoChoi/zimzua/pkg/packet"
	"log"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
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

	obj, err := account.GetInstance().GetAccount(map[string]string{
		"email": email,
	})

	if err != nil {
		res.Code = code.ResultInternalServerError
		res.Message = err.Error()
		log.Println("select account fail. err : ", err.Error())
		return
	}

	if err := obj.ValidLogin(password, token); err != nil {
		res.Code = code.ResultInternalServerError
		res.Message = err.Error()
		log.Println("invalid login. err : ", err.Error())
		return
	}

	res.Message = "로그인 되었습니다."
}
