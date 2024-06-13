package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

func main() {
	// tworzymy klucz 256-bitowy
	aes_256_key := make([]byte, 32)
	rand.Read(aes_256_key)

	// tworzymy blok szyfrujący AES-256
	aes_cipher_block, _ := aes.NewCipher(aes_256_key)
	block_size := aes_cipher_block.BlockSize()

	// wiadomość oryginalna musi być wielokrotnością wielkości bloku szyfrujacego
	original_message := []byte("Hello, World!")
	padded_message := pkcs7Padding(original_message, block_size)

	// CBC - cipher block chaining
	// IV - initialization vector
	// pierwszy blok jest XORowany z wektorem inicjalizacyjnym (IV)
	// kolejne bloki są XORowane z poprzednimi zaszyfrowanymi blokami

	// wektor inicjalizacyjny
	iv := make([]byte, block_size)
	rand.Read(iv)

	// przygotowanie enkryptera CBC
	// (pozwala on na szyfrowanie wielu blokow danych,
	// dzieki czemu nie trzeba tego robic recznie uzywajac samego bloku
	// szyfrujacego AES)
	cbc_encrypter := cipher.NewCBCEncrypter(aes_cipher_block, iv)

	// szyfrowanie wiadomosci
	encrypted_message := make([]byte, len(padded_message))
	cbc_encrypter.CryptBlocks(encrypted_message, padded_message)

	fmt.Printf("Original message (text):  %s\n", original_message)
	fmt.Printf("Original message (byte):  %x\n", original_message)
	fmt.Printf("Padded message:           %x\n", padded_message)
	fmt.Printf("Initialization vector:    %x\n", iv)
	fmt.Printf("Encrypted message:        %x\n\n", encrypted_message)

	// odszyfrowanie uzywajac dekryptera CBC i usuniecie paddingu
	// potrzebny jest wektor inicjalizacyjny i oczywiscie blok AES-256 stworzony z klucza
	// IV moze byc publiczny, wazne zeby generowany byl nowy dla kazdej wiadomosci.
	// Klucz AES-256 musi byc tajny, znany tylko przez nadawce i odbiorce.
	plaintext_padded_message := make([]byte, len(encrypted_message))
	cbc_decrypter := cipher.NewCBCDecrypter(aes_cipher_block, iv)
	cbc_decrypter.CryptBlocks(plaintext_padded_message, encrypted_message)
	plaintext_message := pkcs7Unpadding(plaintext_padded_message)

	fmt.Printf("Padded decrypted message (byte): %x\n", plaintext_padded_message)
	fmt.Printf("Plaintext message (byte):        %x\n", plaintext_message)
	fmt.Printf("Plaintext message (text):        %s\n", plaintext_message)
}

func pkcs7Padding(message []byte, blockSize int) []byte {
	padding := blockSize - len(message)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(message, padtext...)
}

func pkcs7Unpadding(message []byte) []byte {
	length := len(message)
	unpadding := int(message[length-1])
	return message[:(length - unpadding)]
}
