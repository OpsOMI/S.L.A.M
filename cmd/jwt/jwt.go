package main

import (
	"fmt"
	"time"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/tokenstore"
	"github.com/google/uuid"
)

func main() {
	jwtManager := tokenstore.NewJWTManager("SLAM", "u549QD5weh9A04n")

	clientID := uuid.New()
	userID := uuid.New()
	username := "john_doe"
	nickname := "Johnny"

	duration := time.Hour * 24

	token, err := jwtManager.GenerateToken(clientID, userID, username, nickname, duration)
	if err != nil {
		fmt.Println("Token generate error:", err)
		return
	}

	fmt.Println("Generated JWT token:", token)
}
