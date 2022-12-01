package web

import (
	"fmt"
	"github.com/Scalingo/sclng-backend-test-v1/internal/web/routes"
	"github.com/sirupsen/logrus"

	//	"github.com/Scalingo/go-handlers"
	"github.com/Scalingo/go-handlers"
	"net/http"
)

type Handler struct {
	router handlers.Router
	port   int
	log    logrus.FieldLogger
}

func NewHandler(log logrus.FieldLogger, port int) Handler {
	router := *handlers.NewRouter(log)
	router.HandleFunc("/ping", routes.PongHandler)
	// Initialize web server and configure the following routes:
	// GET /repos
	// GET /stats

	return Handler{
		router: router,
		port:   port,
		log:    log,
	}
}

func (handler Handler) Serve() {
	handler.log.WithField("port", handler.port).Info("Listening...")
	err := http.ListenAndServe(fmt.Sprintf(":%d", handler.port), handler.router)
	if err != nil {
		logrus.Errorln("Erreur lors du démarrage de l'application")
	}
}
