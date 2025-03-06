package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort   string
	KeycloackURL string
	ClientID     string
	ClientSecret string
	Realm        string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("xz while")
	}

	return &Config{
		ServerPort:   getEnv("SERVER_PORT", "8081"),
		KeycloackURL: getEnv("KEYCLOACK_URL", "http://192.168.0.102:8080"),
		ClientID:     getEnv("CLIENT_ID", "polyclient"),
		ClientSecret: getEnv("CLIENT_SECRET", "jtfzpXIbb0YVGZqeJQb290zxTQ1QdxWt"),
		Realm:        getEnv("REALM", "master"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
