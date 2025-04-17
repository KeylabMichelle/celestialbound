#!/bin/bash

BASE_URL="http://localhost:8080"

echo "CelestialBound API test..."

# Create Player
echo "------------------------------------------------------------------"
echo "Creating player..."
RESPONSE=$(curl -s -X POST "$BASE_URL/player/" \
  -H "Content-Type: application/json" \
  -d '{"player_name": "TestPlayer"}')

PLAYER_ID=$(echo $RESPONSE | jq -r '.player_id')
PLAYER_NAME=$(echo $RESPONSE | jq -r '.player_name')
echo "Created Player ID: $PLAYER_ID and Name: $PLAYER_NAME"

echo "------------------------------------------------------------------"

# Get Player by ID
echo "------------------------------------------------------------------"
echo "Getting player info..."
curl -s "$BASE_URL/player/$PLAYER_ID" | jq
echo "------------------------------------------------------------------"


# Click to Add Stars
echo "------------------------------------------------------------------"
echo "Clicking for stars..."
curl -s -X POST "$BASE_URL/click/$PLAYER_ID" | jq
echo "------------------------------------------------------------------"


# Create a New Jar
echo "------------------------------------------------------------------"
echo "Creating a new jar..."
curl -s -X POST "$BASE_URL/player/$PLAYER_ID/jar/" | jq
echo "------------------------------------------------------------------"

# Get All Jars
echo "------------------------------------------------------------------"
echo "Getting all jars..."
curl -s "$BASE_URL/player/$PLAYER_ID/jar/" | jq
echo "------------------------------------------------------------------"

# Update Player Name and Stars
echo "------------------------------------------------------------------"
echo "Updating player info..."
curl -s -X PUT "$BASE_URL/player/$PLAYER_ID" \
  -H "Content-Type: application/json" \
  -d '{"player_name": "UpdatedName"}' | jq
echo "------------------------------------------------------------------"

# Get All Players
echo "------------------------------------------------------------------"
echo "Listing all players..."
curl -s "$BASE_URL/player/" | jq
echo "------------------------------------------------------------------"

# Delete Player
echo "------------------------------------------------------------------"
echo "Deleting player..."
curl -s -X DELETE "$BASE_URL/player/$PLAYER_ID" | jq
echo "------------------------------------------------------------------"

# Get All Players to make sure it was deleted
echo "------------------------------------------------------------------"
echo "Listing all players..."
curl -s "$BASE_URL/player/" | jq
echo "------------------------------------------------------------------"

echo "✅ Test complete."
