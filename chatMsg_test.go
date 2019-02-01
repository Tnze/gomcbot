package gomcbot

import (
	// "fmt"
	//"github.com/mattn/go-colorable"//On Windows need
	"testing"
)

var jsons = []string{
	`{"extra":[{"color":"green","text":"故我依然"},{"color":"white","text":"™ "},{"color":"gray","text":"Kun_QwQ"},{"color":"white","text":": 为什么想要用炼药锅灭火时总是跳不进去"}],"text":""}`,
	`{"translate":"chat.type.text","with":[{"insertion":"Xi_Xi_Mi","clickEvent":{"action":"suggest_command","value":"/tell Xi_Xi_Mi "},"hoverEvent":{"action":"show_entity","value":{"text":"{name:\"{\\\"text\\\":\\\"Xi_Xi_Mi\\\"}\",id:\"c1445a67-7551-4d7e-813d-65ef170ae51f\",type:\"minecraft:player\"}"}},"text":"Xi_Xi_Mi"},"好像是这个id。。"]}`,
}

var texts = []string{
	"\033[92m故我依然\033[0m\033[97m™ \033[0m\033[37mKun_QwQ\033[0m\033[97m: 为什么想要用炼药锅灭火时总是跳不进去\033[0m",
	"<Xi_Xi_Mi> 好像是这个id。。",
}

func TestChatMsgFormatString(t *testing.T) {
	for i, v := range jsons {
		cm, err := newChatMsg([]byte(v))
		if err != nil {
			t.Error(err)
		}
		if cm.String() != texts[i] {
			t.Errorf("%v Should be %s", cm, texts[i])
		}
	}
}
