package conversion

import (
	"encoding/json"
)

func StructToStruct(src interface{}, dst interface{}) error {
	jsonBytes, err := json.Marshal(src)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonBytes, dst)			
}
