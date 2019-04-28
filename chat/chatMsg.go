package gomcbot

// import (
// 	"encoding/json"
// 	"fmt"

// 	"github.com/Tnze/gomcbot/data"
// )

// //ChatMsg is a message sent by other
// type ChatMsg jsonChat

// type jsonChat struct {
// 	Text string `json:"text"`

// 	Bold          bool   `json:"bold"`          //粗体
// 	Italic        bool   `json:"Italic"`        //斜体
// 	UnderLined    bool   `json:"underlined"`    //下划线
// 	StrikeThrough bool   `json:"strikethrough"` //删除线
// 	Obfuscated    bool   `json:"obfuscated"`    //随机
// 	Color         string `json:"color"`

// 	Translate string            `json:"translate"`
// 	With      []json.RawMessage `json:"with"` // How can go handle an JSON array with Object and String?
// 	Extra     []jsonChat        `json:"extra"`
// }

// func newChatMsg(jsonMsg []byte) (jc ChatMsg, err error) {
// 	if jsonMsg[0] == '"' {
// 		err = json.Unmarshal(jsonMsg, &jc.Text)
// 	} else {
// 		err = json.Unmarshal(jsonMsg, &jc)
// 	}
// 	return
// }

// var colors = map[string]int{
// 	"black":        30,
// 	"dark_blue":    34,
// 	"dark_green":   32,
// 	"dark_aqua":    36,
// 	"dark_red":     31,
// 	"dark_purple":  35,
// 	"gold":         33,
// 	"gray":         37,
// 	"dark_gray":    90,
// 	"blue":         94,
// 	"green":        92,
// 	"aqua":         96,
// 	"red":          91,
// 	"light_purple": 95,
// 	"yellow":       93,
// 	"white":        97,
// }

// // String return the message with escape sequence for ansi color.
// // On windows, you may want print this string using
// // github.com/mattn/go-colorable.
// func (c ChatMsg) String() (s string) {
// 	var format string
// 	if c.Bold {
// 		format += "1;"
// 	}
// 	if c.Italic {
// 		format += "3;"
// 	}
// 	if c.UnderLined {
// 		format += "4;"
// 	}
// 	if c.StrikeThrough {
// 		format += "9;"
// 	}
// 	if c.Color != "" {
// 		format += fmt.Sprintf("%d;", colors[c.Color])
// 	}

// 	if format != "" {
// 		s = "\033[" + format[:len(format)-1] + "m"
// 	}

// 	s += c.Text

// 	if format != "" {
// 		s += "\033[0m"
// 	}

// 	//handle translate
// 	if c.Translate != "" {
// 		args := make([]interface{}, len(c.With))
// 		for i, v := range c.With {
// 			args[i], _ = newChatMsg(v) //ignore error
// 		}

// 		s += fmt.Sprintf(data.enUs[c.Translate], args...)
// 	}

// 	if c.Extra != nil {
// 		for i := range c.Extra {
// 			s += ChatMsg(c.Extra[i]).String()
// 		}
// 	}
// 	return
// }
