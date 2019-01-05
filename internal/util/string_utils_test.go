package util

import (
	"strings"
	"fmt"
	"testing"
	"os"
)

func TestRemoveFlow(t *testing.T) {
	stringMap := make(map[string]interface{})
	stringMap["hostname:repopath1"] = "test1"
	stringMap["hostname:repopath2"] = "test2"
	stringMap["hostname:repopath3"] = "test3"
	stringMap["hostname:repopath4"] = "test4"

	keys := MapKeys(stringMap)
	hostName := "hostname"

	// 변경된 설정의 SVN 정보 리스트업
	data := "SVN_DIR\trepopath1\nSVN_DIR\trepopath3\nSVN_DIR\trepopath5\nSVN_DIR\trepopath7\nSVN_DIR\trepopath9\n"
	var addSvn []string
	items := strings.Split(data, "\n")
	for _, line := range items {
		tmp := strings.Split(line, "\t")
		if tmp[0] != "SVN_DIR" {
			continue
		}

		key := fmt.Sprintf("%s:%s", hostName, tmp[1])

		if Include(keys, key) { // 설정 항목에 기존 항목이 포함되어 있으면 유지
			keys = RemoveFast(keys, key)
		} else { // 포함되어 있지 않으면 추가
			addSvn = append(addSvn, key)
		}
	}

	t.Log("== remove key ==")
	// 이제 keys에는 설정 변경으로 인해 제거된 항목들만 존재함. Redis에서도 해당 정보 제거
	for _, removeKey := range keys {
		t.Log(removeKey)
	}

	t.Log("== add key ==")
	// 추가될 항목은 redis에 추가
	for _, addKey := range addSvn {
		t.Log(addKey)
	}
}

func TestRemoveWhiteSpace(t *testing.T) {
	str := "this is test string"
	t.Log(RemoveWhiteSpace(str))

	str = "this      is   test              string"
	t.Log(RemoveWhiteSpace(str))

	str = "this\tis\ttest\tstring"
	t.Log(RemoveWhiteSpace(str))

	str = "this\nis\ntest\nstring"
	t.Log(RemoveWhiteSpace(str))
}

func TestTrimSpecialCharacter(t *testing.T) {
	test := TrimSpecialCharacter("ast-01-game02-003")
	t.Log(test)
}

func TestUnicode(t *testing.T) {
	correct := "C:\\Users\\Yongho1037\\Desktop\\태국TOS업데이트절차.txt"
	t.Log(correct == TrimUnicodeForPath("\u202aC:\\Users\\Yongho1037\\Desktop\\태국TOS업데이트절차.txt"))
	t.Log(correct == TrimUnicodeForPath("\xe2\x80\xaaC:\\Users\\Yongho1037\\Desktop\\태국TOS업데이트절차.txt"))
	t.Log(correct == TrimUnicodeForPath("C:\\Users\\Yongho1037\\Desktop\\태국TOS업데이트절차.txt"))
	t.Log("씨:\\Users\\Yongho1037\\Desktop\\태국TOS업데이트절차.txt" == TrimUnicodeForPath("씨:\\Users\\Yongho1037\\Desktop\\태국TOS업데이트절차.txt"))
}

func TestTrimUnicode(t *testing.T) {
	text := "\u202aC:\\Users\\Yongho1037\\Desktop\\태국TOS업데이트절차.txt"

	s, err := os.Stat(text)
	if err != nil {
		t.Log("unicode exist err : ", err)
	}

	if s == nil {
		t.Log("unicode exist s is nil")
	}

	s, err = os.Stat(TrimUnicodeForPath(text))
	if err != nil {
		t.Log("trim unicode err : ", err)
	}

	if s == nil {
		t.Log("trim unicode s is nil")
	}
}

func TestStringLimit(t *testing.T) {
	targetProcessName := "test-asdfasdfjkasdfjasdfkjashdfkajsdhf"
	processName := "test-asdfasdfjk"
	limitableProcessName := targetProcessName[:len(processName)]
	t.Log(limitableProcessName)
	if !EqualsWithoutCase(limitableProcessName, processName) {
		t.Error("not equal")
	}
}
