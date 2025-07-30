package domains

import (
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/utils"
)

type IMapper interface {
	Common() commons.ICommonMapper
	Utils() utils.IUtilMapper
}

type mappers struct {
	common commons.ICommonMapper
	utils  utils.IUtilMapper
}

func NewMappers() IMapper {
	common := commons.NewMapper()
	utils := utils.NewMapper()

	return &mappers{
		common: common,
		utils:  utils,
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
