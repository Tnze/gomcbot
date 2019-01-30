package gomcbot

import (
	"fmt"
	//"github.com/mattn/go-colorable"//On Windows
	"testing"
)

func TestChatMsgToString(t *testing.T) {
	s := ` {"extra":[{"color":"green","text":"故我依然"},{"color":"white","text":"™ "},{"color":"gray","text":"Kun_QwQ"},{"color":"white","text":": 为什么想要用炼药锅灭火时总是跳不进去"}],"text":""}`
	cm, err := newChatMsg(s)
	if err != nil {
		t.Error(err)
	}
	if cm.ToString() != "故我依然™ Kun_QwQ: 为什么想要用炼药锅灭火时总是跳不进去" {
		t.Error("Should be 故我依然™ Kun_QwQ: 为什么想要用炼药锅灭火时总是跳不进去")
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

	s := ` {"extra":[{"color":"green","text":"故我依然"},{"color":"white","text":"™ "},{"color":"gray","text":"Kun_QwQ"},{"color":"white","text":": 为什么想要用炼药锅灭火时总是跳不进去"}],"text":""}`
	cm, err := newChatMsg(s)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(cm)
}
