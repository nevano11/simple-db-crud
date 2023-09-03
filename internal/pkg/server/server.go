package server

import (
	"context"
	"net/http"
	"simple-db-crud/internal/pkg/config"
	"time"

	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	logger *logrus.Logger
	config *config.Config
	server *http.Server
}

func New(config *config.Config, handler http.Handler) (*ApiServer, error) {
	// * Create logger instance
	logger, err := configureLogger(config.Logger.LogLevel)
	if err != nil {
		return nil, err
	}

	// * Create apiserver instance
	apiServer := &ApiServer{
		logger: logger,
		config: config,
		server: &http.Server{
			Addr:              ":" + config.Server.Port,
			Handler:           handler,
			ReadHeaderTimeout: 2 << 20,
			ReadTimeout:       10 * time.Second,
			WriteTimeout:      10 * time.Second,
		},
	}

	return apiServer, nil
}

func (s *ApiServer) Start() error {
	s.logger.Infof("Server is starting on http://localhost:%s", s.config.Server.Port)
	s.logger.Debugf("Application config is %s", s.config)

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
