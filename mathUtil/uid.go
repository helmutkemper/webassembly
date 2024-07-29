package mathUtil

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

func GetUID() (UID string) {
	// Obtenha a data e hora atual no formato UNIX timestamp
	timestamp := time.Now().UnixNano()

	// Converta o timestamp para uma string hexadecimal
	timestampHex := fmt.Sprintf("%x", timestamp)

	// Gere bytes aleatórios
	bytes := make([]byte, 16) // 16 bytes = 128 bits
	if _, err := rand.Read(bytes); err != nil {
		log.Printf("mathUtil.GetUID().error: %v", err)
		return ""
	}

	// Combine o timestamp hexadecimal com os bytes aleatórios
	randomHex := hex.EncodeToString(bytes)
	uniqueID := timestampHex + randomHex

	return uniqueID
}
