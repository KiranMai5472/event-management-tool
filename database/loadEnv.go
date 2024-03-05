package database

import (
	"github.com/spf13/viper"
)

// Config used to get the database values and details from env file
type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"PORT"`
	JwtKey         string `mapstructure:"JWT_SECRET_KEY"`
	ClientOrigin   string `mapstructure:"CLIENT_ORIGIN"`
}

// Cache use to get the caching value from the evv file
type Cache struct {
	CacheStatus        string `mapstructure:"CACHE_STATUS"`
	GetLanguages       int    `mapstructure:"GETLANGUAGES"`
	GetLocalLanguages  int    `mapstructure:"GETLOCALLANGUAGES"`
	GetDevices         int    `mapstructure:"GETDEVICES"`
	GetCountry         int    `mapstructure:"GETCOUNTRY"`
	GetContentType     int    `mapstructure:"GETCONTENTTYPE"`
	GetPreviewMimeType int    `mapstructure:"GETPREVIEWMIMETYPE"`
	GetParentalRating  int    `mapstructure:"GETPARENTALRATING"`
	GetStreamingType   int    `mapstructure:"GETSTREAMINGTYPE"`
	GetCastRole        int    `mapstructure:"GETCASTROLE"`
	GetBusiness        int    `mapstructure:"GETBUSINESS"`
}

// LoadConfig function used to read the values of database from env file
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

// LoadCache function used to read the values of caching from env file
func LoadCache(path string) (cache Cache, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cache)
	return
}
