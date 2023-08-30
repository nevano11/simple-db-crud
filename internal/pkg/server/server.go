package server

import (
	"net/http"
	"simple-db-crud/internal/pkg/config"

	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	logger *logrus.Logger
	config *config.Config
	server *http.Server
}

func New(config *config.Config) (*ApiServer, error) {
	// * Create logger instance
	logger, err := configureLogger(config.Logger.LogLevel)
	if err != nil {
		return nil, err
	}

	// * Create server
	server := &http.Server{
		Addr: ":" + config.Server.Port,
	}

	// * Create apiserver instance
	apiServer := &ApiServer{
		logger: logger,
		config: config,
		server: server,
	}

	// * Add router
	apiServer.server.Handler = apiServer.configureRoutes()
	return apiServer, nil
}

func (s *ApiServer) Start() error {
	s.logger.Infof("Server is starting on http://localhost:%s", s.config.Server.Port)
	s.logger.Debugf("Application config is %s", s.config)

	return s.server.ListenAndServe()
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
