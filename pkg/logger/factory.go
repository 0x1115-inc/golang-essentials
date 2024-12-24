package logger

import (
	"fmt"
	"strconv"
)

func GetLogger(provider string, args map[string]interface{}) VLogger {

	// Parameter validation
	logLevel, err := strconv.ParseInt(fmt.Sprintf("%s", args["level"]), 10, 64)
	if err != nil {
		return nil
	}

	if logLevel < LevelDebug || logLevel > LevelFatal {
		return nil
	}

	switch provider {
	case "simple":
		return NewSimpleLogger(int(logLevel))
	default:
		return nil
	}
}
