package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"fmt"
	"io"
	"net"
)

func generateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	return privateKey, &privateKey.PublicKey
}

func sendPublicKey(conn net.Conn, pub *rsa.PublicKey) {
	encoder := gob.NewEncoder(conn)
	encoder.Encode(pub)
}

func receivePublicKey(conn net.Conn) *rsa.PublicKey {
	var pub rsa.PublicKey
	decoder := gob.NewDecoder(conn)
	decoder.Decode(&pub)
	return &pub
}

func encryptMessage(msg string, pub *rsa.PublicKey) []byte {
	ciphertext, _ := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(msg))
	return ciphertext
}

func decryptMessage(ciphertext []byte, priv *rsa.PrivateKey) string {
	plaintext, _ := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	return string(plaintext)
}

func receiveMessages(conn net.Conn, priv *rsa.PrivateKey) {
	buf := make([]byte, 256)
	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("🔌 Connection closed.")
			return
		}
		decrypted := decryptMessage(buf[:n], priv)
		fmt.Println("Friend:", decrypted)
	}
}