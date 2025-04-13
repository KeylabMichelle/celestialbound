package utils

// CalculateUpgradeCost calculates the cost for the next upgrade based on the current level.
func CalculateUpgradeCost(jarLevel int) int {
	// Base cost for level 0 is 100, and each level increases the cost by 10%
	baseCost := 100 // Base cost for level 0

	cost := float64(baseCost) * (1.1 * float64(jarLevel))

	return int(cost)
}
