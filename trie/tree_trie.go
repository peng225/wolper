package trie

import (
	"strings"

	"github.com/peng225/wolper/util"
)

type TreeTrie struct {
	c        byte
	terminal bool
	sibling  *TreeTrie
	child    *TreeTrie
}

func (treeTrie *TreeTrie) query(key string, include string, exclude string, uniq bool, current string) []string {
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
		if !strings.Contains(exclude, string(currentNode.c)) &&
			!(uniq && strings.Contains(current, string(currentNode.c))) {
			newInclude := include
			if firstChar == '.' {
				if strings.Contains(include, string(currentNode.c)) {
					newInclude = string(util.DeleteChar([]byte(include), currentNode.c))
				}
				current += string(currentNode.c)
				result = append(result, currentNode.query(key, newInclude, exclude, uniq, current)...)
				current = current[:len(current)-1]
			} else if currentNode.c == firstChar {
				if strings.Contains(include, string(currentNode.c)) {
					newInclude = string(util.DeleteChar([]byte(include), currentNode.c))
				}
				current += string(currentNode.c)
				result = append(result, currentNode.query(key, newInclude, exclude, uniq, current)...)
				break
			}
		}
		currentNode = currentNode.sibling
	}
	return result
}

func (treeTrie *TreeTrie) Query(key string, include string, exclude string, uniq bool) []string {
	return treeTrie.query(key, include, exclude, uniq, "")
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
	} else if treeTrie.child.c > firstChar {
		treeTrie.child = &TreeTrie{firstChar, false, treeTrie.child, nil}
	}

	currentNode := treeTrie.child
	for currentNode != nil {
		if currentNode.c == firstChar {
			currentNode.Add(key)
			return
		} else if currentNode.sibling == nil {
			currentNode.sibling = &TreeTrie{firstChar, false, nil, nil}
		} else if firstChar < currentNode.sibling.c {
			prevSibling := currentNode.sibling
			currentNode.sibling = &TreeTrie{firstChar, false, prevSibling, nil}
		}
		currentNode = currentNode.sibling
	}
}

func (treeTrie *TreeTrie) Delete(key string) {
}
