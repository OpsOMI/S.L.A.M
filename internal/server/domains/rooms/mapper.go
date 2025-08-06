package rooms

import (
	"github.com/OpsOMI/S.L.A.M/internal/adapters/postgres/sqlc/pgqueries"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/utils"
	"github.com/google/uuid"
)

type IRoomsMapper interface {
	One(
		dbModel *pgqueries.Room,
	) *Room

	Many(
		dbModels []pgqueries.Room,
		count int64,
	) *Rooms

	CreateParams(
		ownerID uuid.UUID,
		code, password string,
	) pgqueries.CreateRoomParams
}

type mapper struct {
	utils utils.IUtilMapper
}

// NewMapper creates a new instance of rooms mapper.
func NewMapper(
	utils utils.IUtilMapper,
) *mapper {
	return &mapper{
		utils: utils,
	}
}

func (m *mapper) One(
	dbModel *pgqueries.Room,
) *Room {
	if dbModel == nil {
		return nil
	}

	return &Room{
		ID:        dbModel.ID,
		OwnerID:   dbModel.OwnerID,
		Code:      dbModel.Code,
		CreatedAt: dbModel.CreatedAt.Time,
	}
}

func (m *mapper) Many(
	dbModels []pgqueries.Room,
	count int64,
) *Rooms {
	var appModels []Room
	for _, dbModel := range dbModels {
		appModels = append(appModels, *m.One(&dbModel))
	}

	return &Rooms{
		Items:      appModels,
		TotalCount: count,
	}
}

func (m *mapper) CreateParams(
	ownerID uuid.UUID,
	code, password string,
) pgqueries.CreateRoomParams {
	return pgqueries.CreateRoomParams{
		OwnerID:  ownerID,
		Code:     code,
		Password: password,
	}
}
