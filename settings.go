package gomcbot

import pk "github.com/Tnze/gomcbot/packet"

// Settings 客户端设置
type Settings struct {
	Locale             string //地区
	ViewDistance       int    //视距
	ChatMode           int    //聊天模式
	ChatColors         bool   //聊天颜色
	DisplayedSkinParts uint8  //皮肤显示
	MainHand           int    //主手
	ReciveMap          bool   //接收地图数据
}

/*
	Used by Settings.DisplayedSkinParts.
	For each bits set if shows match part.
*/
const (
	_ = 1 << iota
	Jacket
	LeftSleeve
	RightSleeve
	LeftPantsLeg
	RightPantsLeg
	Hat
)

//DefaultSettings 定义了客户端的默认设置
var DefaultSettings = Settings{
	Locale:             "zh_CN",
	ViewDistance:       15,
	ChatMode:           0,
	DisplayedSkinParts: Jacket | LeftSleeve | RightSleeve | LeftPantsLeg | RightPantsLeg | Hat,
	MainHand:           1,
	ReciveMap:          true,
}

func (s *Settings) pack() (p *pk.Packet) {
	p = new(pk.Packet)
	p.ID = 0x04
	p.Data = append(p.Data, pk.PackString(s.Locale)...)
	p.Data = append(p.Data, byte(s.ViewDistance))
	p.Data = append(p.Data, pk.PackVarInt(int32(s.ChatMode))...)
	p.Data = append(p.Data, pk.PackBoolean(s.ChatColors), byte(s.DisplayedSkinParts))
	p.Data = append(p.Data, pk.PackVarInt(int32(s.MainHand))...)
	return
}

// Settings set a Settings to Game
func (g *Game) Settings(set Settings) {
	g.settings = set
}
