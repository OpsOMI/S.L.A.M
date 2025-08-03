package users

import (
	"context"

	"github.com/OpsOMI/S.L.A.M/internal/server/domains/users"
	"github.com/OpsOMI/S.L.A.M/internal/server/repositories"
	"github.com/OpsOMI/S.L.A.M/internal/server/services/utils"
	"github.com/OpsOMI/S.L.A.M/pkg"
	"github.com/google/uuid"
)

type IUserService interface {
	Login(
		ctx context.Context,
		username, password string,
	) (*users.User, error)

	GetByID(
		ctx context.Context,
		id string,
	) (*users.User, error)

	GetByUsername(
		ctx context.Context,
		username string,
	) (*users.User, error)

	GetByNickname(
		ctx context.Context,
		nickname string,
	) (*users.User, error)

	GetByPrivateCode(
		ctx context.Context,
		privateCode string,
	) (*users.User, error)

	CreateUser(
		ctx context.Context,
		nickname, privateCode, username, password, role string,
	) (*uuid.UUID, error)

	ChangeNickname(
		ctx context.Context,
		id, nickname string,
	) error

	DeleteByID(
		ctx context.Context,
		id string,
	) error
}

type service struct {
	utils        utils.IUtilServices
	packages     pkg.IPackages
	repositories repositories.IRepositories
}

func NewService(
	utils utils.IUtilServices,
	packages pkg.IPackages,
	repositories repositories.IRepositories,
) IUserService {
	return &service{
		utils:        utils,
		packages:     packages,
		repositories: repositories,
	}
}
