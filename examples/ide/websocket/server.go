package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdh"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

type Message struct {
	From      string `json:"from"`
	To        string `json:"to,omitempty"`
	Group     string `json:"group,omitempty"`
	Type      string `json:"type"` // "private","group","broadcast","key"
	Payload   string `json:"payload"`
	Timestamp int64  `json:"timestamp"`
	MessageID string `json:"messageId"`
}

type Client struct {
	User      string
	Conn      *websocket.Conn
	SharedKey []byte       // AES key derived via ECDH
	Send      chan Message // outbound messages
}

var (
	clients   = make(map[string]*Client)
	groups    = make(map[string]map[string]bool)
	clientsMu sync.RWMutex
)

func main() {
	http.HandleFunc("/ws", wsHandler)
	fmt.Println("WebSocket server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// upgrade to WS
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	// first message must be key-exchange with user name
	var initMsg Message
	if err := conn.ReadJSON(&initMsg); err != nil || initMsg.Type != "key" {
		conn.Close()
		return
	}

	// ECDH key exchange
	curve := ecdh.P256()
	serverPriv, _ := curve.GenerateKey(rand.Reader)
	clientPubBytes, _ := base64.StdEncoding.DecodeString(initMsg.Payload)
	clientPubKey, _ := curve.NewPublicKey(clientPubBytes)
	shared, _ := serverPriv.ECDH(clientPubKey)

	// register client
	clientsMu.Lock()
	clients[initMsg.From] = &Client{
		User:      initMsg.From,
		Conn:      conn,
		SharedKey: shared[:16], // use first 16 bytes for AES-128
		Send:      make(chan Message, 16),
	}
	clientsMu.Unlock()

	// send server's public key back
	serverPubBytes := serverPriv.PublicKey().Bytes()
	conn.WriteJSON(Message{
		From:    "SERVER",
		Type:    "key",
		Payload: base64.StdEncoding.EncodeToString(serverPubBytes),
	})

	// launch reader and writer
	go clientWriter(clients[initMsg.From])
	clientReader(clients[initMsg.From])
}

// read from client
func clientReader(c *Client) {
	defer unregister(c)
	for {
		var encMsg Message
		if err := c.Conn.ReadJSON(&encMsg); err != nil {
			return
		}
		// decrypt payload
		data, err := decrypt(encMsg.Payload, c.SharedKey)
		if err != nil {
			continue
		}
		var msg Message
		if err := json.Unmarshal(data, &msg); err != nil {
			continue
		}
		routeMessage(msg)
	}
}

// write to client
func clientWriter(c *Client) {
	for msg := range c.Send {
		// encrypt payload
		plain, _ := json.Marshal(msg)
		enc, _ := encrypt(plain, c.SharedKey)
		c.Conn.WriteJSON(Message{
			From:      msg.From,
			To:        msg.To,
			Group:     msg.Group,
			Type:      msg.Type,
			Payload:   enc,
			Timestamp: msg.Timestamp,
			MessageID: msg.MessageID,
		})
	}
}

func unregister(c *Client) {
	clientsMu.Lock()
	delete(clients, c.User)
	for _, members := range groups {
		delete(members, c.User)
	}
	close(c.Send)
	clientsMu.Unlock()
}

func routeMessage(msg Message) {
	clientsMu.RLock()
	defer clientsMu.RUnlock()

	switch msg.Type {
	case "private":
		if dst, ok := clients[msg.To]; ok {
			dst.Send <- msg
		}
	case "group":
		if members, ok := groups[msg.Group]; ok {
			for user := range members {
				if user != msg.From {
					clients[user].Send <- msg
				}
			}
		}
	case "broadcast":
		for _, cl := range clients {
			if cl.User != msg.From {
				cl.Send <- msg
			}
		}
	case "join":
		// Payload is group name
		g := msg.Payload
		if groups[g] == nil {
			groups[g] = make(map[string]bool)
		}
		groups[g][msg.From] = true
	case "leave":
		g := msg.Payload
		delete(groups[g], msg.From)
	}
}

// AES-GCM encrypt/decrypt

func encrypt(plain, key []byte) (string, error) {
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	rand.Read(nonce)
	ct := gcm.Seal(nonce, nonce, plain, nil)
	return base64.StdEncoding.EncodeToString(ct), nil
}

func decrypt(enc string, key []byte) ([]byte, error) {
	data, _ := base64.StdEncoding.DecodeString(enc)
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce, ct := data[:nonceSize], data[nonceSize:]
	return gcm.Open(nil, nonce, ct, nil)
}
