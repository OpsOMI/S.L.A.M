package clients

import (
	"context"

	"github.com/OpsOMI/S.L.A.M/internal/server/config"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/clients"
	"github.com/OpsOMI/S.L.A.M/internal/server/repositories"
	"github.com/OpsOMI/S.L.A.M/internal/server/services/utils"
	"github.com/OpsOMI/S.L.A.M/pkg"
)

type IClientService interface {
	GetByID(
		ctx context.Context,
		id string,
	) (*clients.Client, error)

	GetByClientKey(
		ctx context.Context,
		clientKey string,
	) (*clients.Client, error)

	GetByUserID(
		ctx context.Context,
		userID string,
	) (*clients.Clients, error)

	RevokeByID(
		ctx context.Context,
		id string,
	) error

	DeleteByID(
		ctx context.Context,
		id string,
	) error

	IsExistByID(
		ctx context.Context,
		id string,
	) (*bool, error)

	IsExistByClientKey(
		ctx context.Context,
		clientKey string,
	) (*bool, error)

	IsRevoked(
		ctx context.Context,
		id string,
	) (*bool, error)

	CreateClient(
		serverConfig *config.Configs,
		clientID, nickname string,
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
) IClientService {
	return &service{
		utils:        utils,
		packages:     packages,
		repositories: repositories,
	}
}
