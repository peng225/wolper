package trie

type TreeTrie struct {
	c        byte
	terminal bool
	sibling  *TreeTrie
	child    *TreeTrie
}

func (treeTrie *TreeTrie) Query(key string, include string, exclude string, current string) []string {
	if key == "" {
		if treeTrie.terminal {
			return []string{current}
		} else {
			return make([]string, 0)
		}
	}

	firstChar := key[0]
	key = key[1:]
	currentNode := treeTrie.child
	result := make([]string, 0)
	for currentNode != nil {
		if currentNode.c == firstChar {
			current += string(firstChar)
			result = append(result, currentNode.Query(key, include, exclude, current)...)
			// current = current[:len(current)-2]
			break
		}
		currentNode = currentNode.sibling
	}
	return result
}

func (treeTrie *TreeTrie) Add(key string) {
	if key == "" {
		treeTrie.terminal = true
		return
	}
	firstChar := key[0]
	key = key[1:]

	if treeTrie.child == nil {
		treeTrie.child = &TreeTrie{firstChar, false, nil, nil}
	}

	currentNode := treeTrie.child
	for currentNode != nil {
		if currentNode.c == firstChar {
			currentNode.Add(key)
			return
		} else if currentNode.sibling == nil {
			currentNode.sibling = &TreeTrie{firstChar, false, nil, nil}
		} else if firstChar > currentNode.sibling.c {
			prevSibling := currentNode.sibling
			currentNode.sibling = &TreeTrie{firstChar, false, prevSibling, nil}
		}
		currentNode = currentNode.sibling
	}

}

func (treeTrie *TreeTrie) Delete(key string) {
}
