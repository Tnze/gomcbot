package gomcbot

import (
	"./authenticate"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type encryptionRequest struct {
	ServerID          string
	PublicKeyLength   int
	PublicKey         []byte
	VerifyTokenLength int
	VerifyToken       []byte
}

func unpackEncryptionRequest(p Packet) (er encryptionRequest) {
	i := 0
	serverID, length := unpackString(p.Data[i:])
	i += length
	publicKeyLength, length := unpackVarInt(p.Data[i:])
	i += length
	publicKey := p.Data[i : i+publicKeyLength]
	i += publicKeyLength
	verifyTokenLength, length := unpackVarInt(p.Data[i:])
	i += length
	verifyToken := p.Data[i : i+verifyTokenLength]
	return encryptionRequest{
		ServerID:          serverID,
		PublicKeyLength:   publicKeyLength,
		PublicKey:         publicKey,
		VerifyTokenLength: verifyTokenLength,
		VerifyToken:       verifyToken,
	}
}

func unpackLoginSuccess(p Packet) (uuid string, userName string) {
	len := 0
	uuid, len = unpackString(p.Data[len:])
	userName, _ = unpackString(p.Data[len:])
	return
}

// AuthDigest computes a special SHA-1 digest required for Minecraft web
// authentication on Premium servers (online-mode=true).
// Source: http://wiki.vg/Protocol_Encryption#Server
//
// Also many, many thanks to SirCmpwn and his wonderful gist (C#):
// https://gist.github.com/SirCmpwn/404223052379e82f91e6
func AuthDigest(serverID, sharedSecret, publicKey string) string {
	h := sha1.New()
	io.WriteString(h, serverID)
	io.WriteString(h, sharedSecret)
	io.WriteString(h, publicKey)
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

func loginAuth(auth authenticate.Response, er encryptionRequest) error {
	digest := AuthDigest(er.ServerID, string(er.VerifyToken), string(er.PublicKey))
	resp, err := http.Post("https://sessionserver.mojang.com/session/minecraft/join",
		"application/json",
		strings.NewReader(fmt.Sprintf(`{"accessToken": %q,"selectedProfile": %q,"serverId": %q}`,
			auth.AccessToken, auth.SelectedProfile.ID, digest)),
	)
	if err != nil {
		return fmt.Errorf("post fail: %v", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(resp)
	if resp.Status != "204 No Content" {
		return fmt.Errorf("auth fail: %s", string(body))
	}
	return nil
}

// AES/CFB8 with random key
func newSymmetricEncryption() (key []byte, esteam, dsteam cipher.Stream) {
	key = make([]byte, 16)
	rand.Read(key) //生成密钥

	b, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	esteam = cipher.NewCFBDecrypter(b, key)
	dsteam = cipher.NewCFBDecrypter(b, key)
	return
}

// 1024-bit RSA
func genEncryptionKeyResponse(shareSecret, publicKey, verifyToken []byte) (erp *Packet, err error) {
	pk, err := x509.ParsePKCS1PublicKey(publicKey)
	if err != nil {
		err = fmt.Errorf("decode public key fail: %v", err)
		return
	}

	cryptPK, err := rsa.EncryptPKCS1v15(rand.Reader, pk, shareSecret)
	if err != nil {
		err = fmt.Errorf("encryption share secret fail: %v", err)
		return
	}
	verifyT, err := rsa.EncryptPKCS1v15(rand.Reader, pk, verifyToken)
	if err != nil {
		err = fmt.Errorf("encryption verfy tokenfail: %v", err)
		return
	}

	var data []byte
	data = append(data, packVarInt(len(cryptPK))...)
	data = append(data, cryptPK...)
	data = append(data, packVarInt(len(verifyT))...)
	data = append(data, verifyT...)
	erp = &Packet{
		ID:   0x01,
		Data: data,
	}
	return
}
