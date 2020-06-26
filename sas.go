package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"text/template"
	"time"
)

// GetToken returns a token using the given resource URI and access key
func GetToken(resourceURI string, sasKey string) string {
	// Storage service version
	sv := "2019-10-10"
	// Storage service permissions
	sp := "w"
	// Signed resource type (blob)
	sr := "b"
	// Signed protocol
	spr := "https"

	// Format expire time
	expireTime := time.Now().Add(time.Duration(2) * time.Second)
	// Signed Expiry
	se := expireTime.UTC().Format(time.RFC3339)

	// Escape uri string
	uri := template.URLQueryEscaper(resourceURI)
	// Get Signature
	sigData := uri + "\n" + se
	rawSig := getHmac256(sigData, sasKey)
	// Signature
	sig := template.URLQueryEscaper(rawSig)

	return fmt.Sprintf("sv=%s&sp=%s&sr=%s&spr=%s&se=%s&sig=%s", sv, sp, sr, spr, se, sig)
}

func getHmac256(str string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
