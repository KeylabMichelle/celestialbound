/* Jar logic */

package services

import (
	"errors" // for HTTP status codes
	"time"   //

	"github.com/keylab/celestialbound/backend/models"
	"github.com/keylab/celestialbound/backend/utils"
)

type JarService struct{}

func (jarService *JarService) CreateNewJar(playerID string) error {

	// Check if playerID exists in the map
	if playerStates[playerID] == nil {
		return errors.New("player not found")
	}

	// Create a new jar with default values
	newJar := models.Jar{
		JarID:          utils.GenerateUniqueID(),
		JarLevel:       1,
		StarsStored:    0,
		StarsPerSecond: 2,
		UpgradeCost:    100,
		MaxCapacity:    500,
		// -- Boost mechanics (for future use) ---
		PassiveMultiplier:     1.0,
		PassiveBoostExpiresAt: time.Time{},
	}

	// Add the new jar to the player's jars
	playerState := playerStates[playerID]
	playerState.Jars = append(playerState.Jars, newJar)

	return nil
}

func (JarService *JarService) GetAllJars(playerID string) ([]models.Jar, error) {

	// Check if playerID exists in the map
	if playerStates[playerID] == nil {
		return nil, errors.New("player not found")
	}
	// Return the jars of the player
	return playerStates[playerID].Jars, nil
}

func (JarService *JarService) ClickJar(playerID string, jarID string) (*models.Jar, error) {
	player := playerStates[playerID]
	if player == nil {
		return nil, errors.New("player not found")
	}

	// Find the target jar by ID
	var targetJar *models.Jar
	for i := range player.Jars {
		if player.Jars[i].JarID == jarID {
			targetJar = &player.Jars[i]
			break
		}
	}

	// Add stars to the jar
	if targetJar.StarsStored+player.StarsPerClick <= targetJar.MaxCapacity {
		targetJar.StarsStored += player.StarsPerClick
	} else {
		targetJar.StarsStored = targetJar.MaxCapacity
	}

	return targetJar, nil
}
