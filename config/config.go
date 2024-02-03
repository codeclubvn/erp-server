package config

import (
	"erp/utils/constants"
	"fmt"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	ConfigDefaultFile = "config/config.yml"
	ConfigReleaseFile = "config/config.release.yml"
	ConfigDevFile     = "config/config.dev.yml"
	configType        = "yml"
)

var (
	configEnv     = "./config/app.env"
	configTypeEnv = "env"
	configEnvName = "app"
)

type (
	Config struct {
		Env            Env        `mapstructure:"env"`
		Debug          bool       `mapstructure:"debug"`
		ContextTimeout int        `mapstructure:"contextTimeout"`
		Server         Server     `mapstructure:"server"`
		Services       Services   `mapstructure:"services"`
		Database       Database   `mapstructure:"database"`
		Logger         Logger     `mapstructure:"logger"`
		Jwt            Jwt        `mapstructure:"jwt"`
		Cloudinary     Cloudinary `mapstructure:"cloudinary"`
	}

	Server struct {
		Host       string `mapstructure:"host"`
		Env        string `mapstructure:"env"`
		UseRedis   bool   `mapstructure:"useRedis"`
		Port       int    `mapstructure:"port"`
		UploadPath string `mapstructure:"uploadPath"`
	}

	Database struct {
		Driver   string `mapstructure:"driver"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
		SSLMode  string `mapstructure:"sslmode"`
		TimeZone string `mapstructure:"timeZone"`
	}

	Env struct {
		Env           string `mapstructure:"ENV"`
		CloudinaryURL string `mapstructure:"CLOUDINARY_URL"`
	}

	Jwt struct {
		Secret                string `mapstructure:"secret"`
		AccessTokenExpiresIn  int64  `mapstructure:"accessTokenExpiresIn"`
		RefreshTokenExpiresIn int64  `mapstructure:"refreshTokenExpiresIn"`
	}

	Logger struct {
		Level  string `mapstructure:"level"`
		Format string `mapstructure:"format"`
		Prefix string `mapstructure:"prefix"`
	}

	Services struct {
	}

	Cloudinary struct {
		CloudName string `mapstructure:"cloudName"`
		ApiKey    string `mapstructure:"apiKey"`
		ApiSecret string `mapstructure:"apiSecret"`
		PublicId  string `mapstructure:"publicId"`
		URL       string `mapstructure:"url"`
	}
)

func initEnv(conf *Config) {
	if err := LoadConfigEnv(configEnv, configTypeEnv); err != nil {
		fmt.Printf("unable decode into config struct, %v", err)
	}
	if err := UnmarsharConfig(&conf.Env); err != nil {
		fmt.Printf("unable decode into config struct, %v", err)
	}
	SetEnv(conf)
}

func NewConfig() *Config {
	conf := &Config{}
	initEnv(conf)
	initConfig()
	if err := UnmarsharConfig(conf); err != nil {
		fmt.Printf("unable decode into config struct, %v", err)
	}
	return conf
}

func LoadConfigEnv(configFile, configType string) (err error) {
	viper.SetConfigType(configType)
	viper.SetConfigFile(configFile)

	if err = viper.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
	}
	return
}

func UnmarsharConfig[E any](config *E) error {
	return viper.Unmarshal(config)

}

func SetEnv(config *Config) {
	v := reflect.ValueOf(config.Env)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() != "" {
			os.Setenv(v.Type().Field(i).Tag.Get("mapstructure"), v.Field(i).Interface().(string))
		}
	}
}

func initConfig() {
	var configFile string
	switch os.Getenv("ENV") {
	case constants.Prod:
		configFile = ConfigReleaseFile
		fmt.Printf("gin mode %s,config%s\n", gin.EnvGinMode, ConfigReleaseFile)
	case constants.Dev:
		configFile = ConfigDevFile
		fmt.Printf("gin mode: %s,config file: %s\n", gin.EnvGinMode, ConfigDevFile)
	default:
		configFile = ConfigDefaultFile
		fmt.Printf("gin mode %s,config%s\n", gin.EnvGinMode, ConfigDefaultFile)
	}
	viper.SetConfigType(configType)
	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err.Error())
	}
}
