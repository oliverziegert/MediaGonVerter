package model

type S3Config struct {
	AccessKeyId     string `yaml:"access-key-id"`
	Address         string `yaml:"address"`
	Bucket          string `yaml:"bucket"`
	Region          string `yaml:"region"`
	SecretAccessKey string `yaml:"secret-access-key"`
}
