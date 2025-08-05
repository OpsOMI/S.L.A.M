package services

import (
	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/server/repositories"
	"github.com/OpsOMI/S.L.A.M/internal/server/services/clients"
	"github.com/OpsOMI/S.L.A.M/internal/server/services/rooms"
	"github.com/OpsOMI/S.L.A.M/internal/server/services/users"
	"github.com/OpsOMI/S.L.A.M/internal/server/services/utils"
	"github.com/OpsOMI/S.L.A.M/pkg"
)

type IServices interface {
	Utils() utils.IUtilServices
	Users() users.IUserService
	Clients() clients.IClientService
	Rooms() rooms.IRoomService
}

type services struct {
	utils   utils.IUtilServices
	users   users.IUserService
	clients clients.IClientService
	rooms   rooms.IRoomService
}

func NewServices(
	logger logger.ILogger,
	packages pkg.IPackages,
	repositories repositories.IRepositories,
) IServices {
	utils := utils.NewServices()
	users := users.NewService(utils, packages, repositories)
	clients := clients.NewService(utils, packages, repositories)
	rooms := rooms.NewService(utils, packages, repositories)

	return &services{
		utils:   utils,
		users:   users,
		clients: clients,
		rooms:   rooms,
	}
}

func (s *services) Utils() utils.IUtilServices {
	return s.utils
}

func (s *services) Users() users.IUserService {
	return s.users
}

func (s *services) Clients() clients.IClientService {
	return s.clients
}

func (s *services) Rooms() rooms.IRoomService {
	return s.rooms
}
