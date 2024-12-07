package config

// VConfig is an interface that defines methods for retrieving configuration values.
// It provides methods to get values of various types such as int, int64, uint64, float64,
// string, string slice, and string map.
//
// Methods:
//   - GetInt(key string) int: Retrieves an integer value for the given key.
//   - GetInt64(key string) int64: Retrieves an int64 value for the given key.
//   - GetUint64(key string) uint64: Retrieves a uint64 value for the given key.
//   - GetFloat64(key string) float64: Retrieves a float64 value for the given key.
//   - GetString(key string) string: Retrieves a string value for the given key.
//   - GetStringSlice(key string) []string: Retrieves a slice of strings for the given key.
//   - GetStringMap(key string) map[string]interface{}: Retrieves a map of string to interface{} for the given key.
type VConfig interface {
	GetInt(string) int
	GetInt64(string) int64
	GetUint64(string) uint64
	GetFloat64(string) float64
	GetString(string) string
	GetStringSlice(string) []string
	GetStringMap(string) map[string]interface{}
}
