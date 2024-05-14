package utils

import "github.com/spf13/viper"

// Config stores all configuration of the application
// The values are read by viper from a config file or environment variables
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path) //tells viper the path to look for a config 
	viper.SetConfigName("app") //tells viper the name of the config file
	viper.SetConfigType("env")

	viper.AutomaticEnv() //reads matching envs

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return 
}
