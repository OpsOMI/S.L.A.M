package services

import (
	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/server/repositories"
	"github.com/OpsOMI/S.L.A.M/internal/server/services/users"
	"github.com/OpsOMI/S.L.A.M/internal/server/services/utils"
	"github.com/OpsOMI/S.L.A.M/pkg"
)

type IServices interface {
	Utils() utils.IUtilServices
	Users() users.IUserService
}

type services struct {
	utils utils.IUtilServices
	users users.IUserService
}

func NewServices(
	logger logger.ILogger,
	packages pkg.IPackages,
	repositories repositories.IRepositories,
) IServices {
	utils := utils.NewServices()
	users := users.NewService(utils, packages, repositories)

	return &services{
		utils: utils,
		users: users,
	}
}

func (s *services) Utils() utils.IUtilServices {
	return s.utils
}

func (s *services) Users() users.IUserService {
	return s.users
}
