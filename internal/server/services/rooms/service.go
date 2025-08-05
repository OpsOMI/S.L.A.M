package rooms

import (
	"github.com/OpsOMI/S.L.A.M/internal/server/repositories"
	"github.com/OpsOMI/S.L.A.M/internal/server/services/utils"
	"github.com/OpsOMI/S.L.A.M/pkg"
)

type IRoomService interface {
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
