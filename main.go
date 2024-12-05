package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"prechecks/api/v1"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"prechecks/pkg/logger"

	"prechecks/conf"
)

func main() {

	logger.InitLogger("app.log")

	// Initialize configration
	conf := conf.New()

	// Initialize router
	r := chi.NewRouter()

	// V1 - API
	r.Route("/api", func(r chi.Router) {

		r.Route("/v1", func(r chi.Router) {

			// Group unprotected routess
			r.Group(func(r chi.Router) {
				r.Mount("/", v1.Router())
			})

		})

	})

	// Listening
	appPort := fmt.Sprintf(":%d", conf.AppPort)
	logger.Log.Info("Listening at :%d", conf.AppPort)

	listener, err := net.Listen("tcp", appPort)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err,
		}).Error(fmt.Sprintf("Failed to listen on port %d\n", appPort))
		os.Exit(0)
	}

	if err = http.Serve(listener, r); err != nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err,
		}).Error(fmt.Sprintf("Failed to serve on port %s", appPort))
		os.Exit(0)
	}
}
