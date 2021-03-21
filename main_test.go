package main

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestRunMap(t *testing.T) {
	// ["TrieMap","insert","search","search","startsWith","insert","search"]
	// [[],["apple"],["apple"],["app"],["app"],["app"],["app"]]
	trie := ConstructorMap()
	trie.Insert("apple")
	assert.True(t, trie.Search("apple"))
	assert.False(t, trie.Search("app"))
	assert.True(t, trie.StartsWith("app"))
	trie.Insert("app")
	assert.True(t, trie.Search("app"))
}

func TestRunList(t *testing.T) {
	// ["TrieMap","insert","search","search","startsWith","insert","search"]
	// [[],["apple"],["apple"],["app"],["app"],["app"],["app"]]
	trie := ConstructorList()
	trie.Insert("apple")
	assert.True(t, trie.Search("apple"))
	assert.False(t, trie.Search("app"))
	assert.True(t, trie.StartsWith("app"))
	trie.Insert("app")
	assert.True(t, trie.Search("app"))
}

const letterLowercase = "abcdefghijklmnopqrstuvwxyz"
const lettersAll = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lettersAndNumbers = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func RandStringBytes(letters string) string {
	n := rand.Intn(15)
	if n == 0 {
		return RandStringBytes(letters)
	}
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func insert(b *testing.B, trie TrieInterface, letters string) {
	b.StopTimer()
	wordsNum := 100000
	words := make([]string, 0, wordsNum)
	for i := 0; i < wordsNum; i++ {
		words = append(words, RandStringBytes(letters))
	}
	b.StartTimer()
	b.ReportAllocs()

	for j := 0; j < b.N; j++ {
		for _, w := range words {
			trie.Insert(w)
		}
	}
}

func BenchmarkMapLowerCaseInsert(b *testing.B) {
	trie := ConstructorMap()
	insert(b, &trie, letterLowercase)
}

func BenchmarkListLowerCaseInsert(b *testing.B) {
	trie := ConstructorList()
	insert(b, &trie, letterLowercase)
}

func BenchmarkMapAllCaseInsert(b *testing.B) {
	trie := ConstructorMap()
	insert(b, &trie, lettersAll)
}

func BenchmarkListAllCaseInsert(b *testing.B) {
	trie := ConstructorList()
	insert(b, &trie, lettersAll)
}

func BenchmarkMapWithNumbersInsert(b *testing.B) {
	trie := ConstructorMap()
	insert(b, &trie, lettersAndNumbers)
}

func BenchmarkListWithNumbersInsert(b *testing.B) {
	trie := ConstructorList()
	insert(b, &trie, lettersAndNumbers)
}
