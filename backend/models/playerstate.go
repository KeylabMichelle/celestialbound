package models

type PlayerState struct {
	PlayerStateID string `json:"player_state_id"` // Unique identifier for the player state
	PlayerID      string `json:"player_id"`       // Unique identifier for the player
	JarID         string `json:"jar_id"`          // Unique identifier for the jar
}
