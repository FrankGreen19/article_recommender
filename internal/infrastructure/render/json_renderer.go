package render

import (
	"encoding/json"
	"net/http"
)

type JSONRenderer struct{}

func (r JSONRenderer) Render(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
