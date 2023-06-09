package initialize

import (
	"fmt"

	"github.com/retail-ai-test/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() error {
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("./%s-dev.yaml", configFilePrefix)

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panic("Failed to read config file: %v", err)
		return err
	}

	if err := v.Unmarshal(&config.ServerConf); err != nil {
		zap.S().Panic("Failed to unmarshal config file: %v", err)
		return err
	}
	return nil
}
