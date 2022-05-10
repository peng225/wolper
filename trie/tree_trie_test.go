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
	posExcludeList := make([]string, 5)
	suite.Equal(make([]string, 0), suite.tr.Query("hello", "", "", posExcludeList, false))
	posExcludeList = make([]string, 0)
	suite.Equal(make([]string, 0), suite.tr.Query("", "", "", posExcludeList, false))
}

func (suite *TreeTrieSuite) TestEmptyAdded() {
	posExcludeList := make([]string, 5)
	suite.tr.Add("")
	suite.Equal(make([]string, 0), suite.tr.Query("hello", "", "", posExcludeList, false))
	posExcludeList = make([]string, 0)
	suite.Equal([]string{""}, suite.tr.Query("", "", "", posExcludeList, false))
}

func (suite *TreeTrieSuite) TestOneWord() {
	posExcludeList := make([]string, 5)
	suite.tr.Add("hello")
	suite.Equal([]string{"hello"}, suite.tr.Query("hello", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("honey", "", "", posExcludeList, false))
	posExcludeList = make([]string, 4)
	suite.Equal(make([]string, 0), suite.tr.Query("hell", "", "", posExcludeList, false))
}

func (suite *TreeTrieSuite) TestTwoWords() {
	posExcludeList := make([]string, 5)
	suite.tr.Add("hello")
	suite.tr.Add("point")
	suite.Equal([]string{"hello"}, suite.tr.Query("hello", "", "", posExcludeList, false))
	suite.Equal([]string{"point"}, suite.tr.Query("point", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("honey", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("power", "", "", posExcludeList, false))
	posExcludeList = make([]string, 4)
	suite.Equal(make([]string, 0), suite.tr.Query("hell", "", "", posExcludeList, false))
	posExcludeList = make([]string, 3)
	suite.Equal(make([]string, 0), suite.tr.Query("poi", "", "", posExcludeList, false))
}

func (suite *TreeTrieSuite) TestTwoWordsSharingPrefixInOrder() {
	posExcludeList := make([]string, 5)
	suite.tr.Add("point")
	suite.tr.Add("power")
	suite.Equal([]string{"point"}, suite.tr.Query("point", "", "", posExcludeList, false))
	suite.Equal([]string{"power"}, suite.tr.Query("power", "", "", posExcludeList, false))
	suite.Equal([]string{"point", "power"}, suite.tr.Query("po...", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("honey", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("pondy", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("po", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("pow", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("poi", "", "", posExcludeList, false))
}

func (suite *TreeTrieSuite) TestTwoWordsSharingPrefixOutOfOrder() {
	posExcludeList := make([]string, 5)
	suite.tr.Add("power")
	suite.tr.Add("point")
	suite.Equal([]string{"power"}, suite.tr.Query("power", "", "", posExcludeList, false))
	suite.Equal([]string{"point"}, suite.tr.Query("point", "", "", posExcludeList, false))
	suite.Equal([]string{"point", "power"}, suite.tr.Query("po...", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("honey", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("pondy", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("po", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("pow", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("poi", "", "", posExcludeList, false))
}

func (suite *TreeTrieSuite) TestThreeWordsSharingPrefixABC() {
	posExcludeList := make([]string, 5)
	suite.tr.Add("point")
	suite.tr.Add("polis")
	suite.tr.Add("power")
	suite.Equal([]string{"point"}, suite.tr.Query("point", "", "", posExcludeList, false))
	suite.Equal([]string{"polis"}, suite.tr.Query("polis", "", "", posExcludeList, false))
	suite.Equal([]string{"power"}, suite.tr.Query("power", "", "", posExcludeList, false))
	suite.Equal([]string{"point", "polis", "power"}, suite.tr.Query("po...", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("honey", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("pondy", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("po", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("pow", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("poi", "", "", posExcludeList, false))
}

func (suite *TreeTrieSuite) TestThreeWordsSharingPrefixCBA() {
	posExcludeList := make([]string, 5)
	suite.tr.Add("power")
	suite.tr.Add("polis")
	suite.tr.Add("point")
	suite.Equal([]string{"power"}, suite.tr.Query("power", "", "", posExcludeList, false))
	suite.Equal([]string{"polis"}, suite.tr.Query("polis", "", "", posExcludeList, false))
	suite.Equal([]string{"point"}, suite.tr.Query("point", "", "", posExcludeList, false))
	suite.Equal([]string{"point", "polis", "power"}, suite.tr.Query("po...", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("honey", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("pondy", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("po", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("pow", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("poi", "", "", posExcludeList, false))
}

func (suite *TreeTrieSuite) TestThreeWordsSharingPrefixACB() {
	posExcludeList := make([]string, 5)
	suite.tr.Add("point")
	suite.tr.Add("power")
	suite.tr.Add("polis")
	suite.Equal([]string{"point"}, suite.tr.Query("point", "", "", posExcludeList, false))
	suite.Equal([]string{"power"}, suite.tr.Query("power", "", "", posExcludeList, false))
	suite.Equal([]string{"polis"}, suite.tr.Query("polis", "", "", posExcludeList, false))
	suite.Equal([]string{"point", "polis", "power"}, suite.tr.Query("po...", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("honey", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("pondy", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("po", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("pow", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("poi", "", "", posExcludeList, false))
}

func (suite *TreeTrieSuite) TestOneWordWithInclude() {
	posExcludeList := make([]string, 5)
	suite.tr.Add("hello")
	suite.Equal([]string{"hello"}, suite.tr.Query("hello", "l", "", posExcludeList, false))
	suite.Equal([]string{"hello"}, suite.tr.Query("hello", "ll", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("hello", "lll", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("hello", "x", "", posExcludeList, false))
}

func (suite *TreeTrieSuite) TestTwoWordsSharingPrefixWithInclude() {
	posExcludeList := make([]string, 5)
	suite.tr.Add("point")
	suite.tr.Add("power")
	suite.Equal([]string{"point"}, suite.tr.Query("point", "op", "", posExcludeList, false))
	suite.Equal([]string{"power"}, suite.tr.Query("power", "op", "", posExcludeList, false))
	suite.Equal([]string{"point", "power"}, suite.tr.Query("po...", "op", "", posExcludeList, false))
	suite.Equal([]string{"power"}, suite.tr.Query("po...", "opw", "", posExcludeList, false))
}

func (suite *TreeTrieSuite) TestTwoWordsSharingPrefixWithDuplicatedInclude() {
	posExcludeList := make([]string, 5)
	suite.tr.Add("skill")
	suite.tr.Add("skull")
	suite.Equal([]string{"skill"}, suite.tr.Query("skill", "ll", "", posExcludeList, false))
	suite.Equal([]string{"skull"}, suite.tr.Query("skull", "ll", "", posExcludeList, false))
	suite.Equal([]string{"skill", "skull"}, suite.tr.Query("sk...", "ll", "", posExcludeList, false))
}

func (suite *TreeTrieSuite) TestOneWordWithExclude() {
	posExcludeList := make([]string, 5)
	suite.tr.Add("hello")
	suite.Equal([]string{"hello"}, suite.tr.Query("hello", "", "x", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("hello", "", "h", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("hello", "", "xl", posExcludeList, false))
}

func (suite *TreeTrieSuite) TestTwoWordsWithExclude() {
	posExcludeList := make([]string, 5)
	suite.tr.Add("point")
	suite.tr.Add("power")
	suite.Equal([]string{"point"}, suite.tr.Query("point", "", "xyz", posExcludeList, false))
	suite.Equal([]string{"power"}, suite.tr.Query("power", "", "xyz", posExcludeList, false))
	suite.Equal([]string{"point", "power"}, suite.tr.Query("po...", "", "xyz", posExcludeList, false))
	suite.Equal([]string{"power"}, suite.tr.Query("po...", "", "xt", posExcludeList, false))
	suite.Equal([]string{"point"}, suite.tr.Query("po...", "", "xe", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("po...", "", "xte", posExcludeList, false))
}

func (suite *TreeTrieSuite) TestTwoWordsContainedInOrder() {
	posExcludeList := make([]string, 5)
	suite.tr.Add("tex")
	suite.Equal(make([]string, 0), suite.tr.Query("te", "", "", posExcludeList, false))
	suite.Equal([]string{"tex"}, suite.tr.Query("tex", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("text", "", "", posExcludeList, false))
	suite.tr.Add("text")
	suite.Equal(make([]string, 0), suite.tr.Query("te", "", "", posExcludeList, false))
	suite.Equal([]string{"tex"}, suite.tr.Query("tex", "", "", posExcludeList, false))
	suite.Equal([]string{"text"}, suite.tr.Query("text", "", "", posExcludeList, false))
}

func (suite *TreeTrieSuite) TestTwoWordsContainedOutOfOrder() {
	posExcludeList := make([]string, 5)
	suite.tr.Add("text")
	suite.Equal(make([]string, 0), suite.tr.Query("te", "", "", posExcludeList, false))
	suite.Equal(make([]string, 0), suite.tr.Query("tex", "", "", posExcludeList, false))
	suite.Equal([]string{"text"}, suite.tr.Query("text", "", "", posExcludeList, false))
	suite.tr.Add("tex")
	suite.Equal(make([]string, 0), suite.tr.Query("te", "", "", posExcludeList, false))
	suite.Equal([]string{"tex"}, suite.tr.Query("tex", "", "", posExcludeList, false))
	suite.Equal([]string{"text"}, suite.tr.Query("text", "", "", posExcludeList, false))
}

func (suite *TreeTrieSuite) TestThreeWordsDifferentLength() {
	posExcludeList := make([]string, 5)
	suite.tr.Add("skll")
	suite.tr.Add("skill")
	suite.tr.Add("skull")
	suite.Equal([]string{"skll"}, suite.tr.Query("skll", "ll", "", posExcludeList, false))
	suite.Equal([]string{"skill"}, suite.tr.Query("skill", "ll", "", posExcludeList, false))
	suite.Equal([]string{"skull"}, suite.tr.Query("skull", "ll", "", posExcludeList, false))
	suite.Equal([]string{"skill", "skull"}, suite.tr.Query("sk...", "ll", "", posExcludeList, false))
}

func (suite *TreeTrieSuite) TestThreeWordsUniq() {
	posExcludeList := make([]string, 5)
	suite.tr.Add("space")
	suite.tr.Add("speed")
	suite.tr.Add("split")
	suite.Equal([]string{"space", "split"}, suite.tr.Query("sp...", "", "", posExcludeList, true))
}

func (suite *TreeTrieSuite) TestTwoWordsWithPosExclude() {
	posExcludeList := []string{"", "", "w", "", ""}
	suite.tr.Add("point")
	suite.tr.Add("power")
	suite.Equal([]string{"point"}, suite.tr.Query("po...", "", "", posExcludeList, false))
	posExcludeList = []string{"", "", "i", "", ""}
	suite.Equal([]string{"power"}, suite.tr.Query("po...", "", "", posExcludeList, false))
	posExcludeList = []string{"", "", "z", "", ""}
	suite.Equal([]string{"point", "power"}, suite.tr.Query("po...", "", "", posExcludeList, false))
	posExcludeList = []string{"", "", "iw", "", ""}
	suite.Equal(make([]string, 0), suite.tr.Query("po...", "", "", posExcludeList, false))
}

/*******************************/
/* Run tests                   */
/*******************************/
func TestTreeTrieSuite(t *testing.T) {
	suite.Run(t, new(TreeTrieSuite))
}
