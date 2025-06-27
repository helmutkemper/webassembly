package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdh"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"syscall/js"
	"time"
)

type Message struct {
	From      string `json:"from"`
	To        string `json:"to,omitempty"`
	Group     string `json:"group,omitempty"`
	Type      string `json:"type"`
	Payload   string `json:"payload"`
	Timestamp int64  `json:"timestamp"`
	MessageID string `json:"messageId"`
}

var (
	ws        js.Value
	sharedKey []byte
	onMessage js.Func
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// connect button handler
	js.Global().Get("document").Call("getElementById", "connectBtn").
		Call("addEventListener", "click", js.FuncOf(connect))
	select {}
}

func connect(this js.Value, args []js.Value) any {
	url := "ws://" + js.Global().Get("location").Get("host").String() + "/ws"
	ws = js.Global().Get("WebSocket").New(url)

	// generate ECDH keys
	curve := ecdh.P256()
	priv, _ := curve.GenerateKey(rand.Reader)
	pubBytes := priv.PublicKey().Bytes()

	// send key+username
	ws.Call("send", toJSON(Message{
		From:    js.Global().Get("document").Call("getElementById", "username").Get("value").String(),
		Type:    "key",
		Payload: base64.StdEncoding.EncodeToString(pubBytes),
	}))

	ws.Call("addEventListener", "message", js.FuncOf(onWSMessage(priv)))

	return nil
}

func onWSMessage(priv *ecdh.PrivateKey) func(this js.Value, args []js.Value) any {
	return func(this js.Value, args []js.Value) any {
		var msg Message
		json.Unmarshal([]byte(args[0].Get("data").String()), &msg)

		switch msg.Type {
		case "key":
			// server public
			serverPub, _ := base64.StdEncoding.DecodeString(msg.Payload)
			curve := ecdh.P256()
			pubKey, _ := curve.NewPublicKey(serverPub)
			shared, _ := priv.ECDH(pubKey)
			sharedKey = shared[:16]
			fmt.Println("Shared key established")
		default:
			// decrypt then dispatch
			plain, _ := decrypt(msg.Payload, sharedKey)
			json.Unmarshal(plain, &msg)
			fmt.Printf("[%s] %s\n", msg.From, msg.Payload)
		}
		return nil
	}
}

func SendMessage(to, group, typ, text string) {
	payload := text
	plain, _ := json.Marshal(Message{
		From:      "", // your username field
		To:        to,
		Group:     group,
		Type:      typ,
		Payload:   text,
		Timestamp: time.Now().UnixNano() / 1e6,
		MessageID: fmt.Sprintf("%d", rand.Int63()),
	})
	enc, _ := encrypt(plain, sharedKey)
	ws.Call("send", toJSON(Message{
		From:      "", // username
		To:        to,
		Group:     group,
		Type:      typ,
		Payload:   enc,
		Timestamp: time.Now().UnixNano() / 1e6,
		MessageID: fmt.Sprintf("%d", rand.Int63()),
	}))
}

func toJSON(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

// AES-GCM helpers

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
