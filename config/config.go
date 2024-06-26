package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	GHTMX struct {
		Version string `mapstructure:"version"`

		Port int `mapstructure:"port"`
	} `mapstructure:"ghtmx"`

	Storage struct {
		Username string `mapstructure:"username"`

		Password string `mapstructure:"password"`

		Name string `mapstructure:"name"`

		Addr string `mapstructure:"addr"`
	} `mapstructure:"storage"`
}

func LoadConfig() Config {
	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if _, ok := os.LookupEnv("STORAGE_USERNAME"); ok {
		err := viper.BindEnv("storage.username", "STORAGE_USERNAME")
		if err != nil {
			return Config{}
		}
	}

	if _, ok := os.LookupEnv("STORAGE_PASSWORD"); ok {
		err := viper.BindEnv("storage.password", "STORAGE_PASSWORD")
		if err != nil {
			return Config{}
		}
	}

	if _, ok := os.LookupEnv("STORAGE_ADDR"); ok {
		err := viper.BindEnv("storage.addr", "STORAGE_ADDR")
		if err != nil {
			return Config{}
		}
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
