package gomcbot

import (
	"github.com/Tnze/gomcbot/network"
)

// Clint is the Object used to access Minecraft server
type Clint struct {
	conn *network.Conn

	Info      PlayerInfo
	abilities PlayerAbilities
	settings  Settings
	player    Player
	wd        world //the map data

	events chan Event
	motion chan func() //used to submit a function and HandleGame do
}
