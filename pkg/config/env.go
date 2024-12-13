package config

import (
	"os"
	"strconv"
	"strings"
)

const (
	EnvDefaultDelimiter = ","
)

type EnvConfig struct{
	delimiter string
}

func (e *EnvConfig) GetInt(key string) int {
	val, _ := strconv.Atoi(os.Getenv(key))
	return val
}

func (e *EnvConfig) GetInt64(key string) int64 {
	val, _ := strconv.ParseInt(os.Getenv(key), 10, 64)
	return val
}

func (e *EnvConfig) GetUint64(key string) uint64 {
	val, _ := strconv.ParseUint(os.Getenv(key), 10, 64)
	return val
}

func (e *EnvConfig) GetFloat64(key string) float64 {
	val, _ := strconv.ParseFloat(os.Getenv(key), 64)
	return val
}

func (e *EnvConfig) GetString(key string) string {
	return os.Getenv(key)
}

func (e *EnvConfig) GetStringSlice(key string) []string {
	return strings.Split(os.Getenv(key), e.delimiter)
}

func (e *EnvConfig) GetStringMap(key string) map[string]interface{} {
	result := make(map[string]interface{})
	pairs := strings.Split(os.Getenv(key), e.delimiter)
	for _, pair := range pairs {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) == 2 {
			result[kv[0]] = kv[1]
		}
	}
	return result
}

func NewEnvConfig() *EnvConfig {
	return &EnvConfig{
		delimiter: EnvDefaultDelimiter,
	}
}