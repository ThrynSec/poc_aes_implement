package control

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"net"
	"os"

	"github.com/gin-gonic/gin"
)

func Nameserver(c *gin.Context) {
	nameserver, _ := net.LookupNS(c.Param("nameserver"))

	for _, ns := range nameserver {
		c.String(200, ns.Host)
	}
}

func NameserverAES(c *gin.Context) {
	ciphertext := c.Param("nameserver")
	finalStr := DecryptReturn(ciphertext)
	nameserver, _ := net.LookupNS(finalStr)

	for _, ns := range nameserver {
		c.String(200, ns.Host)
	}
}

func DecryptReturn(nameserver string) string {
	ciphertext := nameserver

	key := getKey()

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		fmt.Println(err)
	}

	return string(plaintext)
}

func DecryptAES(c *gin.Context) {
	ciphertext := c.Param("nameserver")
	finalStr := DecryptReturn(ciphertext)

	c.String(200, finalStr)
}

func getKey() []byte {
	key, _ := os.ReadFile("./tmp_aes.key")
	return key
}
