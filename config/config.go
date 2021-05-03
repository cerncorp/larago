package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// mixing file and env
type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`

	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"database"`
}

var Cfg Config

func init() {
	ReadFile(&Cfg)
	// readEnv(&Cfg)
	fmt.Printf("%+v", Cfg)
	fmt.Println("config package initialized")
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func ReadFile(cfg *Config) {
	f, err := os.Open("env.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

// github.com/kelseyhightower/envconfig
// func ReadEnv(cfg *Config) {
// 	err := envconfig.Process("", cfg)
// 	if err != nil {
// 		processError(err)
// 	}
// }

// func main() {
//     var cfg Config
//     readFile(&cfg)
//     readEnv(&cfg)
//     fmt.Printf("%+v", cfg)
// }
