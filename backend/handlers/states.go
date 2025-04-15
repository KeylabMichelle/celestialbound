package handlers

import (
	"github.com/keylab/celestialbound/backend/models"
)

// In-memory storage for player states
var playerStates = make(map[string]*models.PlayerState)
