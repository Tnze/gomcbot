package gomcbot

/*
	Here is all events you can regist.
	When event happens, match signal will be senting to Game.Event
*/
const (
	DisconnectEvent = iota
	PlayerSpawnEvent
)

// GetEvents returns a int type channal.
// When event happends, a event ID was be sended into this chan
// Note that HandleGame will block if you don't recive from Events
func (g *Game) GetEvents() <-chan int {
	return g.events
}
