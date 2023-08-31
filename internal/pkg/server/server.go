package server

import (
	"simple-db-crud/internal/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	logger *logrus.Logger
	config *config.Config
	engine *gin.Engine
}

func New(config *config.Config) (*ApiServer, error) {
	// * Create logger instance
	logger, err := configureLogger(config.Logger.LogLevel)
	if err != nil {
		return nil, err
	}

	// * Create server
	engine := gin.Default()

	// * Create apiserver instance
	apiServer := &ApiServer{
		logger: logger,
		config: config,
		engine: engine,
	}

	// * Add router
	apiServer.configureRoutes()
	return apiServer, nil
}

func (s *ApiServer) Start() error {
	s.logger.Infof("Server is starting on http://localhost:%s", s.config.Server.Port)
	s.logger.Debugf("Application config is %s", s.config)

	return s.engine.Run(":" + s.config.Server.Port)
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
