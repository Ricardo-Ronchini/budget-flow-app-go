package common

import (
	"os"
	"strings"
)

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func GetEnvArray(key string, defaultValue []string) []string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	items := strings.Split(value, ",")
	for i, item := range items {
		items[i] = strings.TrimSpace(item)
	}

	return items
}
