package mutube

import (
	"crypto/rand"
	"encoding/base32"
	"time"
)

const maxSessionLength = 30 * 25 * time.Hour

func generateSessionId() string {
	bytes := make([]byte, 15)
	rand.Read(bytes)
	sessionId := base32.StdEncoding.EncodeToString(bytes)
	return sessionId
}
