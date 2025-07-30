package users

import "net"

type Connection struct {
	Conn        net.Conn
	User        *User
	CurrentRoom string
}
