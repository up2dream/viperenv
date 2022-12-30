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
