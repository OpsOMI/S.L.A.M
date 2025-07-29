package domains

import "time"

type User struct {
	ID         string
	Nickname   string
	PrivateKey string
	CreatedAt  time.Time
}
