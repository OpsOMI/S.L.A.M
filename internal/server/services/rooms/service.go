package rooms

import (
	"context"

	"github.com/OpsOMI/S.L.A.M/internal/server/domains/rooms"
	"github.com/OpsOMI/S.L.A.M/internal/server/repositories"
	"github.com/OpsOMI/S.L.A.M/internal/server/services/utils"
	"github.com/OpsOMI/S.L.A.M/pkg"
)

type IRoomService interface {
	GetByID(
		ctx context.Context,
		id string,
	) (*rooms.Room, error)

	GetByCode(
		ctx context.Context,
		code string,
	) (*rooms.Room, error)

	GetByOwnerID(
		ctx context.Context,
		ownerID string,
	) (*rooms.Rooms, error)

	DeleteByID(
		ctx context.Context,
		id string,
	) error

	IsExistByID(
		ctx context.Context,
		id string,
	) (*bool, error)

	IsExistByCode(
		ctx context.Context,
		code string,
	) (*bool, error)

	IsExistByOwnerID(
		ctx context.Context,
		ownerID string,
	) (*bool, error)
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
) IRoomService {
	return &service{
		utils:        utils,
		packages:     packages,
		repositories: repositories,
	}
}
