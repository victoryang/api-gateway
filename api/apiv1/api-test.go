package apiv1

import (
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    WriteSuccessResponse(w, "Welcome to api-gateway")
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
    WriteSuccessResponse(w, "echo OK")
}