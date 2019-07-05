package config

import (
    "github.com/joeshaw/envdecode"
    "sync"
    "errors"
    "os"
    "path"
    "encoding/json"
    "io/ioutil"
    "gopkg.in/yaml.v2"
    "fmt"
)

type MongoConfig struct {
    Host     string `json:"host" yaml:"host"`
    Port     int    `json:"port" yaml:"port"`
    Database string `json:"database" yaml:"database"`
    Username string `json:"username" yaml:"username"`
    Password string `json:"username" yaml:"password"`
}

type ZimzuaConfig struct {
    Mongo       MongoConfig              `yaml:"mongo"`
}

var curConfig *ZimzuaConfig
var once sync.Once

// GetInstance 환경설정 싱글턴 객체 반환
func GetInstance(configFile ...string) *ZimzuaConfig {
    once.Do(func() {
        curConfig = new(ZimzuaConfig)
        if len(configFile) == 0 {
            panic(errors.New("config file is invalid"))
        }
        if err := LoadConfig(configFile[0], curConfig); err != nil {
            panic(errors.New("fail load config"))
        }

        fmt.Println(fmt.Sprintf("success load config. %v\n", curConfig))
    })

    return curConfig
}

// LoadConfig 설정파일에서 설정을 불러온다.
// NOTE: filepath가 유효하지 않을 경우 환경변수에서 설정을 가져온다
func LoadConfig(filepath string, o interface{}) error {
    if _, err := os.Stat(filepath); os.IsNotExist(err) {
        return envdecode.Decode(o)
    }

    if filepath == "" || o == nil {
        return nil
    }

    ext := path.Ext(filepath)

    if ext == ".json" {
        return loadJSONConfig(filepath, o)
    }

    if ext == ".yml" || ext == ".yaml" {
        return loadYMLConfig(filepath, o)
    }

    return fmt.Errorf("not exist config")
}

func loadJSONConfig(filepath string, o interface{}) error {
    file, err := os.Open(filepath)
    if err != nil {
        return err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)

    err = decoder.Decode(&o)
    if err != nil {
        return err
    }

    return nil
}

func loadYMLConfig(filepath string, o interface{}) error {
    file, err := os.Open(filepath)
    if err != nil {
        return err
    }
    defer file.Close()

    b, err := ioutil.ReadAll(file)

    return yaml.Unmarshal(b, o)
}
