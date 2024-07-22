package bip39

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/ottosch/lastseed/src/bip39/wordlist"
)

// VerifyWordCount checks if the number of words matches a suitable number of bip39 seed words
func VerifyWordCount(words []string) {
	actualCount := len(words) + 1
	if actualCount < 12 || actualCount > 24 || actualCount%3 != 0 {
		fmt.Fprintf(os.Stderr, "%d word(s) input. Expected 11, 14, 17, 20 or 23\n", len(words))
		os.Exit(1)
	}
}

// VerifyValidWords checks if all words belong to bip39 word list
func VerifyValidWords(words []string) {
	for _, w := range words {
		_, found := wordlist.WordMap[w]
		if !found {
			fmt.Fprintf(os.Stderr, "Word \"%s\" is not part of the BIP39 word list\n", w)
			os.Exit(1)
		}
	}
}

// ValidChecksum performs checksum validation
func ValidChecksum(fullBitstring *big.Int, checksumSize uint) bool {
	wordCount := int(checksumSize) * 3

	withoutChecksumBits := new(big.Int).Set(fullBitstring)
	withoutChecksumBits.Rsh(withoutChecksumBits, checksumSize)

	withoutChecksumBitCount := (wordCount - 1) * 11 // we still don't have the last word here
	withoutChecksumByteLen := (withoutChecksumBitCount + 7) / 8
	withoutChecksumBytes := make([]byte, withoutChecksumByteLen)
	withoutChecksumBits.FillBytes(withoutChecksumBytes)

	hashBits := sha256.Sum256(withoutChecksumBytes)[0] >> (8 - checksumSize) // ex 4 bits: 01010111000110001101 -> 00001000
	bitmask := byte((1 << checksumSize) - 1)                                 // 00001111
	checksumBits := hashBits & bitmask                                       // 00001000

	withChecksumBitCount := wordCount*11 + int(checksumSize) // now all words + checksum
	withChecksumByteLen := (withChecksumBitCount + 7) / 8
	withChecksumBytes := make([]byte, withChecksumByteLen)
	fullBitstring.FillBytes(withChecksumBytes)

	inputChecksumBits := withChecksumBytes[len(withChecksumBytes)-1] & bitmask
	return checksumBits == inputChecksumBits
}
