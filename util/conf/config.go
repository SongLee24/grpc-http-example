package conf

import "github.com/spf13/viper"

const (
	ConfigFilePath = "./config"
)

func init() {
	viper.AddConfigPath(ConfigFilePath)
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Read config.yaml failed: " + err.Error())
	}
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetInt64(key string) int64 {
	return viper.GetInt64(key)
}
