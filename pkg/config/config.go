package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env    string `yaml:"env"`
	Server server `yaml:"server"`
	DB     db     `yaml:"db"`
}

type server struct {
	Port        string        `yaml:"port"`
	Host        string        `yaml:"host"`
	RWTimeout   time.Duration `yaml:"rw_timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type db struct {
	Port   string `yaml:"port"`
	DBName string `yaml:"dbname"`
	User   string `yaml:"user"`
}

func MustLoad() *Config {
	path := os.Getenv("PATH_TO_CONFIG_FILE")
	if path == "" {
		log.Fatal("path to config is empty")
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		log.Fatalf("err with unmarshalling config file: %w", err)
	}

	return &cfg
}
