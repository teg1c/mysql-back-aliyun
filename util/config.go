package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

// Config Mirror Config
type Config struct {
	Database           string `yaml:"DATABASE"`
	MysqlHost          string `yaml:"MYSQL_HOST"`
	MysqlUsername      string `yaml:"MYSQL_USERNAME"`
	MysqlPassword      string `yaml:"MYSQL_PASSWORD"`
	BackDir            string `yaml:"BACK_DIR"`
	ContainerName      string `yaml:"CONTAINER__NAME"`
	OSSAccessKeyID     string `yaml:"OSS_ACCESS_KEY_ID"`
	OSSAccessKeySecret string `yaml:"OSS_ACCESS_KEY_SECRET"`
	OSSEndpoint        string `yaml:"OSS_ENDPOINT"`
	OSSBucket          string `yaml:"OSS_BUCKET"`
	FullPath           string
}

func LoadConfig(configPath string) (conf *Config, err error) {
	content, err := getYamlContent(configPath)
	if err != nil {
		return
	}
	conf = new(Config)
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		return
	}
	err = conf.ValidateConfig()
	return
}

func getYamlContent(yamlPath string) (content []byte, err error) {
	ymlFile, err := os.Open(yamlPath)
	if err != nil {
		return
	}
	defer ymlFile.Close()
	content, err = ioutil.ReadAll(ymlFile)
	return
}

func (config *Config) ValidateConfig() (err error) {
	if config.Database == "" {
		err = errors.New("missing configuration: DATABASE")
		return
	}

	if config.MysqlHost == "" {
		err = errors.New("missing configuration: MYSQL_HOST")
		return
	}

	if config.MysqlUsername == "" {
		err = errors.New("missing configuration: MYSQL_USERNAME")
		return
	}

	if config.MysqlPassword == "" {
		err = errors.New("missing configuration: MYSQL_PASSWORD")
		return
	}

	if config.OSSEndpoint == "" {
		err = errors.New("missing configuration: OSS_ENDPOINT")
		return
	}

	if config.OSSBucket == "" {
		err = errors.New("missing configuration: OSS_BUCKET")
		return
	}

	backUpFileName := fmt.Sprintf("%s-%s.sql", config.Database, time.Now().Format("20060102"))
	config.FullPath = fmt.Sprintf("%s%s", config.BackDir, backUpFileName)
	return
}
