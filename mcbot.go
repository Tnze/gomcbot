package gomcbot

import (
	// "bufio"
	// "bytes"
	"fmt"
	// "net"

	"github.com/Tnze/gomcbot/network"
	pk "github.com/Tnze/gomcbot/network/packet"
)

//ProtocalVersion is the protocal version
// 477 for 1.14
const ProtocalVersion = 477

// PingAndList chack server status and list online player
// Return a JSON string about server status.
// see JSON format at https://wiki.vg/Server_List_Ping#Response
func PingAndList(addr string, port int) (string, error) {
	conn, err := network.DialMC(fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return "", err
	}

	//握手
	err = conn.WritePacket(pk.Marshal(
		0x00,                       //Handshake packet ID
		pk.VarInt(ProtocalVersion), //Protocal version
		pk.String(addr),            //Server's address
		pk.UnsignedShort(port),
		pk.Byte(1),
	))
	if err != nil {
		return "", fmt.Errorf("send handshake packect fail: %v", err)
	}

	//请求服务器状态
	err = conn.WritePacket(pk.Marshal(0))
	if err != nil {
		return "", fmt.Errorf("send list packect fail: %v", err)
	}

	//服务器返回状态
	recv, err := conn.ReadPacket()
	if err != nil {
		return "", fmt.Errorf("recv list packect fail: %v", err)
	}
	var s pk.String
	err = recv.Scan(&s)
	return string(s), err
}

// JoinServer connect a Minecraft server for playing the game.
// func JoinServer(addr string, port int, auth *Auth) (g *Client, err error) {
// 	//连接
// 	g = new(Game)
// 	g.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", addr, port))
// 	if err != nil {
// 		err = fmt.Errorf("cannot connect the server %q: %v", addr, err)
// 		return
// 	}

// 	//init Game
// 	g.settings = DefaultSettings //默认设置
// 	g.reciver = bufio.NewReader(g.conn)
// 	g.sender = g.conn
// 	g.wd.Entities = make(map[int32]Entity)
// 	g.wd.chunks = make(map[chunkLoc]*Chunk)
// 	g.events = make(chan Event)
// 	g.motion = make(chan func())

// 	//握手
// 	hsPacket := newHandshakePacket(ProtocalVersion, addr, port, 2) //构造握手包
// 	err = g.sendPacket(hsPacket)                                   //握手
// 	if err != nil {
// 		err = fmt.Errorf("send handshake packect fail: %v", err)
// 		return
// 	}

// 	//登录
// 	lsPacket := newLoginStartPacket(p.Name)
// 	err = g.sendPacket(lsPacket) //LoginStart
// 	if err != nil {
// 		err = fmt.Errorf("send login start packect fail: %v", err)
// 		return
// 	}
// 	for {
// 		//Recive Packet
// 		var pack *pk.Packet
// 		pack, err = g.recvPacket()
// 		if err != nil {
// 			err = fmt.Errorf("recv packet at state Login fail: %v", err)
// 			return
// 		}

// 		//Handle Packet
// 		switch pack.ID {
// 		case 0x00: //Disconnect
// 			s, _ := pk.UnpackString(bytes.NewReader(pack.Data))
// 			err = fmt.Errorf("connect disconnected by server because: %s", s)
// 			return
// 		case 0x01: //Encryption Request
// 			handleEncryptionRequest(g, pack, p)
// 		case 0x02: //Login Success
// 			// uuid, l := pk.UnpackString(pack.Data)
// 			// name, _ := unpackString(pack.Data[l:])
// 			return //switches the connection state to PLAY.
// 		case 0x03: //Set Compression
// 			threshold, _ := pk.UnpackVarInt(bytes.NewReader(pack.Data))
// 			g.threshold = int(threshold)
// 		case 0x04: //Login Plugin Request
// 			fmt.Println("Waring Login Plugin Request")
// 		}
// 	}
// }
