package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var GlobalConfig Server

func ViperConfig() {
	v := viper.New()
	v.SetConfigFile("./config/config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&GlobalConfig); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&GlobalConfig); err != nil {
		panic(err)
	}
}

type Server struct {
	Window   Window   `mapstructure:"window" json:"window" yaml:"window"`
	Time     Time     `mapstructure:"time" json:"time" yaml:"time"`
	Language Language `mapstructure:"language" json:"language" yaml:"language"`
}
type Window struct {
	ReminderWindowWidth      int  `mapstructure:"reminder-window-width" json:"reminder-window-width" yaml:"reminder-window-width"`
	ReminderWindowHeight     int  `mapstructure:"reminder-window-height" json:"reminder-window-height" yaml:"reminder-window-height"`
	ReminderWindowFullScreen bool `mapstructure:"reminder-window-full-screen" json:"reminder-window-full-screen" yaml:"reminder-window-full-screen"`
	SettingWindowWidth       int  `mapstructure:"setting-window-width" json:"setting-window-width" yaml:"setting-window-width"`
	SettingWindowHeight      int  `mapstructure:"setting-window-height" json:"setting-window-height" yaml:"setting-window-height"`
}
type Time struct {
	RestTime     int `mapstructure:"rest-time" json:"rest-time" yaml:"rest-time"`
	RestDuration int `mapstructure:"rest-duration" json:"rest-duration" yaml:"rest-duration"`
}
type Language struct {
	Language string `mapstructure:"language" json:"language" yaml:"language"`
}
