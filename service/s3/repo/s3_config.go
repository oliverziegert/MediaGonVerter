package repo

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"os"
	"pc-ziegert.de/media_service/common/constant"
	e "pc-ziegert.de/media_service/common/error"
	l "pc-ziegert.de/media_service/common/log"
	m "pc-ziegert.de/media_service/common/model"
)

type S3ConfigRepo struct {
	s3Configs map[string]m.S3Config
}

func NewS3Repo() *S3ConfigRepo {
	return &S3ConfigRepo{
		s3Configs: loadS3Configs(),
	}
}

// --- S3 config functions ---

// GetS3Configs retrieves all s3 configs.
func (i *S3ConfigRepo) GetS3Configs() *e.Error {
	return nil
}

// GetS3Config retrieves an s3 by primary key.
func (i *S3ConfigRepo) GetS3Config(ctx *gin.Context) (*m.S3Config, *e.Error) {
	host := ctx.GetString(constant.ContextKeyDracoonForwardedHostHeader)
	s3Config, ok := i.s3Configs[host]
	if !ok {
		return nil, e.NewError(e.ValIdInvalid, "S3 Config not found!")
	}
	return &s3Config, nil
}

func loadS3Configs() map[string]m.S3Config {
	s3Configs := make(map[string]m.S3Config)
	var configFilePath string

	for _, path := range constant.GetS3ConfigFilePaths() {
		if _, err := os.Stat(path); err == nil {
			configFilePath = path
		}
	}

	if configFilePath == "" {
		err := e.NewError(e.ValIdInvalid, "No config file found")
		l.Fatalf(err.StackTrace())
	}

	l.Infof("Using config file: %s", configFilePath)

	readFile(&s3Configs, configFilePath)

	return s3Configs
}

type importS3Config struct {
	m.S3Config `yaml:",inline"`
	Hostnames  []string `yaml:"hostnames"`
}

func readFile(s3Conf *map[string]m.S3Config, filename string) {
	var importS3Configs []importS3Config
	s3ConfMap := *s3Conf
	buf, err := os.ReadFile(filename)
	if err != nil {
		l.Fatalf("Config file missing! %s", err)
	}

	err = yaml.Unmarshal(buf, &importS3Configs)
	if err != nil {
		l.Fatalf("Config file could not be loaded! %s", err)
	}

	for _, iS3Config := range importS3Configs {
		for _, hostname := range iS3Config.Hostnames {
			s3Config := m.S3Config{
				AccessKeyId:     iS3Config.AccessKeyId,
				Address:         iS3Config.Address,
				Bucket:          iS3Config.Bucket,
				Region:          iS3Config.Region,
				SecretAccessKey: iS3Config.SecretAccessKey,
			}
			s3ConfMap[hostname] = s3Config
		}
	}
	return
}
