package models

import (
	"time"

	"github.com/google/uuid"
)

// Player represents a user of the application, whether registered or guest.
type Player struct {
	PlayerID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Username        string    `gorm:"uniqueIndex;not null"`
	ProfilePhoto    string
	AccountID       *uuid.UUID
	Account         *Account `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;"`
	Sessions        []Session
	GameParticipants []GameParticipant
	WonGames        []Game `gorm:"foreignKey:WinnerID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Account represents the registered credentials of a player.
type Account struct {
	AccountID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email           string    `gorm:"uniqueIndex;not null"`
	PasswordHash    *string
	EmailVerifiedAt *time.Time
	PlayerID        uuid.UUID `gorm:"uniqueIndex;not null"`
	Player          *Player   `gorm:"foreignKey:PlayerID;constraint:OnDelete:CASCADE;"`
	OAuthIdentities []OAuthIdentity
	LocalCredential *LocalCredential
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// OAuthIdentity stores an external OAuth provider identity linked to an account.
type OAuthIdentity struct {
	IdentityID     uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Provider       string    `gorm:"not null"`
	ProviderUserID string    `gorm:"not null"`
	AccountID      uuid.UUID `gorm:"not null"`
	Account        *Account  `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// LocalCredential stores the password hash for accounts using email/password login.
type LocalCredential struct {
	CredentialID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	PasswordHash string    `gorm:"not null"`
	AccountID    uuid.UUID `gorm:"uniqueIndex;not null"`
	Account      *Account  `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Session represents an authenticated session for a player.
type Session struct {
	SessionID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TokenHash string    `gorm:"not null"`
	CreatedAt time.Time
	ExpiresAt time.Time
	RevokedAt *time.Time
	PlayerID  uuid.UUID `gorm:"not null"`
	Player    *Player   `gorm:"foreignKey:PlayerID;constraint:OnDelete:CASCADE;"`
}

// Lobby represents a game lobby.
type Lobby struct {
	LobbyID      uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Code         string    `gorm:"uniqueIndex;not null"`
	Status       string    `gorm:"not null;default:'waiting'"`
	IsPublic     bool      `gorm:"default:false"`
	CreatorID    uuid.UUID `gorm:"not null"`
	Creator      *Player   `gorm:"foreignKey:CreatorID;constraint:OnDelete:CASCADE;"`
	HostID       uuid.UUID `gorm:"not null"`
	Host         *Player   `gorm:"foreignKey:HostID;constraint:OnDelete:CASCADE;"`
	Participants []LobbyParticipant
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// LobbyParticipant links a player to a lobby with their role and status.
type LobbyParticipant struct {
	LobbyParticipantID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	LobbyID            uuid.UUID `gorm:"not null"`
	Lobby              *Lobby    `gorm:"foreignKey:LobbyID;constraint:OnDelete:CASCADE;"`
	PlayerID           uuid.UUID `gorm:"not null"`
	Player             *Player   `gorm:"foreignKey:PlayerID;constraint:OnDelete:CASCADE;"`
	Role               string    `gorm:"not null;default:'player'"`
	Status             string    `gorm:"not null;default:'active'"`
	JoinedAt           time.Time
	ExitAt             *time.Time
}

// TableName overrides the table name for LobbyParticipant.
func (LobbyParticipant) TableName() string {
	return "lobby_participants"
}

// Game represents a single played game inside a lobby.
type Game struct {
	GameID     uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	LobbyID    uuid.UUID `gorm:"not null"`
	Lobby      *Lobby    `gorm:"foreignKey:LobbyID;constraint:OnDelete:CASCADE;"`
	Status     string    `gorm:"not null;default:'waiting'"`
	StartedAt  *time.Time
	FinishedAt *time.Time
	WinnerID   *uuid.UUID
	Winner     *Player `gorm:"foreignKey:WinnerID;constraint:OnDelete:SET NULL;"`
	Participants []GameParticipant
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// GameParticipant links a player to a specific game.
type GameParticipant struct {
	GameParticipantID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	GameID            uuid.UUID `gorm:"not null"`
	Game              *Game     `gorm:"foreignKey:GameID;constraint:OnDelete:CASCADE;"`
	PlayerID          uuid.UUID `gorm:"not null"`
	Player            *Player   `gorm:"foreignKey:PlayerID;constraint:OnDelete:CASCADE;"`
	PlayerIndex       int       `gorm:"not null"`
	FinalScore        *int
	HasQuit           bool `gorm:"default:false"`
	ReplacedByAI      bool `gorm:"default:false"`
}

// TableName overrides the table name for GameParticipant.
func (GameParticipant) TableName() string {
	return "game_participants"
}
