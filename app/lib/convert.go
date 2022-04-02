package lib

import "encoding/json"

func StructToMap(obj interface{}) (value map[string]interface{}, err error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &value)
	return
}
