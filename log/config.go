package log

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"reflect"
	"fmt"
)

const confPath = "./configuration.yaml"

type ConfigArray struct {
	LogPath string `yaml:"log_path"`
	LocalAddress string `yaml:"local_address"`
}

var GlobalConfig ConfigArray

func init() {
	conf, err := ioutil.ReadFile(confPath)
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(conf, &GlobalConfig)
	fmt.Println("Initializing...")
}

func GetConfig (k string) string {
	config_ins := reflect.ValueOf(GlobalConfig)
	return config_ins.FieldByName(k).String()
}

func SetConfig (k string, v string) {
	config_ins := reflect.ValueOf(GlobalConfig)
	config_ins.FieldByName(k).SetString(v)
}

