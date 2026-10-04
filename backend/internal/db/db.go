package db

import (
	"backend/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewConnection opens a connection to the PostgreSQL database using the provided DSN.
func NewConnection(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// AutoMigrate creates or updates the database schema to match the GORM models.
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Player{},
		&models.Account{},
		&models.OAuthIdentity{},
		&models.LocalCredential{},
		&models.Session{},
		&models.Lobby{},
		&models.LobbyParticipant{},
		&models.Game{},
		&models.GameParticipant{},
	)
}
