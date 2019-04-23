package apiv1

import (
    "net/http"
    Log "github.com/victoryang/api-gateway/log"

    "github.com/gorilla/mux"
)

func RegisterAPIv1(r *mux.Router) http.Handler {
    Log.Debug("Register v1 api...")

    r.HandleFunc("/", hello).Methods("GET")
    r.HandleFunc("/health", handleHealth).Methods("GET")

    return r
}