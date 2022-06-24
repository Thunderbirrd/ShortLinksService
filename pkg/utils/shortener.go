package utils

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	uuid "github.com/nu7hatch/gouuid"
	"math/big"
	"os"
)

const length = 10

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))

	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding

	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(encoded)
}

func GenerateShortUrl(longUrl string) string {
	u, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	urlHashBytes := sha256Of(longUrl + u.String())
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))

	return finalString[:length]
}
