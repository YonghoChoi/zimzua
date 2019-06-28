package util

import (
	"errors"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type SortString []string

func (s SortString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SortString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortString) Len() int {
	return len(s)
}

func DeleteElementOfStringArray(slice []string, target string) ([]string, error) {
	size := len(slice)
	for i := 0; i < size; i++ {
		if slice[i] == target {
			return append(slice[:i], slice[i+1:]...), nil
		}
	}

	return nil, errors.New("not exist target element in slice")
}

// TransformToKorean : 한글 출력의 경우 깨짐 현상을 방지하기 위해 EUCKR 인코딩
func TransformToKorean(target string) string {
	utf8 := transform.NewReader(strings.NewReader(target), korean.EUCKR.NewDecoder())
	decBytes, err := ioutil.ReadAll(utf8)
	if err != nil { // 한글 인코딩에 실패하는 경우 원문 반환
		return target
	}

	return string(decBytes)
}

func AddQuotation(origin string) string {
	return "'" + RemoveQuotation(origin) + "'"
}

func AddDoubleQuotation(origin string) string {
	return "\"" + RemoveQuotation(origin) + "\""
}

func RemoveQuotation(origin string) string {
	origin = strings.Replace(origin, "'", "", -1)
	origin = strings.Replace(origin, "\"", "", -1)
	return origin
}

func RemoveBracket(origin string) string {
	origin = strings.Replace(origin, "(", "", -1)
	origin = strings.Replace(origin, ")", "", -1)
	return origin
}

func RemoveSlash(origin string) string {
	origin = strings.Replace(origin, "/", "", -1)
	return origin
}

// atoiWithoutError : 변환에 실패하더라도 원활한 동작을 위해 허용
func AtoiWithoutError(origin string) int {
	num, err := strconv.Atoi(origin)
	if err != nil {
		log.Printf("can not convert string to integer. value : %s\n", origin)
	}

	return num
}

func TrimSpecialCharacter(str string) string {
	var result string
	var validID = regexp.MustCompile(`^[a-zA-Z0-9]$`)
	for _, s := range []byte(str) {
		if validID.MatchString(string(s)) {
			result = result + string(s)
		}
	}

	return result
}

func EqualStringWithoutCase(src string, target string) bool {
	return strings.ToLower(src) == strings.ToLower(target)
}

func Index(arr []string, target string) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func Include(arr []string, target string) bool {
	return Index(arr, target) >= 0
}

// array의 순서가 중요하지 않은 경우 지울 위치에 마지막 값을 대입하는 것으로 대상 제거
func RemoveFast(arr []string, target string) []string {
	index := Index(arr, target)
	arr[index] = arr[len(arr)-1]
	arr[len(arr)-1] = ""
	return arr[:len(arr)-1]
}

// array 순서를 유지해야 하는 경우 새로운 배열 생성 (복사 비용 발생)
func RemoveSlow(arr []string, target string) []string {
	index := Index(arr, target)
	copy(arr[index:], arr[index+1:])
	arr[len(arr)-1] = ""
	return arr[:len(arr)-1]
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func RemoveWhiteSpace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func EqualsWithoutCase(src string, target string) bool {
	return strings.ToLower(src) == strings.ToLower(target)
}

func NotEquals(src string, target string) bool {
	return strings.ToLower(src) != strings.ToLower(target)
}
