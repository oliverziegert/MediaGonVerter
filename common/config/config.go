package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"pc-ziegert.de/media_service/common/constant"
	"pc-ziegert.de/media_service/common/log"
)

type Config struct {
	Auth struct {
		Jwt struct {
			PublicKey string `yaml:"public-key"`
		} `yaml:"jwt"`
	} `yaml:"auth"`

	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`

	Data struct {
		S3 struct {
			AccessKeyId     string `yaml:"access-key-id"`
			Address         string `yaml:"address"`
			Bucket          string `yaml:"bucket"`
			Region          string `yaml:"region"`
			SecretAccessKey string `yaml:"secret-access-key"`
		} `yaml:"s3"`
		Expiration string `yaml:"expiration"`
	} `yaml:"data"`

	HTTP struct {
		Gzip struct {
			Enabled bool `yaml:"enabled"`
		} `yaml:"gzip"`

		Origin struct {
			All struct {
				Enabled bool `yaml:"enabled"`
			} `yaml:"all"`
		} `yaml:"origin"`

		Port uint16 `yaml:"port"`
	} `yaml:"http"`

	RabbitMQ struct {
		IP          string `yaml:"ip"`
		Port        uint16 `yaml:"port"`
		Username    string `yaml:"username"`
		Password    string `yaml:"password"`
		VirtualHost string `yaml:"virtual-host"`
	} `yaml:"rabbitmq"`

	Redis struct {
		Network  string `yaml:"network"`
		IP       string `yaml:"ip"`
		Port     uint16 `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DB       uint8  `yaml:"db"`
	} `yaml:"redis"`

	Service struct {
		Core struct {
			Protocol string `yaml:"protocol"`
			Host     string `yaml:"host"`
			Port     uint16 `yaml:"port"`
		} `yaml:"core"`
		Internal struct {
			Communication struct {
				Auth struct {
					Token string `yaml:"token"`
				} `yaml:"auth"`
			} `yaml:"communication"`
		} `yaml:"internal"`
	} `yaml:"service"`
}

func LoadConfig() *Config {
	var conf Config
	readFile(&conf, constant.ConfigFileName)
	return &conf
}

func readFile(conf *Config, filename string) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Config file missing! %s", err)
	}

	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		log.Fatalf("Config file could not be loaded! %s", err)
	}
}
