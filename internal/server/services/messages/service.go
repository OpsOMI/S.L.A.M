package messages

import (
	"github.com/OpsOMI/S.L.A.M/internal/server/repositories"
	"github.com/OpsOMI/S.L.A.M/internal/server/services/rooms"
	"github.com/OpsOMI/S.L.A.M/internal/server/services/users"
	"github.com/OpsOMI/S.L.A.M/internal/server/services/utils"
	"github.com/OpsOMI/S.L.A.M/pkg"
)

type IMessageService interface {
}

type service struct {
	utils        utils.IUtilServices
	packages     pkg.IPackages
	repositories repositories.IRepositories
	users        users.IUserService
	rooms        rooms.IRoomService
}

func NewService(
	utils utils.IUtilServices,
	packages pkg.IPackages,
	repositories repositories.IRepositories,
	users users.IUserService,
	rooms rooms.IRoomService,
) IMessageService {
	return &service{
		utils:        utils,
		packages:     packages,
		repositories: repositories,
		users:        users,
		rooms:        rooms,
	}
}
