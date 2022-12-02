package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	authHttp "github.com/martinyonathann/users-service/internal/auth/delivery/http"
	"github.com/martinyonathann/users-service/internal/auth/repository"
	"github.com/martinyonathann/users-service/internal/auth/usecase"
	"github.com/martinyonathann/users-service/internal/middleware"
	"github.com/martinyonathann/users-service/pkg/metric"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	_, err := metric.CreateMetrics(s.cfg.Metrics.URL, s.cfg.Metrics.ServiceName)
	if err != nil {
		s.logger.Errorf("CreateMatrics Error: %s", err)
	}
	s.logger.Infof(
		"Metrics available URL: %s, ServiceName: %s",
		s.cfg.Metrics.URL,
		s.cfg.Metrics.ServiceName,
	)

	// Init repositories
	authRepo := repository.NewAuthRepository(s.db)

	// Init useCases
	authUC := usecase.NewAuthUseCase(s.cfg, authRepo, s.logger)

	// Init Handler
	authHandlers := authHttp.NewAuthHandlers(s.cfg, authUC, nil, s.logger)

	mw := middleware.NewMiddlewareManager(nil, authUC, s.cfg, []string{"*"}, s.logger)

	e.Use(mw.RequestLoggerMiddleware)

	//ToDo : create swagger docs

	//ToDo : validate ssl

	//ToDo : CORSWithConfig

	//ToDo : RecoverWithConfig

	v1 := e.Group("/api/v1")

	health := v1.Group("/health")
	authGroup := v1.Group("/auth")

	authHttp.MapAuthRoutes(authGroup, authHandlers, mw)

	health.GET("", func(c echo.Context) error {
		s.logger.Infof("Health check RequestID: %s", middleware.GetRequestID(c))
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})

	})

	return nil
}
