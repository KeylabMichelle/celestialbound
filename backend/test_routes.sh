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
echo "Created Player ID: $PLAYER_ID"

echo "------------------------------------------------------------------"

# Get Player by ID
echo "Getting player info..."
curl -s "$BASE_URL/player/$PLAYER_ID" | jq
echo "------------------------------------------------------------------"

# Create a New Jar
echo "Creating a new jar..."
curl -s -X POST "$BASE_URL/player/$PLAYER_ID/jar/" | jq
echo "------------------------------------------------------------------"

# Get All Jars
echo "Getting all jars..."
JARS=$(curl -s "$BASE_URL/player/$PLAYER_ID/jar/")
echo "$JARS" | jq
echo "------------------------------------------------------------------"

# Extract first JarID
JAR_ID=$(echo "$JARS" | jq -r '.[0].jar_id')
echo "Target Jar ID: $JAR_ID"

# Click on the specific Jar
echo "Clicking on jar to generate stars..."
RESPONSE=$(curl -s -X POST "$BASE_URL/player/$PLAYER_ID/jar/$JAR_ID/click")
STARS=$(echo "$RESPONSE" | jq -r '.stars_stored')
echo "Stars stored in jar after click: $STARS"
echo "------------------------------------------------------------------"

# Get All Jars
echo "Getting all jars..."
JARS=$(curl -s "$BASE_URL/player/$PLAYER_ID/jar/")
echo "$JARS" | jq
echo "------------------------------------------------------------------"


# Update Player Name
echo "Updating player name..."
curl -s -X PUT "$BASE_URL/player/$PLAYER_ID" \
  -H "Content-Type: application/json" \
  -d '{"player_name": "UpdatedName"}' | jq
echo "------------------------------------------------------------------"

# List All Players
echo "Listing all players..."
curl -s "$BASE_URL/player/" | jq
echo "------------------------------------------------------------------"

# Delete Player
echo "Deleting player..."
curl -s -X DELETE "$BASE_URL/player/$PLAYER_ID" | jq
echo "------------------------------------------------------------------"

# Confirm Deletion
echo "Listing all players to confirm deletion..."
curl -s "$BASE_URL/player/" | jq
echo "------------------------------------------------------------------"

echo "Test complete."
