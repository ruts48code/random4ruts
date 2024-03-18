package random4ruts

import (
	"crypto/rand"
	"log"
	"math/big"
)

func GetRandomInt(n int64) int64 {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(n))
	return nBig.Int64()
}

func GetRandomBytes(n int64) []byte {
	output := make([]byte, n)
	_, err := rand.Read(output)
	if err != nil {
		log.Printf("Error: random4ruts-GetRandomBytes - read random error - %v\n", err)
		return make([]byte, 0)
	}
	return output
}

func GetRandomString(n int) string {
	salt := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	output := ""
	for i := 0; i < n; i++ {
		output = output + string(salt[GetRandomInt(int64(len(salt)))])
	}
	return output
}
