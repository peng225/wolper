package trie

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

/*******************************/
/* Test set up                 */
/*******************************/
type TreeTrieSuite struct {
	suite.Suite
	tr Trie
}

func (suite *TreeTrieSuite) SetupTest() {
	suite.tr = &TreeTrie{}
}

/*******************************/
/* Test cases                  */
/*******************************/
func (suite *TreeTrieSuite) TestEmpty() {
	suite.Equal(make([]string, 0), suite.tr.Query("hello", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("", "", ""))
}

func (suite *TreeTrieSuite) TestEmptyAdded() {
	suite.tr.Add("")
	suite.Equal(make([]string, 0), suite.tr.Query("hello", "", ""))
	suite.Equal([]string{""}, suite.tr.Query("", "", ""))
}

func (suite *TreeTrieSuite) TestOneWord() {
	suite.tr.Add("hello")
	suite.Equal([]string{"hello"}, suite.tr.Query("hello", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("honey", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("hell", "", ""))
}

func (suite *TreeTrieSuite) TestTwoWords() {
	suite.tr.Add("hello")
	suite.tr.Add("point")
	suite.Equal([]string{"hello"}, suite.tr.Query("hello", "", ""))
	suite.Equal([]string{"point"}, suite.tr.Query("point", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("honey", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("power", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("hell", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("poi", "", ""))
}

func (suite *TreeTrieSuite) TestTwoWordsSharingPrefixInOrder() {
	suite.tr.Add("point")
	suite.tr.Add("power")
	suite.Equal([]string{"point"}, suite.tr.Query("point", "", ""))
	suite.Equal([]string{"power"}, suite.tr.Query("power", "", ""))
	suite.Equal([]string{"point", "power"}, suite.tr.Query("po...", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("honey", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("pondy", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("po", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("pow", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("poi", "", ""))
}

func (suite *TreeTrieSuite) TestTwoWordsSharingPrefixOutOfOrder() {
	suite.tr.Add("power")
	suite.tr.Add("point")
	suite.Equal([]string{"power"}, suite.tr.Query("power", "", ""))
	suite.Equal([]string{"point"}, suite.tr.Query("point", "", ""))
	suite.Equal([]string{"point", "power"}, suite.tr.Query("po...", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("honey", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("pondy", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("po", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("pow", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("poi", "", ""))
}

func (suite *TreeTrieSuite) TestThreeWordsSharingPrefixABC() {
	suite.tr.Add("point")
	suite.tr.Add("polis")
	suite.tr.Add("power")
	suite.Equal([]string{"point"}, suite.tr.Query("point", "", ""))
	suite.Equal([]string{"polis"}, suite.tr.Query("polis", "", ""))
	suite.Equal([]string{"power"}, suite.tr.Query("power", "", ""))
	suite.Equal([]string{"point", "polis", "power"}, suite.tr.Query("po...", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("honey", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("pondy", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("po", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("pow", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("poi", "", ""))
}

func (suite *TreeTrieSuite) TestThreeWordsSharingPrefixCBA() {
	suite.tr.Add("power")
	suite.tr.Add("polis")
	suite.tr.Add("point")
	suite.Equal([]string{"power"}, suite.tr.Query("power", "", ""))
	suite.Equal([]string{"polis"}, suite.tr.Query("polis", "", ""))
	suite.Equal([]string{"point"}, suite.tr.Query("point", "", ""))
	suite.Equal([]string{"point", "polis", "power"}, suite.tr.Query("po...", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("honey", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("pondy", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("po", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("pow", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("poi", "", ""))
}

func (suite *TreeTrieSuite) TestThreeWordsSharingPrefixACB() {
	suite.tr.Add("point")
	suite.tr.Add("power")
	suite.tr.Add("polis")
	suite.Equal([]string{"point"}, suite.tr.Query("point", "", ""))
	suite.Equal([]string{"power"}, suite.tr.Query("power", "", ""))
	suite.Equal([]string{"polis"}, suite.tr.Query("polis", "", ""))
	suite.Equal([]string{"point", "polis", "power"}, suite.tr.Query("po...", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("honey", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("pondy", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("po", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("pow", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("poi", "", ""))
}

func (suite *TreeTrieSuite) TestOneWordWithInclude() {
	suite.tr.Add("hello")
	suite.Equal([]string{"hello"}, suite.tr.Query("hello", "l", ""))
	suite.Equal([]string{"hello"}, suite.tr.Query("hello", "ll", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("hello", "lll", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("hello", "x", ""))
}

func (suite *TreeTrieSuite) TestTwoWordsSharingPrefixWithInclude() {
	suite.tr.Add("point")
	suite.tr.Add("power")
	suite.Equal([]string{"point"}, suite.tr.Query("point", "op", ""))
	suite.Equal([]string{"power"}, suite.tr.Query("power", "op", ""))
	suite.Equal([]string{"point", "power"}, suite.tr.Query("po...", "op", ""))
	suite.Equal([]string{"power"}, suite.tr.Query("po...", "opw", ""))
}

func (suite *TreeTrieSuite) TestOneWordWithExclude() {
	suite.tr.Add("hello")
	suite.Equal([]string{"hello"}, suite.tr.Query("hello", "", "x"))
	suite.Equal(make([]string, 0), suite.tr.Query("hello", "", "h"))
	suite.Equal(make([]string, 0), suite.tr.Query("hello", "", "xl"))
}

func (suite *TreeTrieSuite) TestTwoWordsWithExclude() {
	suite.tr.Add("point")
	suite.tr.Add("power")
	suite.Equal([]string{"point"}, suite.tr.Query("point", "", "xyz"))
	suite.Equal([]string{"power"}, suite.tr.Query("power", "", "xyz"))
	suite.Equal([]string{"point", "power"}, suite.tr.Query("po...", "", "xyz"))
	suite.Equal([]string{"power"}, suite.tr.Query("po...", "", "xt"))
	suite.Equal([]string{"point"}, suite.tr.Query("po...", "", "xe"))
	suite.Equal(make([]string, 0), suite.tr.Query("po...", "", "xte"))
}

func (suite *TreeTrieSuite) TestTwoWordsContainedInOrder() {
	suite.tr.Add("tex")
	suite.Equal(make([]string, 0), suite.tr.Query("te", "", ""))
	suite.Equal([]string{"tex"}, suite.tr.Query("tex", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("text", "", ""))
	suite.tr.Add("text")
	suite.Equal(make([]string, 0), suite.tr.Query("te", "", ""))
	suite.Equal([]string{"tex"}, suite.tr.Query("tex", "", ""))
	suite.Equal([]string{"text"}, suite.tr.Query("text", "", ""))
}

func (suite *TreeTrieSuite) TestTwoWordsContainedOutOfOrder() {
	suite.tr.Add("text")
	suite.Equal(make([]string, 0), suite.tr.Query("te", "", ""))
	suite.Equal(make([]string, 0), suite.tr.Query("tex", "", ""))
	suite.Equal([]string{"text"}, suite.tr.Query("text", "", ""))
	suite.tr.Add("tex")
	suite.Equal(make([]string, 0), suite.tr.Query("te", "", ""))
	suite.Equal([]string{"tex"}, suite.tr.Query("tex", "", ""))
	suite.Equal([]string{"text"}, suite.tr.Query("text", "", ""))
}

/*******************************/
/* Run tests                   */
/*******************************/
func TestTreeTrieSuite(t *testing.T) {
	suite.Run(t, new(TreeTrieSuite))
}
