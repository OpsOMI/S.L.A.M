package connection

import (
	"net"
	"sync"

	"github.com/google/uuid"
)

type ClientInfo struct {
	Conn   net.Conn
	RoomID string
}

type ConnectionManager struct {
	clients map[uuid.UUID]*ClientInfo
	mu      sync.RWMutex
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		clients: make(map[uuid.UUID]*ClientInfo),
	}
}

func (cm *ConnectionManager) Register(
	clientID uuid.UUID,
	conn net.Conn,
) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.clients[clientID] = &ClientInfo{
		Conn:   conn,
		RoomID: "",
	}

}

func (cm *ConnectionManager) SetClientRoom(clientID uuid.UUID, roomID string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	if client, ok := cm.clients[clientID]; ok {
		client.RoomID = roomID
	}
}

func (cm *ConnectionManager) GetClientRoom(clientID uuid.UUID) (string, bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	if client, ok := cm.clients[clientID]; ok {
		return client.RoomID, true
	}
	return "", false
}

func (cm *ConnectionManager) Unregister(clientID uuid.UUID) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	delete(cm.clients, clientID)
}

func (cm *ConnectionManager) GetConn(clientID uuid.UUID) (net.Conn, bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	if client, ok := cm.clients[clientID]; ok {
		return client.Conn, true
	}
	return nil, false
}
