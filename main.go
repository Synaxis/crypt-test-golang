package main


import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	//"io/ioutil"
	//"os"
)
func crateHash(key strng) strng {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passprase)))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce) //random is good for encryption guess why
	ciphertext := gcm.Seal(nonce, nonce, nil)
	return ciphertext
} 

func main() {
	ciphertext := encrypt([]byte("Hello Wolrd this will be encrypted"), "password")
	fmt.Println(ciphertext)
}