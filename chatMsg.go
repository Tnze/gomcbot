package gomcbot

import (
	"encoding/json"
	"fmt"
)

//ChatMsg is a message sent by other
type ChatMsg jsonChat

type jsonChat struct {
	Text string `json:"text"`

	Bold          bool   `json:"bold"`          //粗体
	Italic        bool   `json:"Italic"`        //斜体
	UnderLined    bool   `json:"underlined"`    //下划线
	StrikeThrough bool   `json:"strikethrough"` //删除线
	Obfuscated    bool   `json:"obfuscated"`    //随机
	Color         string `json:"color"`

	Translate string        `json:"translate"`
	With      []interface{} `json:"with"` // How can go handle an JSON array with Object and String?
	Extra     []jsonChat    `json:"extra"`
}

func newChatMsg(jsonMsg string) (ChatMsg, error) {
	var jc jsonChat
	err := json.Unmarshal([]byte(jsonMsg), &jc)
	return ChatMsg(jc), err
}

//ToString convert a ChatMsg to string without using format
func (c ChatMsg) ToString() (s string) {
	s += c.Text
	if c.Translate == "chat.type.text" {
		if msg, ok := c.With[1].(string); ok {
			jc, _ := toChatMsg(c.With[0])
			s += fmt.Sprintf("<%s> %s", jc.ToString(), msg)
		}
	}
	if c.Extra != nil {
		for i := range c.Extra {
			s += ChatMsg(c.Extra[i]).ToString()
		}
	}
	return
}

// String return the message with escape sequence for ansi color.
// On windows, you may want print this string using
// github.com/mattn/go-colorable.
func (c ChatMsg) String() (s string) {
	s += "\033["
	if c.Bold {
		s += "1;"
	}
	if c.Italic {
		s += "3;"
	}
	if c.UnderLined {
		s += "4;"
	}
	if c.StrikeThrough {
		s += "9;"
	}
	switch c.Color {
	case "black":
		s += "30;"
	case "dark_blue":
		s += "34;"
	case "dark_green":
		s += "32;"
	case "dark_aqua":
		s += "36;"
	case "dark_red":
		s += "31;"
	case "dark_purple":
		s += "35;"
	case "gold":
		s += "33;"
	case "gray":
		s += "37;"
	case "dark_gray":
		s += "90;"
	case "blue":
		s += "94;"
	case "green":
		s += "92;"
	case "aqua":
		s += "96;"
	case "red":
		s += "91;"
	case "light_purple":
		s += "95;"
	case "yellow":
		s += "93;"
	case "white":
		s += "97;"
	}
	s = s[:len(s)-1] + "m"

	s += c.Text
	if c.Translate == "chat.type.text" {
		if msg, ok := c.With[1].(string); ok {
			jc, _ := toChatMsg(c.With[0])
			s += fmt.Sprintf("<%s> %s", jc.ToString(), msg)
		}
	}

	s += "\033[0m"

	if c.Extra != nil {
		for i := range c.Extra {
			s += ChatMsg(c.Extra[i]).String()
		}
	}
	return
}

func toChatMsg(value interface{}) (c ChatMsg, err error) {
	var b []byte
	b, err = json.Marshal(value)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &c)
	return
}
