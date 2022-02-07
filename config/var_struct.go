package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var (
	ConfigData Config
	Dir        string
)

type Config struct {
	TotalWidth   float32   `yaml:"totalWidth"`
	TotalHeight  float32   `yaml:"totalHeight"`
	TargetWidth  float32   `yaml:"targetWidth"`
	TargetHeight float32   `yaml:"targetHeight"`
	SleepTime    int       `yaml:"sleepTime"`
	Scale        float32   `yaml:"scale"`
	CommendText  string    `yaml:"commendText"`
	DisName      []string  `yaml:"disName"`
	FontSize     []float32 `yaml:"fontSize"`
}

func init() {
	var err error
	Dir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadFile(fmt.Sprintf("%s/config/config.yaml", Dir))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(bytes, &ConfigData)
	if err != nil {
		panic(err)
	}
	TargetDis = (ConfigData.TotalWidth/2 - ConfigData.TargetWidth*3.5) / 3

}
