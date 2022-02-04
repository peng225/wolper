package dictionary

import (
    "io/ioutil"
    "strings"
    "os"
    "fmt"
    "bufio"

	"github.com/peng225/wolper/trie"
)

type Dictionary struct {
    tr trie.Trie
}

func (dict *Dictionary) cleanUp(text string) string {
    retText := strings.Replace(text, ".", " ", -1)
    retText = strings.Replace(retText, "?", " ", -1)
    retText = strings.Replace(retText, "!", " ", -1)
    retText = strings.Replace(retText, "#", " ", -1)
    retText = strings.Replace(retText, "$", " ", -1)
    retText = strings.Replace(retText, "&", " ", -1)
    retText = strings.Replace(retText, "(", " ", -1)
    retText = strings.Replace(retText, ")", " ", -1)
    retText = strings.Replace(retText, ",", " ", -1)
    retText = strings.Replace(retText, ":", " ", -1)
    retText = strings.Replace(retText, ";", " ", -1)
    retText = strings.Replace(retText, "\"", " ", -1)
    retText = strings.Replace(retText, "'", " ", -1)
    retText = strings.Replace(retText, "~", " ", -1)
    retText = strings.Replace(retText, "=", " ", -1)
    retText = strings.Replace(retText, "\n", "", -1)
    retText = strings.ToLower(retText)
    return retText
}

func (dict *Dictionary) Build(inputDir string, outputFile string) {
    files, err := ioutil.ReadDir(inputDir)
    if err != nil {
        panic(err)
    }

    ofp, err := os.Create(outputFile)
    if err != nil {
        panic(err)
    }
    defer ofp.Close()

    words := make(map[string]bool)
    for _, file := range files {
        if file.IsDir() {
            continue
        }

        fileFullPath := inputDir + "/" + file.Name()
        // allow input both form of "path" and "path/" for inputDir
        fileFullPath = strings.Replace(fileFullPath, "//", "/", -1)
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

    writer := bufio.NewWriter(ofp)
    const FLUSH_THRESHOLD = 8192
    for word, _ := range words {
        _, err = writer.WriteString(word + "\n")
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

func (dict *Dictionary) Query(key string, include string, exclude string) []string {
    return dict.tr.Query(key, include, exclude, "")
}

