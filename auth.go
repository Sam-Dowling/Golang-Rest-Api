package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
)

const (
	signingKeyPath = "/home/sam/keys/server.key.pem"
	serverCertPath = "/home/sam/keys/server-local.cert.pem"
)

var signingKey, _ = ioutil.ReadFile(signingKeyPath)

//SessionTokenResponse holds the format of the sessionToken
type SessionTokenResponse struct {
	Token string `json:"sessionToken"`
}

func readToken(r *http.Request) (*jwt.Token, error) {
	return jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
}

func isValidToken(token *jwt.Token) bool {
	if !token.Valid {
		return false
	}
	sessionID := token.Claims["jti"].(string)
	if !getSession(sessionID) {
		return false
	}
	return true
}

func writeJSON(w http.ResponseWriter, status int, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(response)
}

func startSession() string {
	sessionID := base64.URLEncoding.EncodeToString(uuid.NewV4().Bytes())
	addSession(sessionID)
	return sessionID
}

func generateToken(sessionID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["jti"] = sessionID
	token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	return token.SignedString(signingKey)
}

func sha256Hash(pass string, salt string) string {
	s := fmt.Sprintf("%s%s", salt, pass)
	h := sha256.New()
	h.Write([]byte(s))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}
