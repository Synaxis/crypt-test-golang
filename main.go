package main


import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"crypto/rand"
	//"io/ioutil"
	//"os"
)
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)	
	if err != nil {
		fmt.Println("error with NewGCM line23\n")
	}

	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce) //random is good for encryption
	ciphertext := gcm.Seal(nonce, nonce, data, nil) 
	return ciphertext
} 

func main() {
	ciphertext := encrypt([]byte("Hello Wolrd this will be encrypted"), "password")
	fmt.Println(string(ciphertext)) //converts the buffer into a encoded str
}