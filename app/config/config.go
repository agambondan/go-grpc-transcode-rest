package config

import (
	"os"
)

var Config Configuration

type Configuration struct {
	DBDriver      string
	DBHost        string
	DBName        string
	DBUser        string
	DBPassword    string
	DBPort        string
	Port          string
	RedisHost     string
	RedisPort     string
	RedisPassword string
	AccessSecret  string
	RefreshSecret string
	CookiesSecret string
	Salt          string
	CipherKey     string
	Environment   string
}

func (config *Configuration) Init() {
	Config.DBDriver = os.Getenv("DB_DRIVER")
	Config.DBHost = os.Getenv("DB_HOST")
	Config.DBPassword = os.Getenv("DB_PASS")
	Config.DBUser = os.Getenv("DB_USER")
	Config.DBName = os.Getenv("DB_NAME")
	Config.DBPort = os.Getenv("DB_PORT")
	Config.Port = os.Getenv("PORT")
	Config.RedisHost = os.Getenv("REDIS_HOST")
	Config.RedisPort = os.Getenv("REDIS_PORT")
	Config.RedisPassword = os.Getenv("REDIS_PASSWORD")
	Config.AccessSecret = os.Getenv("ACCESS_SECRET")
	Config.RefreshSecret = os.Getenv("REFRESH_SECRET")
	Config.CookiesSecret = os.Getenv("COOKIES_SECRET")
	Config.Salt = os.Getenv("SALT")
	Config.CipherKey = os.Getenv("CIPHER_KEY")
	Config.Environment = os.Getenv("ENVIRONMENT")
}
