package server

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	logger *logrus.Logger
	server *http.Server
}

func New(logLevel, port string, handler http.Handler) (*ApiServer, error) {
	// * Create logger instance
	logger, err := configureLogger(logLevel)
	if err != nil {
		return nil, err
	}

	// * Create apiserver instance
	apiServer := &ApiServer{
		logger: logger,
		server: &http.Server{
			Addr:              ":" + port,
			Handler:           handler,
			ReadHeaderTimeout: 2 << 20,
			ReadTimeout:       10 * time.Second,
			WriteTimeout:      10 * time.Second,
		},
	}

	return apiServer, nil
}

func (s *ApiServer) Start() error {
	s.logger.Info("Server starting on addr http://localhost" + s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *ApiServer) Shutdown(ctx context.Context) {
	s.Shutdown(ctx)
}

func configureLogger(logLevel string) (*logrus.Logger, error) {
	// * Create logger instance
	logger := logrus.New()

	// * Set logLevel
	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}
	logger.SetLevel(lvl)

	return logger, nil
}
