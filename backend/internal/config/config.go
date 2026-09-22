package config

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

type Config struct {
	Port string
}

func GetConfig() (Config, error) {
	envPath := os.Getenv("ENV_FILE")
	if envPath == "" {
		envPath = ".env"
	}

	f, err := os.Open(envPath)
	if err != nil {
		slog.Warn("Env file not found, using process environment", "path", envPath)
		return Config{
			Port: os.Getenv("PORT"),
		}, nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			slog.Info("Skipping malformed env line", "line", line)
			continue
		}
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		if key == "" {
			continue
		}
		if os.Getenv(key) == "" {
			os.Setenv(key, value)
		}
	}
	if err := scanner.Err(); err != nil {
		return Config{}, fmt.Errorf("read env file %s: %w", envPath, err)
	}

	return Config{
		Port: os.Getenv("PORT"),
	}, nil
}
