package trie

type Trie interface {
    Query(key string, include string, exclude string, current string) []string
    Add(key string)
    Delete(key string)
}

