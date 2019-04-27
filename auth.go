package gomcbot

// Auth includes a account
type Auth struct {
	Name string
	UUID string
	AsTk string
}

// 加密请求
func handleEncryptionRequest(g *Game, pack *pk.Packet, auth *Auth) error {
	//创建AES对称加密密钥
	key, encoStream, decoStream := newSymmetricEncryption()

	//解析EncryptionRequest包
	er, err := unpackEncryptionRequest(*pack)
	if err != nil {
		return err
	}
	err = loginAuth(auth.AsTk, auth.Name, auth.UUID, key, *er) //向Mojang验证
	if err != nil {
		return fmt.Errorf("login fail: %v", err)
	}

	// 响应加密请求
	var p *pk.Packet // Encryption Key Response
	p, err = genEncryptionKeyResponse(key, er.PublicKey, er.VerifyToken)
	if err != nil {
		return fmt.Errorf("gen encryption key response fail: %v", err)
	}
	err = g.sendPacket(p)
	if err != nil {
		return err
	}

	//加密连接
	g.reciver = bufio.NewReader(cipher.StreamReader{ //Set reciver for AES
		S: decoStream,
		R: g.conn,
	})
	g.sender = cipher.StreamWriter{
		S: encoStream,
		W: g.conn,
	}
	return nil
}
