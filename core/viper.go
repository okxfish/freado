package core

import (
	"fmt"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("./config/local.yml")
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	//v.WatchConfig()
	//v.OnConfigChange(func(e fsnotify.Event) {}

	return v
}
