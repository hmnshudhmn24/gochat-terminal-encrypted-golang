package main

import (
	"bufio"
	"crypto/rsa"
	"fmt"
	"net"
	"os"
)

func startServer() {
	privateKey, publicKey := generateKeyPair()

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	fmt.Println("🔐 Waiting for connection on port 9090...")
	conn, _ := listener.Accept()
	defer conn.Close()

	sendPublicKey(conn, publicKey)
	clientPubKey := receivePublicKey(conn)

	go receiveMessages(conn, privateKey)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		encrypted := encryptMessage(text, clientPubKey)
		conn.Write(encrypted)
	}
}