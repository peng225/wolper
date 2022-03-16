package dictionary

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"unicode"

	"github.com/peng225/wolper/trie"
)

type Dictionary struct {
	tr trie.Trie
}

func (dict *Dictionary) cleanUp(text string) string {
	retText := strings.ToLower(text)
	for _, v := range text {
		if !unicode.IsLower(v) {
			retText = strings.Replace(retText, string(v), " ", -1)
		}
	}
	return retText
}

func (dict *Dictionary) mergeMaps(map1, map2 map[string]bool) map[string]bool {
	result := map[string]bool{}

	for k, v := range map1 {
		result[k] = v
	}
	for k, v := range map2 {
		result[k] = v
	}
	return result
}

func (dict *Dictionary) wordScan(inputDir string) map[string]bool {
	files, err := os.ReadDir(inputDir)
	if err != nil {
		panic(err)
	}
	words := make(map[string]bool)
	for _, file := range files {
		fileFullPath := path.Join(inputDir, file.Name())

		if file.IsDir() {
			words = dict.mergeMaps(words, dict.wordScan(fileFullPath))
			continue
		}

		fmt.Println(fileFullPath)
		ifp, err := os.Open(fileFullPath)
		if err != nil {
			panic(err)
		}
		defer ifp.Close()

		scanner := bufio.NewScanner(ifp)
		for scanner.Scan() {
			line := dict.cleanUp(scanner.Text())
			tokens := strings.Split(line, " ")
			for _, token := range tokens {
				if token != "" {
					words[token] = true
				}
			}
		}
	}
	return words
}

func (dict *Dictionary) Build(inputDir string, outputFile string) {
	words := dict.wordScan(inputDir)

	ofp, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer ofp.Close()

	writer := bufio.NewWriter(ofp)
	const FLUSH_THRESHOLD = 8192
	for word, _ := range words {
		_, err = writer.WriteString(word + string('\n'))
		if err != nil {
			panic(err)
		}
		if writer.Buffered() >= FLUSH_THRESHOLD {
			writer.Flush()
		}
	}
	writer.Flush()
}

func (dict *Dictionary) Load(inputFile string) {
	ifp, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer ifp.Close()

	dict.tr = &trie.TreeTrie{}
	scanner := bufio.NewScanner(ifp)
	for scanner.Scan() {
		dict.tr.Add(scanner.Text())
	}
}

func (dict *Dictionary) Query(key string, include string, exclude string, uniq bool) []string {
	return dict.tr.Query(key, include, exclude, uniq)
}
