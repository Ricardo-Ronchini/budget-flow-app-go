package common

import (
	"math"
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

func RoundTo(n float32, places int) float32 {
	pow := math.Pow(10, float64(places))
	return float32(math.Round(float64(n)*pow) / pow)
}
