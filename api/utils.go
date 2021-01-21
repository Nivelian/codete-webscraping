package api

import (
	"encoding/json"
	"fmt"
	"github.com/Nivelian/codete-webscraping/helpers"
	"net/http"
)

func InternalServerErr(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, err)
}

func SendStruct(x interface{}, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(x); err != nil {
		return helpers.LogErr(err, "Failed to send struct to response")
	}

	return nil
}
