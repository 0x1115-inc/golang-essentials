package conversion

import (
	"encoding/json"
)

// StructToStruct converts a struct to another struct
// It uses json.Marshal and json.Unmarshal to convert the struct
// The struct fields must be tagged with json
func StructToStruct(src interface{}, dst interface{}) error {
	jsonBytes, err := json.Marshal(src)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonBytes, dst)
}
