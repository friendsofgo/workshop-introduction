package encoding

import "encoding/json"

func JsonUnmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func JsonMarshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}