package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Mongo    mongo
	External external
	Server   server
}

type mongo struct {
	URI      string `envconfig:"MONGO_URI" default:"mongodb://localhost:27017"`
	Database string `envconfig:"DB_NAME" default:"covid"`
}

type server struct {
	Port int `envconfig:"PORT" default:"8080"`
}

type external struct {
	Covid string `envconfig:"COVID_URL"`
}

var cfg Config

func Init() {
	_ = godotenv.Load()
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("read env error : %s", err.Error())
	}
}

func Get() Config {
	return cfg
}
