package handler

import "encoding/json"

func jsonError(msg string) []byte {
	err := struct {
		Message string `json:"message"`
	}{msg}
	b, error := json.Marshal(err)
	if error != nil {
		return []byte(error.Error())
	}
	return b
}
