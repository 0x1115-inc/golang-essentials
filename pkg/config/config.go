package config

type VConfig interface {
	GetInt(string) int
	GetInt64(string) int64
	GetUint64(string) uint64
	GetFloat64(string) float64
	GetString(string) string
	GetStringSlice(string) []string
	GetStringMap(string) map[string]interface{}
}
