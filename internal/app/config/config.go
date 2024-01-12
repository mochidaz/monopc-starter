package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host         string `env:"APP_HOST,default=localhost"`
	Env          string `env:"APP_ENV,default=development"`
	Port         string `env:"APP_PORT,default=3000"`
	AppName      string `env:"APP_NAME,default=mono-starter"`
	Database     Database
	LocalStorage LocalStorage
	JWTConfig    JWTConfig
	//Redis    Redis
}

type Database struct {
	Host            string `env:"DB_HOST,default=localhost"`
	Port            string `env:"DB_PORT,default=5432"`
	User            string
	Password        string
	Name            string
	MaxOpenConns    string `env:"DB_MAX_OPEN_CONNS,default=10"`
	MaxConnLifetime string `env:"DB_MAX_CONN_LIFETIME,default=5m"`
	MaxIdleLifetime string `env:"DB_MAX_IDLE_LIFETIME,default=5m"`
}

type LocalStorage struct {
	BasePath string `env:"LOCAL_STORAGE_BASE_PATH,default=/tmp"`
}

type JWTConfig struct {
	Public    string `env:"JWT_PUBLIC,required"`
	Private   string `env:"JWT_PRIVATE,required"`
	Issuer    string `env:"JWT_ISSUER,required"`
	IssuerCMS string `env:"JWT_ISSUER_CMS,required"`
}

//type Redis struct {
//	Host     string
//	Port     string
//	Password string
//}

// LoadConfig memuat konfigurasi dari file .env dan mengembalikan struktur Config
func LoadConfig(path string) Config {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		Host:    getEnv("APP_HOST", "localhost"),
		Env:     getEnv("APP_ENV", "development"),
		Port:    getEnv("APP_PORT", "8080"),
		AppName: getEnv("APP_NAME", "mono-starter"),
		Database: Database{
			Host:            getEnv("DB_HOST", "localhost"),
			Port:            getEnv("DB_PORT", "5432"),
			User:            getEnv("DB_USER", "postgres"),
			Password:        getEnv("DB_PASSWORD", "postgres"),
			Name:            getEnv("DB_NAME", "postgres"),
			MaxOpenConns:    getEnv("DB_MAX_OPEN_CONNS", "10"),
			MaxConnLifetime: getEnv("DB_MAX_CONN_LIFETIME", "5m"),
			MaxIdleLifetime: getEnv("DB_MAX_IDLE_LIFETIME", "5m"),
		},
		LocalStorage: LocalStorage{
			BasePath: getEnv("LOCAL_STORAGE_BASE_PATH", "/tmp"),
		},
		JWTConfig: JWTConfig{
			Public:    getEnv("JWT_PUBLIC", ""),
			Private:   getEnv("JWT_PRIVATE", ""),
			Issuer:    getEnv("JWT_ISSUER", ""),
			IssuerCMS: getEnv("JWT_ISSUER_CMS", ""),
		},
	}
}

// getEnv membantu mendapatkan variabel lingkungan dengan default value
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
