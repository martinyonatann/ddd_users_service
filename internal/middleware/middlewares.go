package middleware

import (
	"github.com/martinyonathann/users-service/config"
	"github.com/martinyonathann/users-service/internal/auth"
	"github.com/martinyonathann/users-service/internal/session"
	"github.com/martinyonathann/users-service/pkg/logger"
)

type MiddlewareManager struct {
	sessUC  session.UCSession
	authUC  auth.UseCase
	cfg     *config.Config
	origins []string
	logger  logger.Logger
}

// Middleware manager constructor
func NewMiddlewareManager(sessUC session.UCSession, authUC auth.UseCase, cfg *config.Config, origins []string, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{sessUC: sessUC, authUC: authUC, cfg: cfg, origins: origins, logger: logger}
}
