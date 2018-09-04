package common

const (
	// BACKGROUND needs to be drawn
	// beneath everything in the game
	BACKGROUND = -1
	// OBJECTS are small items and
	// around the world objects like
	// chests or bags of loot.
	OBJECTS = 0
	// SERVERPLAYER needs to be
	// displayed over the objects, but
	// still underneath the player.
	SERVERPLAYER = 1
	// PLAYER needs to be drawn
	// above everything in the game
	PLAYER = 2
)