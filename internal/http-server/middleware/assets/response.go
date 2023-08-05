package assets

import (
	"encoding/json"
	"net/http"
)

func ErrorResponse[T any](w http.ResponseWriter, message T, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]T)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
