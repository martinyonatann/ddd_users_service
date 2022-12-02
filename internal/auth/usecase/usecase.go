package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/martinyonathann/users-service/config"
	"github.com/martinyonathann/users-service/internal/auth"
	"github.com/martinyonathann/users-service/internal/models"
	"github.com/martinyonathann/users-service/pkg/logger"
	"github.com/martinyonathann/users-service/pkg/utils"
)

const (
	basePrefix    = "api-auth"
	cacheDuration = 3600
)

type authUC struct {
	cfg      *config.Config
	authRepo auth.Repository
	logger   logger.Logger
}

func NewAuthUseCase(cfg *config.Config, authRepo auth.Repository, logger logger.Logger) auth.UseCase {
	return &authUC{cfg: cfg, authRepo: authRepo, logger: logger}
}

func (u *authUC) Register(ctx context.Context, user *models.User) (*models.UserWithToken, error)
func (u *authUC) Login(ctx context.Context, user *models.User) (*models.UserWithToken, error)
func (u *authUC) Update(ctx context.Context, user *models.User) (*models.User, error)
func (u *authUC) Delete(ctx context.Context, userID uuid.UUID) error
func (u *authUC) GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
func (u *authUC) FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*models.UsersList, error)
func (u *authUC) GetUsers(ctx context.Context, pq *utils.PaginationQuery) (*models.UsersList, error)
func (u *authUC) UploadAvatar(ctx context.Context, userID uuid.UUID, file models.UploadInput) (*models.User, error)
