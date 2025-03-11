package config

import (
	"github.com/BurntSushi/toml"
)

var OSSfg *OSSConfig

type OSSConfig struct {
	AccessKeyID     string
	AccessKeySecret string
	Endpoint        string
	BucketName      string
}

func init() {
	OSSfg = new(OSSConfig)
	_, err := toml.DecodeFile("config/ossConfig.toml", &OSSfg)
	if err != nil {
		panic(err)
	}
}
