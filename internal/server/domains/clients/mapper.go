package clients

import (
	"github.com/OpsOMI/S.L.A.M/internal/adapters/postgres/sqlc/pgqueries"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/utils"
	"github.com/google/uuid"
)

type IClientsMapper interface {
	One(
		dbModel *pgqueries.Client,
	) *Client

	Many(
		dbModels []pgqueries.Client,
		count int64,
	) *Clients

	CreateClient(
		userID, clientKey uuid.UUID,
	) pgqueries.CreateClientParams
}

type mapper struct {
	utils utils.IUtilMapper
}

// NewMapper creates a new instance of userAuthMapper.
func NewMapper(
	utilsMapper utils.IUtilMapper,
) *mapper {
	return &mapper{
		utils: utilsMapper,
	}
}

func (m *mapper) One(
	dbModel *pgqueries.Client,
) *Client {
	if dbModel == nil {
		return nil
	}

	return &Client{
		ID:        dbModel.ID,
		ClientKey: dbModel.ClientKey,
		RevokedAt: &dbModel.RevokedAt.Time,
		CreatedAt: dbModel.CreatedAt.Time,
	}
}

func (m *mapper) Many(
	dbModels []pgqueries.Client,
	count int64,
) *Clients {
	var appModels []Client
	for _, dbModel := range dbModels {
		appModels = append(appModels, *m.One(&dbModel))
	}

	return &Clients{
		Items:      appModels,
		TotalCount: count,
	}
}

func (m *mapper) CreateClient(
	userID, clientKey uuid.UUID,
) pgqueries.CreateClientParams {
	return pgqueries.CreateClientParams{
		UserID:    userID,
		ClientKey: clientKey,
	}
}
