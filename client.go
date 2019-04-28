package gomcbot

import (
	"github.com/Tnze/gomcbot/network"
)

// Client is the Object used to access Minecraft server
type Client struct {
	conn *network.Conn
	Auth

	// Info      PlayerInfo
	// abilities PlayerAbilities
	settings Settings
	// player    Player
	// wd        world //the map data

}

//NewClient init and return a new Client
func NewClient() (c *Client) {
	c = new(Client)

	//init Client
	c.settings = DefaultSettings
	c.Name = "Steve"

	return
}
