package util

import (
	"encoding/json"
)

// some data need to be parsed into struct format
// the data is fetched from possibly database or over http request.
// that data need to be converted to a series of bytes ie an array of bytes
// then easily convert it back to a struct with the help of marshaler

func StructMarshaler(source interface{}, target interface{}) error {
	// turning the data to byte
	dataBytes, _ := json.Marshal(source)

	// parsing it back to a struct
	return json.Unmarshal(dataBytes, target)
}
