package setting

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
)

//Config definition struct
type Config struct {
	LogFile        string //`toml:"somename"`
	InfoLogFile    string
	WarningLogFile string
	ErrorLogFile   string
	GinMode        string
	HTTPPort       int
	DbName         string
	DbUser         string
	DbPassword     string
	DbServer       string
	DbPort         string
	JwtSecret      string
	AccessToken    string
	ThreadId       string
	Host           string
	ClientID       string
	ClientSecret   string
}

//AppSetting common setting aplication
var AppSetting = &Config{}

var ginModes = map[string]string{
	"release": "release",
	"debug":   "debug",
}

//Setup initial setting
func Setup() {
	envFile := os.Getenv("ENV_FILE")

	if len(envFile) == 0 {
		envFile = ".dev.env"
	}

	tomlData, err := ioutil.ReadFile(envFile)
	if err != nil {
		log.Fatalf("Can´t find config file")
	}
	if _, err := toml.Decode(string(tomlData), &AppSetting); err != nil {
		// handle error
		log.Fatalf("Can´t read config file")
	}
}
