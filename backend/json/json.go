package json

import (
	"encoding/json"
	"log"
)

// ToJSON marshal bytes to JSON.
func ToJSON(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		log.Println(err.Error())
	}

	return string(bytes)
}
