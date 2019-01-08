package gomcbot

import (
	"./CFB8"
	pk "./packet"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type encryptionRequest struct {
	ServerID    string
	PublicKey   []byte
	VerifyToken []byte
}

func unpackEncryptionRequest(p pk.Packet) (er encryptionRequest) {
	i := 0
	serverID, length := pk.UnpackString(p.Data[i:])
	i += length
	publicKeyLength, length := pk.UnpackVarInt(p.Data[i:])
	i += length
	publicKey := p.Data[i : i+int(publicKeyLength)]
	i += int(publicKeyLength)
	verifyTokenLength, length := pk.UnpackVarInt(p.Data[i:])
	i += length
	verifyToken := p.Data[i : i+int(verifyTokenLength)]
	return encryptionRequest{
		ServerID:    serverID,
		PublicKey:   publicKey,
		VerifyToken: verifyToken,
	}
}

// AuthDigest computes a special SHA-1 digest required for Minecraft web
// authentication on Premium servers (online-mode=true).
// Source: http://wiki.vg/Protocol_Encryption#Server
//
// Also many, many thanks to SirCmpwn and his wonderful gist (C#):
// https://gist.github.com/SirCmpwn/404223052379e82f91e6
func AuthDigest(serverID string, sharedSecret, publicKey []byte) string {
	h := sha1.New()
	h.Write([]byte(serverID))
	h.Write(sharedSecret)
	h.Write(publicKey)
	hash := h.Sum(nil)

	// Check for negative hashes
	negative := (hash[0] & 0x80) == 0x80
	if negative {
		hash = twosComplement(hash)
	}

	// Trim away zeroes
	res := strings.TrimLeft(fmt.Sprintf("%x", hash), "0")
	if negative {
		res = "-" + res
	}

	return res
}

// little endian
func twosComplement(p []byte) []byte {
	carry := true
	for i := len(p) - 1; i >= 0; i-- {
		p[i] = byte(^p[i])
		if carry {
			carry = p[i] == 0xff
			p[i]++
		}
	}
	return p
}

type request struct {
	AccessToken     string `json:"accessToken"`
	SelectedProfile string `json:"selectedProfile"`
	ServerID        string `json:"serverId"`
}

func loginAuth(AsTk, UUID string, er encryptionRequest) error {
	digest := AuthDigest(er.ServerID, er.VerifyToken, er.PublicKey)
	//Post
	client := http.Client{}
	requestPacket, err := json.Marshal(
		request{
			AccessToken:     AsTk,
			SelectedProfile: UUID,
			ServerID:        digest,
		},
	)
	if err != nil {
		return fmt.Errorf("create request packet to authenticate faile: %v", err)
	}
	PostRequest, err := http.NewRequest(http.MethodPost, "https://sessionserver.mojang.com/session/minecraft/join",
		bytes.NewReader(requestPacket))
	if err != nil {
		return fmt.Errorf("make request error: %v", err)
	}
	PostRequest.Header.Set("User-Agent", "gomcbot")
	PostRequest.Header.Set("Connection", "keep-alive")
	resp, err := client.Do(PostRequest)
	if err != nil {
		return fmt.Errorf("post fail: %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.Status != "204 No Content" {
		return fmt.Errorf("auth fail: %s", string(body))
	}
	return nil
}

// AES/CFB8 with random key
func newSymmetricEncryption() (key []byte, encoStream, decoStream cipher.Stream) {
	key = make([]byte, 16)
	rand.Read(key) //生成密钥

	b, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	decoStream = CFB8.NewCFB8Decrypt(b, key)
	encoStream = CFB8.NewCFB8Encrypt(b, key)
	return
}

// 1024-bit RSA
func genEncryptionKeyResponse(shareSecret, publicKey, verifyToken []byte) (erp *pk.Packet, err error) {

	iPK, err := x509.ParsePKIXPublicKey(publicKey) // Decode Public Key
	if err != nil {
		err = fmt.Errorf("decode public key fail: %v", err)
		return
	}
	rsaKey := iPK.(*rsa.PublicKey)
	cryptPK, err := rsa.EncryptPKCS1v15(rand.Reader, rsaKey, shareSecret)
	if err != nil {
		err = fmt.Errorf("encryption share secret fail: %v", err)
		return
	}
	verifyT, err := rsa.EncryptPKCS1v15(rand.Reader, rsaKey, verifyToken)
	if err != nil {
		err = fmt.Errorf("encryption verfy tokenfail: %v", err)
		return
	}
	var data []byte
	data = append(data, pk.PackVarInt(int32(len(cryptPK)))...)
	data = append(data, cryptPK...)
	data = append(data, pk.PackVarInt(int32(len(verifyT)))...)
	data = append(data, verifyT...)
	erp = &pk.Packet{
		ID:   0x01,
		Data: data,
	}
	return
}