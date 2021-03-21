package main

type TrieList struct {
	isWord        bool
	link          rune
	childrenLinks []*TrieList
}

/** Initialize your data structure here. */
func ConstructorList() TrieList {
	return TrieList{
		isWord:        false,
		link:          0,
		childrenLinks: []*TrieList{},
	}
}

func (t *TrieList) findChild(link rune) (*TrieList, bool) {
	for _, node := range t.childrenLinks {
		if node.link == link {
			return node, true
		}
	}

	return nil, false
}

/** Inserts a word into the trie. */
func (t *TrieList) Insert(word string) {
	currentNode := t
	for _, letter := range word {
		if childNode, ok := currentNode.findChild(letter); ok {
			currentNode = childNode
			continue
		}

		newChildNode := &TrieList{
			link:   letter,
		}
		currentNode.childrenLinks = append(currentNode.childrenLinks, newChildNode)
		currentNode = newChildNode
	}
	currentNode.isWord = true
}

func (t TrieList) search(word string, searchWord bool) bool {
	currentNode := &t
	lastIdx := len(word) - 1
	for idx, k := range word {
		if childNode, ok := currentNode.findChild(k); ok {
			if idx == lastIdx {
				if searchWord {
					return childNode.isWord
				}
				return true
			}
			currentNode = childNode
			continue

		}

		return false
	}

	return false
}

/** Returns if the word is in the trie. */
func (t TrieList) Search(word string) bool {
	return t.search(word, true)
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t TrieList) StartsWith(prefix string) bool {
	return t.search(prefix, false)
}
