package utils

import "github.com/spf13/viper"

type Config struct {
	APPPORT  string `mapstructure:"APP_PORT"`
	APPNAME  string `mapstructure:"APP_NAME"`
	APPDEBUG string `mapstructure:"APP_DEBUG"`

	DBCONNECTION string `mapstructure:"DB_CONNECTION"`
	DBHOST       string `mapstructure:"DB_HOST"`
	DBUSERNAME   string `mapstructure:"DB_USERNAME"`
	DBPASSWORD   string `mapstructure:"DB_PASSWORD"`
	DBDATABASE   string `mapstructure:"DB_DATABASE"`
	DBPORT       string `mapstructure:"DB_PORT"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
