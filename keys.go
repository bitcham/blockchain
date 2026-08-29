package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
)

type KeyPair struct {
	Public  ed25519.PublicKey
	Private ed25519.PrivateKey
}

func NewKeyPair() (KeyPair, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return KeyPair{}, err
	}
	return KeyPair{Public: pub, Private: priv}, nil
}

func (k KeyPair) Address() string {
	return hex.EncodeToString(k.Public)
}

func (k KeyPair) Sign(payload string) string {
	sig := ed25519.Sign(k.Private, []byte(payload))
	return hex.EncodeToString(sig)
}

func Verify(pubHex, payload, sigHex string) bool {
	pub, err := hex.DecodeString(pubHex)
	if err != nil || len(pub) != ed25519.PublicKeySize {
		return false
	}
	sig, err := hex.DecodeString(sigHex)
	if err != nil {
		return false
	}
	return ed25519.Verify(ed25519.PublicKey(pub), []byte(payload), sig)
}
