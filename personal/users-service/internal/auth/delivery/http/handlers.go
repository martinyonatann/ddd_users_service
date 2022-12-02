package http

import (
	"github.com/koinworks/asgard-heimdal/libs/logger"
	"github.com/labstack/echo/v4"
	"github.com/martinyonathann/users-service/config"
	"github.com/martinyonathann/users-service/internal/auth"
	"github.com/martinyonathann/users-service/internal/session"
)

type authHandlers struct {
	cfg    *config.Config
	authUC auth.UseCase
	sessUC session.UCSession
	logger logger.Logger
}

func NewAuthHandlers(cfg *config.Config, authUC auth.UseCase, sessUC session.UCSession, log logger.Logger) auth.Handlers {
	return &authHandlers{cfg: cfg, authUC: authUC, sessUC: sessUC, logger: log}
}

func (h *authHandlers) Register() echo.HandlerFunc
func (h *authHandlers) Login() echo.HandlerFunc
func (h *authHandlers) Logout() echo.HandlerFunc
func (h *authHandlers) Update() echo.HandlerFunc
func (h *authHandlers) Delete() echo.HandlerFunc
func (h *authHandlers) GetUserByID() echo.HandlerFunc
func (h *authHandlers) FindByName() echo.HandlerFunc
func (h *authHandlers) GetUsers() echo.HandlerFunc
func (h *authHandlers) GetMe() echo.HandlerFunc
func (h *authHandlers) UploadAvatar() echo.HandlerFunc
func (h *authHandlers) GetCSRFToken() echo.HandlerFunc
