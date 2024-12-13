package logger

func GetLogger(args map[string]interface{}) VLogger {
	switch args["provider"].(string) {
	case "simple":
		return NewSimpleLogger(args["level"].(int))
	default:
		return nil
	}	
}
