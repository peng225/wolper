package trie

import "strings"

type TreeTrie struct {
	c        byte
	terminal bool
	sibling  *TreeTrie
	child    *TreeTrie
}

func deleteChar(s []byte, c byte) []byte {
	foundIndex := -1
	for i, v := range s {
		if v == c {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		return s
	} else {
		s[foundIndex] = s[len(s)-1]
		return s[:len(s)-1]
	}
}

func (treeTrie *TreeTrie) Query(key string, include string, exclude string, current string) []string {
	if key == "" {
		if treeTrie.terminal && include == "" {
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
		if !strings.Contains(exclude, string(currentNode.c)) {
			newInclude := include
			if firstChar == '.' {
				foundInclude := false
				if strings.Contains(include, string(currentNode.c)) {
					foundInclude = true
					newInclude = string(deleteChar([]byte(include), currentNode.c))
				}
				current += string(currentNode.c)
				result = append(result, currentNode.Query(key, newInclude, exclude, current)...)
				current = current[:len(current)-1]
				if foundInclude {
					include += string(currentNode.c)
				}
			} else if currentNode.c == firstChar {
				if strings.Contains(include, string(currentNode.c)) {
					newInclude = string(deleteChar([]byte(include), currentNode.c))
				}
				current += string(currentNode.c)
				result = append(result, currentNode.Query(key, newInclude, exclude, current)...)
				break
			}
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
