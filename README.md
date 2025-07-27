# 💬 GOCHAT: TERMINAL-BASED ENCRYPTED CHAT APP

**GoChat** is a secure, peer-to-peer chat application built in Go for terminal users. It uses **RSA** for public-private key exchange and **AES** for encrypting the chat messages, ensuring **end-to-end encryption** during communication.

This project is a great demonstration of:
- 🔐 Cryptography in Go
- 📡 Peer-to-peer (P2P) communication
- ⚙️ Networking and concurrency using `net` and `goroutines`



## 🚀 Features

- 🔑 RSA public/private key generation and exchange
- 🔒 AES encrypted chat messages
- 📥 Real-time bidirectional messaging
- 🌐 Works on localhost or over LAN (custom IP/port)
- 🧵 Concurrent send/receive using Goroutines
- 🧪 Simple and minimal CLI interface
- 🔧 No external dependencies — only uses Go’s standard library



## 🧱 Tech Stack

| Component        | Usage                                   |
|------------------|------------------------------------------|
| Language         | Go (Golang)                             |
| Networking       | `net` package                           |
| Cryptography     | `crypto/rsa`, `crypto/aes`, `crypto/rand` |
| Concurrency      | `goroutines`, `channels`                |
| Storage          | N/A (no persistence, real-time only)    |



## 📦 Project Structure

```
gochat-terminal-encrypted-golang/
├── main.go             # Entry point and app mode selection (server/client)
├── server.go           # Server-side networking and encryption handling
├── client.go           # Client-side networking and encryption handling
├── crypto.go           # RSA keygen, AES encryption/decryption logic
├── utils.go            # Helper functions (I/O, etc.)
├── README.md           # Project documentation
```



## 🖥️ Usage

### 1. 📥 Clone and Build

```bash
git clone https://github.com/yourusername/gochat-terminal-encrypted-golang.git
cd gochat-terminal-encrypted-golang
go build -o gochat
```

### 2. 🖥️ Run as Server

```bash
./gochat server --port 9000
```

### 3. 💻 Run as Client

```bash
./gochat client --ip 127.0.0.1 --port 9000
```

> Server and client will automatically exchange RSA keys before the chat begins.



## 🔐 How Encryption Works

1. **RSA key pair** is generated at both ends (client/server).
2. Each peer **exchanges public keys** over the network.
3. A random **AES session key** is generated and encrypted using the peer's RSA public key.
4. All chat messages are **AES-encrypted** with the shared session key.
5. Messages are decrypted at the other end using the session key.



## 💬 Example Session

### Server:

```bash
[+] Listening on port 9000...
[✓] Received client's RSA public key
[✓] Encrypted AES key sent
[🔐] Secure chat started!
You: Hello from server!
```

### Client:

```bash
[✓] Connected to 127.0.0.1:9000
[✓] Sent RSA public key
[✓] Received and decrypted AES key
[🔐] Secure chat started!
You: Hello from client!
```


## 🧪 Sample Commands

```bash
./gochat server --port 9000
./gochat client --ip 192.168.0.105 --port 9000
```



## 📌 Future Plans

- 🌍 Add support for NAT traversal (for remote P2P)
- 💾 Optional chat logging (encrypted)
- 👥 Group chat or multi-user support
- 🔒 Add digital signatures for message authenticity
- 📦 Docker support for easy deployment



## 🛠️ Requirements

- Go 1.18 or later  
- Terminal access to test locally or over LAN
