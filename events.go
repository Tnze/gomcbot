package gomcbot

//Event happends in game and you can recive it from what Game.GetEvent() returns
type Event int

/*
	Here is all events you can regist.
	When event happens, match signal will be senting to Game.Event
*/
const (
	DisconnectEvent Event = iota
	PlayerSpawnEvent
	PlayerDeadEvent
	InventoryChangeEvent
	BlockChangeEvent
)

// GetEvents returns a int type channal.
// When event happends, a event ID was be sended into this chan
// Note that HandleGame will block if you don't recive from Events
func (g *Game) GetEvents() <-chan Event {
	return g.events
}
