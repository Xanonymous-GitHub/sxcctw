package vp

import (
	"github.com/spf13/viper"
	"log"
)

const (
	defaultConfigFilePath = "config/config.yml"
	defaultConfigFileType = "yml"
)

var Cvp *viper.Viper

func init() {
	Cvp = InitViper()
}

// TODO(TU) us auto env.

func InitViper() *viper.Viper {
	v := viper.New()

	v.SetConfigFile(defaultConfigFilePath)
	v.SetConfigType(defaultConfigFileType)

	err := v.ReadInConfig()
	if err != nil {
		log.Println(err)
	}

	v.WatchConfig()

	return v
}
