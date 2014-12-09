package word2vec

import (
	"os"
)

const (
	MaxString     = 100
	VocabHashSize = 30000000
)

type VocabWord struct {
	Word        string
	Count       uint64
	Point, Code []int
	CodeLen     int
}

var Vocab = make(map[string]*VocabWord)
var VocabHash = make(map[string]int)

func LearnVocab(fn string) {
	file, err := os.Open(fn)
}

func addWord(word string) {
	if Vocab[word] == nil {
		Vocab[word] = &VocabWord{
			word,
			1,
			[]int{},
			[]int{},
			0,
		}
		VocabHash[word] = hashWord(word)
	} else {
		Vocab[word].Count++
	}
}

func hashWord(word string) uint {
	var hash uint64 = 0
	for _, c := range word {
		hash = hash*257 + uint64(c)
	}
	return hash % VocabHashSize
}
