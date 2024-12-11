package cache

import "fmt"

// Error code dictionary
const (
	CacheErrorNotFound = 1
)

type CacheError struct {
	Code int
	Message string
}

func (e *CacheError) Error() string {
	return fmt.Sprintf("[Code %d]: %s", e.Code, e.Message)
}

