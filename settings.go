package gomcbot

import pk "./packet"

// Settings 客户端设置
type Settings struct {
	Locale             string //地区
	ViewDistance       int    //视距
	ChatMode           int    //聊天模式
	ChatColors         bool   //聊天颜色
	DisplayedSkinParts uint8  //皮肤显示
	MainHand           int    //主手
}

// Cape enabled
const (
	_ = 1 << iota
	// Jacket 衣服
	Jacket
	// LeftSleeve 左袖子
	LeftSleeve
	// RightSleeve 右袖子
	RightSleeve
	// LeftPantsLeg 左裤子
	LeftPantsLeg
	// RightPantsLeg 右裤子
	RightPantsLeg
	// Hat 帽子
	Hat
)

//DefaultSettings 默认设置
var DefaultSettings = Settings{
	Locale:             "zh_CN",
	ViewDistance:       15,
	ChatMode:           0,
	DisplayedSkinParts: Jacket | LeftSleeve | RightSleeve | LeftPantsLeg | RightPantsLeg | Hat,
	MainHand:           1,
}

func (s *Settings) pack() (p []byte) {
	p = append(p, pk.PackString(s.Locale)...)
	p = append(p, byte(s.ViewDistance))
	p = append(p, pk.PackVarInt(int32(s.ChatMode))...)
	p = append(p, pk.PackBoolean(s.ChatColors), byte(s.DisplayedSkinParts))
	p = append(p, pk.PackVarInt(int32(s.MainHand))...)
	return
}
