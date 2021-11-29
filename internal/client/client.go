package client

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"github.com/ThrynSec/aes_implementation_poc/externals/getweb"
)

func CreateKey(keyStr string) {
	os.WriteFile("./tmp_aes.key", []byte(keyStr), 6044)
}

func SendSecure(strContent string) {
	key := getKey()

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	htmlRequest := gcm.Seal(nonce, nonce, []byte(strContent), nil)

	fmt.Println(getweb.GetAPI("http://localhost:8080/dns/secure/" + string(htmlRequest)))
}

func SendUnsecure(strContent string) {
	fmt.Println(fmt.Println(getweb.GetAPI("http://localhost:8080/dns/unsecure/" + strContent)))
}

func SendMessage(strContent string) {
	key := getKey()

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	htmlRequest := gcm.Seal(nonce, nonce, []byte(strContent), nil)

	fmt.Println(getweb.GetAPI("http://localhost:8080/decrypt/" + string(htmlRequest)))
}

func getKey() []byte {
	key, _ := os.ReadFile("./tmp_aes.key")
	return key
}
