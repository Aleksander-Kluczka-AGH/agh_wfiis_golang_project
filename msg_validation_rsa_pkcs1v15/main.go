package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	// strona A zna klucz prywatny, strona B klucz publiczny
	pkey, _ := rsa.GenerateKey(rand.Reader, 2048)

	// strona A tworzy wiadomość, podpisuje i wysyła do strony B
	message := []byte("Hello, World!")
	hashed := sha256.Sum256(message)
	fmt.Printf("Hashed: %x\n", hashed)

	signature, _ := rsa.SignPKCS1v15(rand.Reader, pkey, crypto.SHA256, hashed[:])
	fmt.Printf("Signature: %x\n", signature)

	// strona B odbiera wiadomość i podpis, sprawdza czy podpis jest poprawny
	err := rsa.VerifyPKCS1v15(&pkey.PublicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		// wiadomość została zmieniona
		fmt.Println("Verification failed")
	} else {
		// wiadomość jest oryginalna
		fmt.Println("Verification succeeded")
	}
}
