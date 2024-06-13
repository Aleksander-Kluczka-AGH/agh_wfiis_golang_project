package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func bruteForce(length int, hash []byte) {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsetLen := len(charset)
	indices := make([]int, length)
	current := make([]byte, length)

	start := time.Now()
	for {
		// Tworzenie aktualnego hasła
		for i, idx := range indices {
			current[i] = charset[idx]
		}

		// Obliczenie hasha aktualnego hasła
		currentHash := md5.Sum(current)
		if hex.EncodeToString(currentHash[:]) == hex.EncodeToString(hash) {
			fmt.Println("Found password: ", string(current))
			fmt.Println("Time elapsed: ", time.Since(start))
			return
		}

		// Przygotowanie nowych indeksów
		for i := length - 1; i >= 0; i-- {
			indices[i]++
			if indices[i] < charsetLen {
				break
			}
			indices[i] = 0
		}
	}
}

func main() {
	// Stworzenie hasła i hasha
	password := "pass"
	passwordHash := md5.Sum([]byte(password))

	// Łamanie hasła
	bruteForce(len(password), passwordHash[:])
}
