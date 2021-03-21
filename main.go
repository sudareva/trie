package main

type TrieInterface interface {
	Insert(string)
	Search(string) bool
	StartsWith(string) bool
}

type TrieMap struct {
	isWord        bool // ключ - является полноценным добавленным словом
	childrenLinks map[rune]*TrieMap
}

/** Initialize your data structure here. */
func ConstructorMap() TrieMap {
	return TrieMap{
		isWord:        false,
		childrenLinks: make(map[rune]*TrieMap),
	}

}

/** Inserts a word into the trie. */
func (t *TrieMap) Insert(word string) {
	currentNode := t
	for _, k := range word {
		if n, ok := currentNode.childrenLinks[k]; ok {
			currentNode = n
			continue
		}
		nextNode := &TrieMap{}
		if len(currentNode.childrenLinks) == 0 {
			currentNode.childrenLinks = map[rune]*TrieMap{k: nextNode}
		} else {
			currentNode.childrenLinks[k] = nextNode
		}
		currentNode = nextNode
	}

	currentNode.isWord = true
}

func (t TrieMap) search(word string, searchWord bool) bool {
	currentNode := &t
	for _, k := range word {
		if n, ok := currentNode.childrenLinks[k]; ok {
			currentNode = n
			continue
		}

		return false
	}
	if searchWord {
		return currentNode.isWord
	}

	return true
}

/** Returns if the word is in the trie. */
func (t TrieMap) Search(word string) bool {
	return t.search(word, true)
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t TrieMap) StartsWith(prefix string) bool {
	return t.search(prefix, false)
}