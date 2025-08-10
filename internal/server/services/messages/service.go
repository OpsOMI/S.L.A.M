package messages

import (
	"context"

	"github.com/OpsOMI/S.L.A.M/internal/server/domains/messages"
	"github.com/OpsOMI/S.L.A.M/internal/server/repositories"
	"github.com/OpsOMI/S.L.A.M/internal/server/services/rooms"
	"github.com/OpsOMI/S.L.A.M/internal/server/services/users"
	"github.com/OpsOMI/S.L.A.M/internal/server/services/utils"
	"github.com/OpsOMI/S.L.A.M/pkg"
	"github.com/google/uuid"
)

type IMessageService interface {
	GetMessagesByRoomCode(
		ctx context.Context,
		roomCode string,
	) (*messages.RoomMessages, error)

	CreateMessage(
		ctx context.Context,
		senderID, roomCode string,
		content string,
	) error

	DeleteMessages(
		ctx context.Context,
	) error

	DeleteMessageByRoomCode(
		ctx context.Context,
		roomCode string,
	) error

	DeleteMessageInRoom(
		ctx context.Context,
		ownerID uuid.UUID,
		roomCode string,
	) error
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
