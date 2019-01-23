package server

import (
	"fmt"
	"net/http"

	"github.com/agparadiso/tfcacoo/pkg/typeform"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func New() http.Handler {
	fmt.Println("Running cacco-typeform-integration server")
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1/").Subrouter()
	api.Handle(`/getfeedback`, typeform.NewHTTPHandler())

	handler := cors.Default().Handler(r)
	return handler
}
