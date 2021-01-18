package utils

import (
	"crypto/rand"
	"math/big"
)

/*
* RandomString - Generated random string based on given length
 */
func RandomString(length int) ([]byte, error) {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	symbols := big.NewInt(int64(len(alphanum)))
	states := big.NewInt(0)
	states.Exp(symbols, big.NewInt(int64(length)), nil)
	r, err := rand.Int(rand.Reader, states)
	if err != nil {
		return nil, err
	}
	var bytes = make([]byte, length)
	r2 := big.NewInt(0)
	symbol := big.NewInt(0)
	for i := range bytes {
		r2.DivMod(r, symbols, symbol)
		r, r2 = r2, r
		bytes[i] = alphanum[symbol.Int64()]
	}
	return bytes, err
}
