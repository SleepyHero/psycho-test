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
	TotalWidth      float32        `yaml:"totalWidth"`
	TotalHeight     float32        `yaml:"totalHeight"`
	TargetWidth     float32        `yaml:"targetWidth"`
	TargetHeight    float32        `yaml:"targetHeight"`
	SleepTime       int            `yaml:"sleepTime"`
	Scale           float32        `yaml:"scale"`
	CommendText     string         `yaml:"commendText"`
	DisName         []string       `yaml:"disName"`
	FontSize        []float32      `yaml:"fontSize"`
	TextExamTitle   string         `yaml:"textExamTitle"`
	TextExamContent string         `yaml:"textExamContent"`
	NumExamTitle    string         `yaml:"numExamTitle"`
	NumExamContent  string         `yaml:"numExamContent"`
	KeyMap          map[string]int `yaml:"keyMap"`
	TextSize        float32        `yaml:"textSize"`
	WindowHeight    float32        `yaml:"windowHeight"`
	RepeatTimes     int            `yaml:"repeatTimes"`
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
	ConfigData.TotalHeight *= ConfigData.Scale
	ConfigData.TotalWidth *= ConfigData.Scale
	ConfigData.TargetWidth *= ConfigData.Scale
	ConfigData.TargetHeight *= ConfigData.Scale
	ConfigData.WindowHeight *= ConfigData.Scale
	TargetDis = (ConfigData.TotalWidth/2 - ConfigData.TargetWidth*3.5) / 3
	BaseHeight = (ConfigData.WindowHeight - ConfigData.TotalHeight) / 2
}
