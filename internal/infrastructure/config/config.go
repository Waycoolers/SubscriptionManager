package config

import (
	"errors"
	"log/slog"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Server *ServerConfig
	DB     *DBConfig
	LogLVL string
}

type ServerConfig struct {
	Host        string
	Port        int
	HTTPTimeout time.Duration
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func Load() (*Config, error) {
	lvl := os.Getenv("LOG_LVL")
	if lvl == "" {
		slog.Warn("Not found LOG_LVL")
	}

	server, err := loadServerConfig()
	if err != nil {
		return nil, err
	}
	db, err := loadDBConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		Server: server,
		DB:     db,
		LogLVL: lvl,
	}, nil
}

func loadServerConfig() (*ServerConfig, error) {
	host := os.Getenv("HTTP_HOST")
	if host == "" {
		host = "localhost"
		slog.Warn("Not found HTTP_HOST. Using default", "host", host)
	}
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
		slog.Warn("Not found HTTP_PORT. Using default", "port", port)
	}
	httpTimeout := os.Getenv("HTTP_TIMEOUT")
	if httpTimeout == "" {
		httpTimeout = "30"
		slog.Warn("Not found HTTP_TIMEOUT. Using default", "timeout", httpTimeout)
	}

	intPort, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}
	intHTTPTimeout, err := strconv.Atoi(httpTimeout)
	if err != nil {
		return nil, err
	}

	return &ServerConfig{
		Host:        host,
		Port:        intPort,
		HTTPTimeout: time.Duration(intHTTPTimeout) * time.Second,
	}, nil
}

func loadDBConfig() (*DBConfig, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		return nil, errors.New("DB_HOST environment variable not set")
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		return nil, errors.New("DB_PORT environment variable not set")
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		return nil, errors.New("DB_USER environment variable not set")
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return nil, errors.New("DB_PASSWORD environment variable not set")
	}
	name := os.Getenv("DB_NAME")
	if name == "" {
		return nil, errors.New("DB_NAME environment variable not set")
	}

	intPort, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}

	return &DBConfig{
		Host:     host,
		Port:     intPort,
		User:     user,
		Password: password,
		Name:     name,
	}, nil
}
