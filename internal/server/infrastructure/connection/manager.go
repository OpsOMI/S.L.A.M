package connection

import (
	"net"
	"sync"
)

type ClientInfo struct {
	Conn   net.Conn
	RoomID string
}

type ConnectionManager struct {
	clients map[string]*ClientInfo
	mu      sync.RWMutex
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		clients: make(map[string]*ClientInfo),
	}
}

func (cm *ConnectionManager) Register(
	clientID string,
	conn net.Conn,
) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.clients[clientID] = &ClientInfo{
		Conn:   conn,
		RoomID: "",
	}

}

func (cm *ConnectionManager) SetClientRoom(clientID, roomID string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	if client, ok := cm.clients[clientID]; ok {
		client.RoomID = roomID
	}
}

func (cm *ConnectionManager) GetClientRoom(clientID string) (string, bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	if client, ok := cm.clients[clientID]; ok {
		return client.RoomID, true
	}
	return "", false
}

func (cm *ConnectionManager) Unregister(clientID string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	delete(cm.clients, clientID)
}

func (cm *ConnectionManager) GetConn(clientID string) (net.Conn, bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	if client, ok := cm.clients[clientID]; ok {
		return client.Conn, true
	}
	return nil, false
}
