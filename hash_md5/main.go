package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	// Hasła do zahashowania
	password1 := "admin0"
	password2 := "admin1"

	// Obliczenie hashy md5
	hash1 := md5.Sum([]byte(password1))
	hash2 := md5.Sum([]byte(password2))

	// Wypisanie stworoznych hashy
	fmt.Println("hash1: ", hex.EncodeToString(hash1[:]))
	fmt.Println("hash2: ", hex.EncodeToString(hash2[:]))
}
