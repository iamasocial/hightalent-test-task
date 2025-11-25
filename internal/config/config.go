package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config represents application configuration, including HTTP and database settings.
type Config struct {
	HTTP `yaml:"http"`
	DB   `yaml:"db"`
}

// HTTP holds configuration for HTTP server, such as port and timeouts.
type HTTP struct {
	Port            string        `yaml:"port"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
	IdleTimeout     time.Duration `yaml:"idle_timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}

// DB holds database connection settings, including credentials and connection pool parameters.
type DB struct {
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	User            string
	Password        string
	Name            string
	SSLMode         string        `yaml:"sslmode"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
}

// LoadConfig loads application configuration from YAML file and environment variables.
// It first load .env, reads the YAML at configPath into Config.
func LoadConfig(configPath string) (*Config, error) {
	// if err := godotenv.Load(); err != nil {
	// 	return nil, fmt.Errorf("failed to load .env file: %s", err)
	// }

	if configPath == "" {
		return nil, fmt.Errorf("config file path is empty")
	}

	if _, err := os.Stat(configPath); err != nil {
		return nil, fmt.Errorf("config file not found: %s", err)
	}

	cfg := Config{}

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	cfg.DB.User = os.Getenv("DB_USER")
	cfg.DB.Password = os.Getenv("DB_PASSWORD")
	cfg.DB.Name = os.Getenv("DB_NAME")

	return &cfg, nil
}
