package models

import "time"

type Jar struct {
	JarID          string `json:"jar_id"`           // Unique identifier for the jar
	JarLevel       int    `json:"jar_level"`        // Current level of the jar (0-10)
	StarsStored    int    `json:"stars_stored"`     // Stars stored in the jar
	StarsPerSecond int    `json:"stars_per_second"` // Stars gained per second (Passive action)
	UpgradeCost    int    `json:"upgrade_cost"`     // Cost for next upgrade of the jar
	MaxCapacity    int    `json:"max_capacity"`     // Maximum capacity of the jar
	// --- Boost mechanics (for future use) ---
	PassiveMultiplier     float64   `json:"passive_multiplier"`       // Multiplier for passive stars
	PassiveBoostExpiresAt time.Time `json:"passive_boost_expires_at"` // Expiration time for passive boost

}
