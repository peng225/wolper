package trie

type Trie interface {
	Query(key string, include string, exclude string, uniq bool) []string
	Add(key string)
	Delete(key string)
}
