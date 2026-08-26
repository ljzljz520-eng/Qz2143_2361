package api

import (
	"community50/model"
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}
func decodeQuery(r *http.Request) model.Query {
	return model.Query{Status: r.URL.Query().Get("status"), Author: r.URL.Query().Get("author"), Text: r.URL.Query().Get("q")}
}
