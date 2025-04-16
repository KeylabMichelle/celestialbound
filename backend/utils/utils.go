package utils

import (
	googleuuid "github.com/google/uuid"
)

// CalculateUpgradeCost calculates the cost for the next upgrade based on the current level.
func CalculateUpgradeCost(jarLevel int) int {
	// Base cost for level 0 is 100, and each level increases the cost by 10%
	baseCost := 100 // Base cost for level 0
	cost := float64(baseCost) * (1.1 * float64(jarLevel))

	return int(cost)
}

func GenerateUniqueID() string {
	// Generate a new UUID
	id := googleuuid.New()

	// Convert the UUID to a string and return it
	return id.String()
}
