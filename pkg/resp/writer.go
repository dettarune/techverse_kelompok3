package resp

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, code int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	jsonData, _ := json.Marshal(data)
	w.Write(jsonData)
}
