package utils

import (
	"encoding/json"
	"io/ioutil"
)

func ReadJSON(filename string, data interface{}) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	if err := json.Unmarshal(bytes, data); err != nil {
		return
	}
	return
}
