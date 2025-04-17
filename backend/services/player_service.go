package services

import (
	"errors"
	"time"

	"github.com/keylab/celestialbound/backend/models"
	"github.com/keylab/celestialbound/backend/utils"
)

// In-memory storage for player states
var playerStates = make(map[string]*models.Player)

type PlayerService struct{}

func (playerService *PlayerService) CreatePlayer(playerName string) (string, error) {

	if playerName == "" {
		return "", errors.New("player name cannot be empty")
	}

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
	playerStates[playerID] = newPlayer

	return newPlayer.PlayerID, nil
}

func (playerService *PlayerService) GetPlayer(playerID string) (*models.Player, error) {

	/* Error handling */

	// Check if playerID is empty
	if playerID == "" {
		return nil, errors.New("playerID cannot be empty")
	}

	// Check if playerID exists in the map
	if playerStates[playerID] == nil {
		return nil, errors.New("player not found")
	}

	return playerStates[playerID], nil
}

func (playerService *PlayerService) GetAllPlayers() ([]*models.Player, error) {
	if len(playerStates) == 0 {
		return nil, errors.New("no players found")
	}

	players := make([]*models.Player, 0, len(playerStates))

	for _, player := range playerStates {
		players = append(players, player)
	}
	return players, nil
}

func (playerService *PlayerService) DeletePlayer(playerID string) error {
	if playerID == "" {
		return errors.New("playerID cannot be empty")
	}

	// Check if playerID exists in the map
	if playerStates[playerID] == nil {
		return errors.New("player not found")
	}

	// Delete the player state
	delete(playerStates, playerID)

	return nil

}

func (playerService *PlayerService) UpdatePlayerName(playerID string, playerName string) error {

	/* Error handling */
	if playerID == "" {
		return errors.New("playerID cannot be empty")
	}

	// Check if playerID exists in the map
	if playerStates[playerID] == nil {
		return errors.New("player not found")
	}

	playerStates[playerID].PlayerName = playerName

	return nil
}
