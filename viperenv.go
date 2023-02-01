// Package viperenv
// 配置文件必须放在config目录下，文件名为config.yaml
// 特定环境的配置文件为config-<profile>.yaml，<profile>为环境变量app.profiles.active的值，如果没有设置环境变量，
// 则使用config.yaml中的app.profiles.active的值，如果没有设置app.profiles.active的值，则使用dev作为<profile>的值。
// 如果有多个附加配置文件，则可以在config.yaml中设置app.profiles.include的值，值可以是一个字符串，也可以是一个字符串数组。
// 例如：
// app:
//
//	profiles:
//	  active: dev
//	  include: [dev, test]
package viperenv

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

const (
	keyProfilesActive  = "app.profiles.active"
	keyProfilesInclude = "app.profiles.include"
)

var Config = readConfigAll()

func readConfigAll() *viper.Viper {
	config := readConfig("")
	config.SetDefault(keyProfilesActive, "dev")
	includeProfiles := config.Get(keyProfilesInclude)
	if includeProfiles != nil {
		if _, ok := includeProfiles.(string); ok {
			configDev := readConfig(includeProfiles.(string))
			config.MergeConfigMap(configDev.AllSettings())
		} else if _, ok := includeProfiles.([]interface{}); ok {
			for _, includeProfile := range includeProfiles.([]interface{}) {
				configInclude := readConfig(includeProfile.(string))
				config.MergeConfigMap(configInclude.AllSettings())
			}
		}
	}
	var activeProfile interface{} = os.Getenv(keyProfilesActive)
	if len(strings.TrimSpace(activeProfile.(string))) == 0 {
		activeProfile = config.Get(keyProfilesActive)
	}
	fmt.Printf("Current config: %s\n", activeProfile)
	if activeProfile != nil {
		configDev := readConfig(activeProfile.(string))
		config.MergeConfigMap(configDev.AllSettings())
	}

	return config
}
func readConfig(env string) *viper.Viper {
	config := viper.New()
	configName := "config"
	if len(strings.TrimSpace(env)) != 0 {
		configName += "-" + env
	}
	config.SetConfigName(configName)
	configType := "yaml"
	config.SetConfigType(configType)
	config.AddConfigPath("./config")
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("Config file(%s.%s) not found, ignore...\n", configName, configType)
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}
	return config
}
