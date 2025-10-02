package helpers

import (
	"crypto/rand"
	"math/big"
)

func GenerateSixDigitCode() (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(900000))

	if err != nil {
		return 0, err
	}

	return int(n.Int64()) + 100000, nil
}
