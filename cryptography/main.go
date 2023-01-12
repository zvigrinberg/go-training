package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

func shaHashing(input string) string {
	plainText := []byte(input)
	sha256Hash := sha256.Sum256(plainText)
	return hex.EncodeToString(sha256Hash[:])
}

func mdHashing(input string) string {
	byteInput := []byte(input)
	md5Hash := md5.Sum(byteInput)
	return hex.EncodeToString(md5Hash[:]) // by referring to it as a string
}

func Encrypt(value []byte, keyPhrase string) []byte {

	aesBlock, err := aes.NewCipher([]byte(keyPhrase))
	if err != nil {
		fmt.Println(err)
	}
	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcmInstance.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	cipheredText := gcmInstance.Seal(nonce, nonce, value, nil)

	return cipheredText
}

func Decrypt(ciphered []byte, keyPhrase string) []byte {
	aesBlock, err := aes.NewCipher([]byte(keyPhrase))
	if err != nil {
		log.Fatalln(err)
	}
	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		log.Fatalln(err)
	}
	nonceSize := gcmInstance.NonceSize()
	nonce, cipheredText := ciphered[:nonceSize], ciphered[nonceSize:]

	originalText, err := gcmInstance.Open(nil, nonce, cipheredText, nil)
	if err != nil {
		log.Fatalln(err)
	}
	return originalText
}

func main() {
	keyPhrase := "abcdefghijklmnopqrstuvwxyz123456"
	encryptedValue:= Encrypt([]byte("Hi There How are you today?? Sending Regards!"), keyPhrase);
	fmt.Println("encrypted value bytes=" , encryptedValue)
	fmt.Println("encrypted value=" , string(encryptedValue))
	fmt.Println("decrypted value bytes=" , Decrypt(encryptedValue, keyPhrase))
	fmt.Println("decrypted value=" , string(Decrypt(encryptedValue, keyPhrase)))



}
