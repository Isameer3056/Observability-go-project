package rest

import (
	"net/http"

	"github.com/Isameer3056/Observability-go-project/pkg/di"
)

func Initialize() {
	exampleController := di.NewExampleController()
	mux := http.NewServeMux()

	mux.HandleFunc("/", exampleController.GetExample)
	http.ListenAndServe(":8080", mux)
}
