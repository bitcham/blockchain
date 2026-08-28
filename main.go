package main

import (
	"crypto/sha256"
	"fmt"
)

func hash(data string) string {
	sum := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", sum)
}

func main() {
	a := "alice sends 10 to bob"
	b := "alice sends 10 to bob"
	c := "alice sends 11 to bob"

	fmt.Println("same input twice:")
	fmt.Println(" ", hash(a))
	fmt.Println(" ", hash(b))
	fmt.Println()
	fmt.Println("one digit changed (10 → 11):")
	fmt.Println(" ", hash(c))
}
