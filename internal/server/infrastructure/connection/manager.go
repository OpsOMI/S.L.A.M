package connection

import (
	"net"
	"sync"

	"github.com/google/uuid"
)

type ClientInfo struct {
	Conn     net.Conn
	UserID   uuid.UUID
	RoomCode string
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
		Conn:     conn,
		RoomCode: "",
	}

}

func (cm *ConnectionManager) SetClientRoom(
	clientID uuid.UUID,
	roomCode string,
) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	if client, ok := cm.clients[clientID]; ok {
		client.RoomCode = roomCode
	}
}

func (cm *ConnectionManager) GetClientRoom(clientID uuid.UUID) (string, bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	client, ok := cm.clients[clientID]
	if !ok || client.RoomCode == "" {
		return "", false
	}
	return client.RoomCode, true
}

func (cm *ConnectionManager) GetConnectionsByRoomCode(roomCode string) ([]net.Conn, bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	var conns []net.Conn
	for _, client := range cm.clients {
		if client.RoomCode == roomCode {
			conns = append(conns, client.Conn)
		}
	}

	if len(conns) == 0 {
		return nil, false
	}
	return conns, true
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

func (cm *ConnectionManager) CountOnlineConnections() int {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return len(cm.clients)
}
