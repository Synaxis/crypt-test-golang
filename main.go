package main

func crateHash(key strng) strng {
	hasher := md5.New()
	hasher.Write([]byte(key)
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passprase)))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce) //random is good for encryption guess why
	ciphertext := gcm.Seal(nonce)
} 

func main() {

}