package services

import (
	"time"

	"github.com/keylab/celestialbound/backend/models"
	"github.com/keylab/celestialbound/backend/utils"
)

// In-memory storage for player states
var playerStates = make(map[string]*models.Player)

type PlayerService struct{}

func (playerService *PlayerService) CreatePlayer(playerName string) string {

	// Generate a unique PlayerID
	playerID := utils.GenerateUniqueID()

	//Generate Jar unique ID
	jarID := utils.GenerateUniqueID()

	// Default jar
	starterJar := models.Jar{
		JarID:          jarID,
		JarLevel:       1,
		StarsStored:    0,
		StarsPerSecond: 2,
		UpgradeCost:    100,
		MaxCapacity:    500,
	}

	// Initialize player state
	newPlayer := &models.Player{
		PlayerID:      playerID,
		PlayerName:    playerName,
		Stars:         0,
		StarsPerClick: 1,
		Jars:          []models.Jar{starterJar},
		// -- Boost mechanics (for future use) ---
		ClickMultiplier:       1.0,
		ClickBoostExpiresAt:   time.Time{},
		PassiveMultiplier:     1.0,
		PassiveBoostExpiresAt: time.Time{},
	}

	// Store the new player state
	playerStates[playerID] = newPlayer // This changes later when there is a DB

	return newPlayer.PlayerID
}

func (playerService *PlayerService) GetPlayer(playerID string) *models.Player {
	return playerStates[playerID]
}

func (playerService *PlayerService) GetAllPlayers() []*models.Player {
	players := make([]*models.Player, 0, len(playerStates))
	for _, player := range playerStates {
		players = append(players, player)
	}
	return players
}

func (playerService *PlayerService) DeletePlayer(playerID string) {
	// Delete the player state
	delete(playerStates, playerID)

}

func (playerService *PlayerService) UpdatePlayerName(playerID string, playerName string) {
	playerStates[playerID].PlayerName = playerName
}
