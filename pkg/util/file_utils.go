package util

import (
	"bytes"
	"compress/gzip"
	"github.com/shirou/gopsutil/host"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func ExtractFileNameWithExt(path string) string {
	return filepath.Base(path) + "." + filepath.Ext(path)
}

func IsValidPath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}

	return nil
}

func MkDir(path string) error {
	if err := IsValidPath(path); err != nil {
		if os.IsNotExist(err) {
			if err = os.Mkdir(path, 0777); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func DownloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func GUnzipData(data []byte) (resData []byte, err error) {
	b := bytes.NewBuffer(data)

	var r io.Reader
	r, err = gzip.NewReader(b)
	if err != nil {
		return
	}

	var resB bytes.Buffer
	_, err = resB.ReadFrom(r)
	if err != nil {
		return
	}

	resData = resB.Bytes()

	return
}

func HomeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func getOS() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", err
	}

	return info.OS, nil
}

func HomeDirByAccount(account string) string {
	osType, _ := getOS()
	if osType == "windows" {
		return "C:\\Users\\" + account
	} else {
		return "/home/" + account
	}
}

func HomeBase() string {
	osType, _ := getOS()
	if osType == "windows" {
		return "C:\\Users"
	} else {
		return "/home"
	}
}

func HomeUser() string {
	home := HomeDir()
	osType, _ := getOS()
	var seperator string
	if osType == "windows" {
		seperator = "\\"
	} else {
		seperator = "/"
	}

	pathSplit := strings.Split(home, seperator)
	if len(pathSplit) > 0 {
		return pathSplit[len(pathSplit)-1]
	}

	return ""
}

func GetFileContent(path string, size int) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	buf := make([]byte, size)
	_, err = file.Read(buf)
	if err != nil && err != io.EOF {
		return "", err
	}

	return string(buf), nil
}

func TrimUnicodeForPath(text string) string {
	// WEB Form에서 복사&붙여넣기 할 경우 앞에 유니코드 문자(EN Space)가 포함되는 문제가 있음
	// 앞 부분의 유니코드 제거
	result := ""
	bytes := []byte(text)
	re := regexp.MustCompile("[[:ascii:]]") // Windows/Linux 모두 경로 앞은 ascii로 시작. 맨 앞이 유니코드인 경우 모두 제거
	for index, byte := range bytes {
		if re.MatchString(string(byte)) {
			result = string(bytes[index:])
			break
		}
	}

	return result
}
