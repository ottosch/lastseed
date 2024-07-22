package seed

import (
	"encoding/hex"
	"math/big"
)

type Result struct {
	bitstring *big.Int
	lastWord  string
}

func NewResult(bitstring *big.Int, lastWord string) *Result {
	return &Result{bitstring, lastWord}
}

func (r *Result) Bitstring() *big.Int {
	return r.bitstring
}

func (r *Result) LastWord() string {
	return r.lastWord
}

// Entropy returns entropy in hex
func (r *Result) Entropy(checksumSize uint) string {
	wordCount := int(checksumSize) * 3

	withoutChecksumBits := new(big.Int).Set(r.bitstring)
	withoutChecksumBits.Rsh(withoutChecksumBits, checksumSize)

	bitCount := wordCount*11 - int(checksumSize) // entropy doesn't have checksum
	byteLen := (bitCount + 7) / 8
	fbs := make([]byte, byteLen)
	withoutChecksumBits.FillBytes(fbs)

	return hex.EncodeToString(fbs)
}
