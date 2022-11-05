package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func fileExists(f string) bool {
	if _, err := os.Stat(f); errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}

func fileNotExists(f string) bool {
	return !fileExists(f)
}

// Proper key injected at build time
var EncryptionKey string = "dummy"

// func encBytes(data []byte, key string) ([]byte, error) {
// 	block, _ := aes.NewCipher([]byte(createHash(key)))
// 	gcm, err := cipher.NewGCM(block)
// 	if err != nil {
// 		return nil, err
// 	}
// 	nonce := make([]byte, gcm.NonceSize())
// 	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
// 		return nil, err
// 	}
// 	ciphertext := gcm.Seal(nonce, nonce, data, nil)
// 	return ciphertext, nil
// }

// func decBytes(data []byte, key string) ([]byte, error) {
// 	keyHash := []byte(createHash(key))
// 	block, err := aes.NewCipher(keyHash)
// 	if err != nil {
// 		return nil, err
// 	}
// 	gcm, err := cipher.NewGCM(block)
// 	if err != nil {
// 		return nil, err
// 	}
// 	nonceSize := gcm.NonceSize()
// 	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
// 	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
// 	if err != nil {
// 		return nil, errors.New("failed to decrypt")
// 	}
// 	return plaintext, nil
// }

// func createHash(key string) string {
// 	hasher := md5.New()
// 	hasher.Write([]byte(key))
// 	return hex.EncodeToString(hasher.Sum(nil))
// }

func rndName() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 10)
	// Read b number of numbers
	rand.Read(b)
	//return fmt.Sprintf("%x", b)[:16]
	return fmt.Sprintf("%x", b)
}
