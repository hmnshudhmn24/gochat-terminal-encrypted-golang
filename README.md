# 💬 GOCHAT: TERMINAL-BASED ENCRYPTED CHAT APP

**GoChat** is a simple peer-to-peer terminal chat app built in Go with end-to-end encryption using RSA.

## 🔐 Features

- Peer-to-peer encrypted communication
- RSA key generation per session
- Terminal interface
- No third-party libraries

## 🚀 How It Works

1. RSA key generation on startup.
2. Exchange of public keys between peers.
3. Messages are encrypted with the recipient's public key and decrypted with the sender's private key.

## 🧪 Usage

### Installation

```bash
git clone https://github.com/yourusername/gochat-terminal-encrypted-golang
cd gochat-terminal-encrypted-golang
go build -o gochat
```

### Start Server

```bash
./gochat
# Choose option 1
```

### Start Client

```bash
./gochat
# Choose option 2
```

## 📄 License

MIT