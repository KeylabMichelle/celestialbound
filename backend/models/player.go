/*
Player state and logic
This file contains the Player struct and its methods for managing player state and actions.
*/

package models

import "time"

type Player struct {
	// --- Player information ---
	PlayerID   string `json:"player_id"`   // Unique identifier for the player
	PlayerName string `json:"player_name"` // Name of the player

	// --- Game state ---
	Stars         int   `json:"stars"`           // Core currency
	Jars          []Jar `json:"jars"`            // List of jars owned by the player
	StarsPerClick int   `json:"stars_per_click"` // Stars to add per click (manual action)

	// --- Boost mechanics (for future use) ---
	ClickMultiplier     float64   `json:"click_multiplier"`       // Multiplier for manual clicks
	ClickBoostExpiresAt time.Time `json:"click_boost_expires_at"` // Expiration time for click boost

	PassiveMultiplier     float64   `json:"passive_multiplier"`       // Multiplier for passive stars to all jars
	PassiveBoostExpiresAt time.Time `json:"passive_boost_expires_at"` // Expiration time for passive boost

}
