package service

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"practice4-sqlx/internal/models"
)

type TransferService struct {
	db *sqlx.DB
}

func NewTransferService(db *sqlx.DB) *TransferService {
	return &TransferService{db: db}
}

func (s *TransferService) TransferBalance(fromID, toID int, amount float64) error {
	tx, err := s.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	var sender, receiver models.User

	if err := tx.Get(&sender, "SELECT * FROM users WHERE id=$1 FOR UPDATE", fromID); err != nil {
		tx.Rollback()
		return fmt.Errorf("sender not found: %w", err)
	}
	if err := tx.Get(&receiver, "SELECT * FROM users WHERE id=$1 FOR UPDATE", toID); err != nil {
		tx.Rollback()
		return fmt.Errorf("receiver not found: %w", err)
	}
	if sender.Balance < amount {
		tx.Rollback()
		return fmt.Errorf("insufficient balance for user ID %d", fromID)
	}

	if _, err := tx.Exec("UPDATE users SET balance = balance - $1 WHERE id=$2", amount, fromID); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update sender: %w", err)
	}
	if _, err := tx.Exec("UPDATE users SET balance = balance + $1 WHERE id=$2", amount, toID); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update receiver: %w", err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Printf("💸 Transferred %.2f from user %d to user %d successfully.\n", amount, fromID, toID)
	return nil
}
