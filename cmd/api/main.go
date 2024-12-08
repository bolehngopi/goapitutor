package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	// "github.com/bolehngopi/goapitutor/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting server on port 8080")

	err := http.ListenAndServe("localhost:8080", r)

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Failed to start server")
	}
}