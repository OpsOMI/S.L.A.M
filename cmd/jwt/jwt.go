package main

import (
	"fmt"
	"time"

	"github.com/OpsOMI/S.L.A.M/internal/server/network/store"
	"github.com/google/uuid"
)

func main() {
	jwtManager := store.NewManager("SLAM", "u549QD5weh9A04n")

	clientID := uuid.New()
	userID := uuid.New()
	username := "john_doe"
	nickname := "Johnny"

	duration := time.Hour * 24

	token, err := jwtManager.GenerateToken(clientID, userID, username, nickname, "user", duration)
	if err != nil {
		fmt.Println("Token generate error:", err)
		return
	}

	fmt.Println("Generated JWT token:", token)
}
