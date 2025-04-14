package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server `yaml:"server"`
	DBConn `yaml:"db-conn"`
}

type Server struct {
	JWTKey       string        `env:"JWTKEY" env-required:"true"`
	Host         string        `yaml:"host" env:"HOST" env-default:"localhost"`
	Port         string        `yaml:"port" env:"PORT" env-default:"8080"`
	ResponseTime time.Duration `yaml:"response_time" env-default:"100ms"`
	RPS          int           `yaml:"rps" env-default:"1000"`
	LogLevel     string        `yaml:"log_level" env-default:"info"`
}

type DBConn struct {
	URL          string `yaml:"url" env:"DB_URL" env-required:"true"`
	MaxOpenConns int    `yaml:"max_open_conns" env-default:"15"`
}

func New(path string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
