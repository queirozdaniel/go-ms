package config

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type (
	dbConfig struct {
		Driver   string `mapstructure:"DB_DRIVER"`
		Host     string `mapstructure:"DB_HOST"`
		Port     string `mapstructure:"DB_PORT"`
		User     string `mapstructure:"DB_USER"`
		Password string `mapstructure:"DB_PASSWORD"`
		Database string `mapstructure:"DB_DATABASE"`
	}

	webServiceConfig struct {
		Port string `mapstructure:"REST_PORT"`
	}

	jwtConfig struct {
		Secret       string `mapstructure:"JWT_SECRET"`
		ExpiresIn    int    `mapstructure:"JWT_EXPIRES_IN"`
		TokenAuthKey *jwtauth.JWTAuth
	}

	appConfig struct {
		DB         dbConfig
		WebService webServiceConfig
		Jwt        jwtConfig
	}
)

func InitConfig(path string) (*appConfig, error) {

	var webConfig *webServiceConfig
	var jwtConfig *jwtConfig
	var dbConfig *dbConfig

	viper.SetConfigName("app_authentication")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv() // force to use environment from OS
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&webConfig)

	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&jwtConfig)

	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&dbConfig)

	if err != nil {
		panic(err)
	}

	config := &appConfig{
		DB:         *dbConfig,
		WebService: *webConfig,
		Jwt:        *jwtConfig,
	}

	return config, nil
}
