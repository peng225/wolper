package dictionary

import (
	"testing"

	"github.com/peng225/wolper/trie"
	"github.com/stretchr/testify/suite"
)

/*******************************/
/* Test set up                 */
/*******************************/
type DictionarySuite struct {
	suite.Suite
	dict Dictionary
}

func (suite *DictionarySuite) SetupTest() {
	suite.dict = Dictionary{&trie.TreeTrie{}}
}

/*******************************/
/* Expose private methods      */
/*******************************/
func CleanUp(dict *Dictionary, text string) string {
	return dict.cleanUp(text)
}

func WordScan(dict *Dictionary, inputDir string) map[string]bool {
	return dict.wordScan(inputDir)
}

/*******************************/
/* Test cases                  */
/*******************************/
func (suite *DictionarySuite) TestCealnupA() {
	suite.Equal("abcde", CleanUp(&suite.dict, "abcde"))
	suite.Equal("abcde", CleanUp(&suite.dict, "Abcde"))
	suite.Equal("a b c d e f g", CleanUp(&suite.dict, "A;b.c,d!e2F\"g"))
	suite.Equal("  xy   z  ", CleanUp(&suite.dict, "89xy ;+z.\n"))
}

func (suite *DictionarySuite) TestWordScan() {
	words := WordScan(&suite.dict, "test_input")
	suite.Contains(words, "hello")
	suite.Contains(words, "world")
	suite.Contains(words, "i")
	suite.Contains(words, "love")
	suite.Contains(words, "golang")
	suite.NotContains(words, "java")
}

func (suite *DictionarySuite) TestWordScanWithSlash() {
	words := WordScan(&suite.dict, "test_input/")
	suite.Contains(words, "hello")
	suite.Contains(words, "world")
	suite.Contains(words, "i")
	suite.Contains(words, "love")
	suite.Contains(words, "golang")
	suite.NotContains(words, "java")
}

/*******************************/
/* Run tests                   */
/*******************************/
func TestTreeTrieSuite(t *testing.T) {
	suite.Run(t, new(DictionarySuite))
}
