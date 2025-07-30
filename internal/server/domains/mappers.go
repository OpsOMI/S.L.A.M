package domains

import (
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/users"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/utils"
)

type IMapper interface {
	Common() commons.ICommonMapper
	Utils() utils.IUtilMapper
	Users() users.IUsersMapper
}

type mappers struct {
	common commons.ICommonMapper
	utils  utils.IUtilMapper
	users  users.IUsersMapper
}

func NewMappers() IMapper {
	common := commons.NewMapper()
	utils := utils.NewMapper()
	users := users.NewMapper(utils)

	return &mappers{
		common: common,
		utils:  utils,
		users:  users,
	}
}

// Common returns the common domain mapper.
func (m *mappers) Common() commons.ICommonMapper {
	return m.common
}

// Utils returns the util domain mapper.
func (m *mappers) Utils() utils.IUtilMapper {
	return m.utils
}

func (m *mappers) Users() users.IUsersMapper {
	return m.users
}
