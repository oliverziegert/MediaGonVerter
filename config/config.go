package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"pc-ziegert.de/media_service/constant"
	"pc-ziegert.de/media_service/log"
)

type Config struct {
	// 	log:
	// 		level:
	// 	data:
	//		s3:
	//			access-key-id:
	//    		address:
	//			bucket:
	//			region:
	//			secret-access-key:
	// 	http:
	// 		port:
	// 	rabbitmq:
	// 		addresses:
	// 		username:
	// 		password:
	// 		virtual-host:
	// 	service:
	// 		core:
	// 			protocol:
	// 			host:
	// 			port:
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
		Addresses   string `yaml:"addresses"`
		Username    string `yaml:"username"`
		Password    string `yaml:"password"`
		VirtualHost string `yaml:"virtual-host"`
	} `yaml:"rabbitmq"`

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
