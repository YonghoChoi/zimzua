package account

import "fmt"

type Account struct {
	ID        string `json:"id" bson:"_id"`
	Name      string `json:"name" bson:"name"`
	Phone     string `json:"phone" bson:"phone"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	LoginType string `json:"login_type" bson:"login_type"`
	Token     string `json:"token" bson:"token"`
	Created   string `json:"created,omitempty" bson:"created,omitempty"`
	Updated   string `json:"updated,omitempty" bson:"updated,omitempty"`
}

func (a Account) ValidReg() error {
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

func (a Account) ValidLogin(password string, token string) error {
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
