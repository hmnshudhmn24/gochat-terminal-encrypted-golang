package main

import (
	"bufio"
	"crypto/rsa"
	"fmt"
	"net"
	"os"
)

func startClient() {
	privateKey, publicKey := generateKeyPair()

	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		fmt.Println("Unable to connect:", err)
		return
	}
	defer conn.Close()

	serverPubKey := receivePublicKey(conn)
	sendPublicKey(conn, publicKey)

	go receiveMessages(conn, privateKey)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		encrypted := encryptMessage(text, serverPubKey)
		conn.Write(encrypted)
	}
}