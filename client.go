package gomcbot

import (
	"github.com/Tnze/gomcbot/network"
)

// Client is the Object used to access Minecraft server
type Client struct {
	conn *network.Conn

	// Info      PlayerInfo
	// abilities PlayerAbilities
	// settings  Settings
	// player    Player
	// wd        world //the map data

}
