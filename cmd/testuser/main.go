package main

import (
	"context"
	"time"
	"workflou/template/auth"
	"workflou/template/db"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	db := db.New()
	defer db.Close()

	auth.NewStore(db).Save(context.Background(), auth.User{
		Email:     "test@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}
