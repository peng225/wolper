package trie

type Trie interface {
	Query(key, include, exclude string, posExcludeList []string, uniq bool) []string
	Add(key string)
	Delete(key string)
}
