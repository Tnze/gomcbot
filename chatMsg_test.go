package gomcbot

import (
	"fmt"
	//"github.com/mattn/go-colorable"//On Windows need
	"testing"
)

var jsons = []string{
	`{"extra":[{"color":"green","text":"故我依然"},{"color":"white","text":"™ "},{"color":"gray","text":"Kun_QwQ"},{"color":"white","text":": 为什么想要用炼药锅灭火时总是跳不进去"}],"text":""}`,
	`{"translate":"chat.type.text","with":[{"insertion":"Xi_Xi_Mi","clickEvent":{"action":"suggest_command","value":"/tell Xi_Xi_Mi "},"hoverEvent":{"action":"show_entity","value":{"text":"{name:\"{\\\"text\\\":\\\"Xi_Xi_Mi\\\"}\",id:\"c1445a67-7551-4d7e-813d-65ef170ae51f\",type:\"minecraft:player\"}"}},"text":"Xi_Xi_Mi"},"好像是这个id。。"]}`,
}

var texts = []string{
	"故我依然™ Kun_QwQ: 为什么想要用炼药锅灭火时总是跳不进去",
	"<Xi_Xi_Mi> 好像是这个id。。",
}

func TestChatMsgToString(t *testing.T) {
	for i, v := range jsons {
		cm, err := newChatMsg(v)
		if err != nil {
			t.Error(err)
		}
		if cm.ToString() != texts[i] {
			t.Error("Should be " + texts[i])
		}
	}

}

func TestChatMsgFormatString(t *testing.T) {
	// Windows:

	// stdout := colorable.NewColorableStdout()
	// s := ` {"extra":[{"color":"green","text":"故我依然"},{"color":"white","text":"™ "},{"color":"gray","text":"Kun_QwQ"},{"color":"white","text":": 为什么想要用炼药锅灭火时总是跳不进去"}],"text":""}`
	// cm, err := newChatMsg(s)
	// if err != nil {
	// 	t.Error(err)
	// }

	// fmt.Fprintln(stdout, cm)

	//Linux/OS X

	//s := ` {"extra":[{"color":"green","text":"故我依然"},{"color":"white","text":"™ "},{"color":"gray","text":"Kun_QwQ"},{"color":"white","text":": 为什么想要用炼药锅灭火时总是跳不进去"}],"text":""}`
	s := `{"translate":"chat.type.text","with":[{"insertion":"Xi_Xi_Mi","clickEvent":{"action":"suggest_command","value":"/tell Xi_Xi_Mi "},"hoverEvent":{"action":"show_entity","value":{"text":"{name:\"{\\\"text\\\":\\\"Xi_Xi_Mi\\\"}\",id:\"c1445a67-7551-4d7e-813d-65ef170ae51f\",type:\"minecraft:player\"}"}},"text":"Xi_Xi_Mi"},"好像是这个id。。"]}`
	cm, err := newChatMsg(s)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(cm)
}
