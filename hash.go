package main

import (
	"crypto/sha256"
	"fmt"
)

func hash(data string) string {
	sum := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", sum)
}
