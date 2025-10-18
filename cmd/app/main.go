package main


import (
	"fmt"
	"log"

	"practice4-sqlx/internal/db"
	"practice4-sqlx/internal/models"
	"practice4-sqlx/internal/repository"
	"practice4-sqlx/internal/service"
)

func main() {
	cfg := db.Config{
		Host:     "localhost",
		Port:     5433,
		User:     "postgres",
		Password: "postgres",
		DBName:   "practice4",
	}
	database, err := db.NewDB(cfg)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	defer database.Close()

	userRepo := repository.NewUserRepository(database)
	transferSvc := service.NewTransferService(database)

	newUser := models.User{Name: "Alice", Email: "alice@example.com", Balance: 100}
	if err := userRepo.Insert(newUser); err != nil {
		log.Println("Insert error:", err)
	}

	users, err := userRepo.GetAll()
	if err != nil {
		log.Println("GetAll error:", err)
	} else {
		fmt.Println("All users:")
		for _, u := range users {
			fmt.Printf("ID: %d | Name: %s | Email: %s | Balance: %.2f\n", u.ID, u.Name, u.Email, u.Balance)
		}
	}

	if err := transferSvc.TransferBalance(1, 2, 25.0); err != nil {
		log.Println("Transfer error:", err)
	}
}
