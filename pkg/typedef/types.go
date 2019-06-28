package typedef

import (
	"fmt"
	"github.com/YonghoChoi/zimzua/pkg/db"
	"github.com/pkg/errors"
)

type AccountInfo struct {
	Id        int64  `json:id`
	Name      string `json:name`
	LoginType string `json:loginType`
	Phone     string `json:phone`
	Email     string `json:email`
	Password  string `json:password`
	Token     string `json:token`
}

func (a AccountInfo) ValidReg() error {
	if a.Name == "" {
		return fmt.Errorf("이름을 입력해주세요.")
	}

	if a.Email == "" {
		return fmt.Errorf("이메일을 입력해주세요.")
	}

	switch a.LoginType {
	case "zimzua":
		if a.Phone == "" {
			return fmt.Errorf("핸드폰 번호를 입력해주세요.")
		}

		if a.Password == "" {
			return fmt.Errorf("비밀번호를 입력해주세요.")
		}
	case "google":
		fallthrough
	case "facebook":
		if a.Token == "" {
			return fmt.Errorf("인증 정보가 유효하지 않습니다.")
		}
	default:
		return fmt.Errorf("유효하지 않은 로그인 유형입니다.")
	}

	return nil
}

func (a AccountInfo) ValidLogin(password string, token string) error {
	switch a.LoginType {
	case "zimzua":
		if a.Password != password {
			return fmt.Errorf("비밀번호가 일치하지 않습니다.")
		}
	case "google":
		fallthrough
	case "facebook":
		if a.Token != token {
			return fmt.Errorf("인증 정보가 유효하지 않습니다.")
		}
	default:
		return fmt.Errorf("유효하지 않은 로그인 유형입니다.")
	}

	return nil
}

func (a *AccountInfo) Insert() error {
	query := fmt.Sprintf("insert into account(name,phone,email,password,loginType,token) values('%s','%s','%s','%s','%s','%s')",
		a.Name, a.Phone, a.Email, a.Password, a.LoginType, a.Token)

	if id, err := db.Insert(query); err != nil {
		return errors.Wrap(err, "query : "+query)
	} else {
		a.Id = id
	}

	return nil
}

func (a *AccountInfo) Select(email string) error {
	query := fmt.Sprintf("select id,name,phone,email,password,loginType,token from account where email='%s'",
		email)

	if err := db.GetDB().QueryRow(query).Scan(
		&a.Id,
		&a.Name,
		&a.Phone,
		&a.Email,
		&a.Password,
		&a.LoginType,
		&a.Token,
	); err != nil {
		return errors.Wrap(err, "query : "+query)
	}

	return nil
}

type Point struct {
	Lon float64
	Lat float64
}

type Storage struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Location Point  `json:"location"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
	Dist     string `json:"dist"`
}

func (s Storage) Print() {
	fmt.Printf("%s\t%s\t%s\t%s\t%v\t%s\t%s\t%s\n", s.Id, s.Name, s.Phone, s.Address, s.Location, s.Created, s.Updated, s.Dist)
}
