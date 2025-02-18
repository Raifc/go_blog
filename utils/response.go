package utils

import (
    "encoding/json"
    "net/http"
)

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
    response, err := json.Marshal(payload)
    if err != nil {
        RespondWithError(w, NewAppError("Failed to encode response", http.StatusInternalServerError))
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(response)
}
