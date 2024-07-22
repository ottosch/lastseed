package seed

import (
	"math"
	"math/big"
	"strings"

	"github.com/ottosch/lastseed/src/bip39"
	"github.com/ottosch/lastseed/src/bip39/wordlist"
)

type Seed struct {
	bitstring *big.Int // the bits of the seed words
	words     []string
	results   []*Result
}

func NewSeed(wordsStr string) *Seed {
	words := strings.Split(wordsStr, " ")

	bip39.VerifyWordCount(words)
	bip39.VerifyValidWords(words)

	s := &Seed{words: words}
	s.fillBitstring(words)

	s.calcLastWords()
	return s
}

func (s *Seed) GetWords() []string {
	return s.words
}

func (s *Seed) GetWordCount() int {
	return len(s.words) + 1
}

func (s *Seed) GetChecksumSize() uint {
	return uint(s.GetWordCount() / 3)
}

func (s *Seed) GetBitstring() *big.Int {
	return new(big.Int).Set(s.bitstring)
}

func (s *Seed) GetResults() []*Result {
	return s.results
}

func (s *Seed) calcLastWords() {
	totalWords := int(math.Pow(2, float64(11-s.GetChecksumSize())))
	results := make([]*Result, 0, totalWords)

	for i, bipWord := range wordlist.Wordlist {
		fullBitstring := addWordToBitstring(s.GetBitstring(), i)
		if bip39.ValidChecksum(fullBitstring, s.GetChecksumSize()) {
			results = append(results, NewResult(fullBitstring, bipWord))
		}
	}

	s.results = results
}

// fillBitstring sets the seed bit string according to the words indices
func (s *Seed) fillBitstring(words []string) {
	num := new(big.Int)
	for _, w := range words {
		index := int64(wordlist.WordMap[w])
		indexBig := new(big.Int).SetInt64(index)
		num.Or(num.Lsh(num, 11), indexBig) // num << 11 | index
	}

	s.bitstring = num
}

// addWordToBitstring adds the word index to the current bit string
func addWordToBitstring(bitstring *big.Int, wordIndex int) *big.Int {
	fullBitstring := new(big.Int).Set(bitstring)
	lastWordIndex := new(big.Int).SetInt64(int64(wordIndex))
	fullBitstring.Or(fullBitstring.Lsh(fullBitstring, 11), lastWordIndex)
	return fullBitstring
}
