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
	suite.Equal(make([]string, 0), suite.tr.Query("hello", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("", "", "", false))
}

func (suite *TreeTrieSuite) TestEmptyAdded() {
	suite.tr.Add("")
	suite.Equal(make([]string, 0), suite.tr.Query("hello", "", "", false))
	suite.Equal([]string{""}, suite.tr.Query("", "", "", false))
}

func (suite *TreeTrieSuite) TestOneWord() {
	suite.tr.Add("hello")
	suite.Equal([]string{"hello"}, suite.tr.Query("hello", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("honey", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("hell", "", "", false))
}

func (suite *TreeTrieSuite) TestTwoWords() {
	suite.tr.Add("hello")
	suite.tr.Add("point")
	suite.Equal([]string{"hello"}, suite.tr.Query("hello", "", "", false))
	suite.Equal([]string{"point"}, suite.tr.Query("point", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("honey", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("power", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("hell", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("poi", "", "", false))
}

func (suite *TreeTrieSuite) TestTwoWordsSharingPrefixInOrder() {
	suite.tr.Add("point")
	suite.tr.Add("power")
	suite.Equal([]string{"point"}, suite.tr.Query("point", "", "", false))
	suite.Equal([]string{"power"}, suite.tr.Query("power", "", "", false))
	suite.Equal([]string{"point", "power"}, suite.tr.Query("po...", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("honey", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("pondy", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("po", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("pow", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("poi", "", "", false))
}

func (suite *TreeTrieSuite) TestTwoWordsSharingPrefixOutOfOrder() {
	suite.tr.Add("power")
	suite.tr.Add("point")
	suite.Equal([]string{"power"}, suite.tr.Query("power", "", "", false))
	suite.Equal([]string{"point"}, suite.tr.Query("point", "", "", false))
	suite.Equal([]string{"point", "power"}, suite.tr.Query("po...", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("honey", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("pondy", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("po", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("pow", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("poi", "", "", false))
}

func (suite *TreeTrieSuite) TestThreeWordsSharingPrefixABC() {
	suite.tr.Add("point")
	suite.tr.Add("polis")
	suite.tr.Add("power")
	suite.Equal([]string{"point"}, suite.tr.Query("point", "", "", false))
	suite.Equal([]string{"polis"}, suite.tr.Query("polis", "", "", false))
	suite.Equal([]string{"power"}, suite.tr.Query("power", "", "", false))
	suite.Equal([]string{"point", "polis", "power"}, suite.tr.Query("po...", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("honey", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("pondy", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("po", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("pow", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("poi", "", "", false))
}

func (suite *TreeTrieSuite) TestThreeWordsSharingPrefixCBA() {
	suite.tr.Add("power")
	suite.tr.Add("polis")
	suite.tr.Add("point")
	suite.Equal([]string{"power"}, suite.tr.Query("power", "", "", false))
	suite.Equal([]string{"polis"}, suite.tr.Query("polis", "", "", false))
	suite.Equal([]string{"point"}, suite.tr.Query("point", "", "", false))
	suite.Equal([]string{"point", "polis", "power"}, suite.tr.Query("po...", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("honey", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("pondy", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("po", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("pow", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("poi", "", "", false))
}

func (suite *TreeTrieSuite) TestThreeWordsSharingPrefixACB() {
	suite.tr.Add("point")
	suite.tr.Add("power")
	suite.tr.Add("polis")
	suite.Equal([]string{"point"}, suite.tr.Query("point", "", "", false))
	suite.Equal([]string{"power"}, suite.tr.Query("power", "", "", false))
	suite.Equal([]string{"polis"}, suite.tr.Query("polis", "", "", false))
	suite.Equal([]string{"point", "polis", "power"}, suite.tr.Query("po...", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("honey", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("pondy", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("po", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("pow", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("poi", "", "", false))
}

func (suite *TreeTrieSuite) TestOneWordWithInclude() {
	suite.tr.Add("hello")
	suite.Equal([]string{"hello"}, suite.tr.Query("hello", "l", "", false))
	suite.Equal([]string{"hello"}, suite.tr.Query("hello", "ll", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("hello", "lll", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("hello", "x", "", false))
}

func (suite *TreeTrieSuite) TestTwoWordsSharingPrefixWithInclude() {
	suite.tr.Add("point")
	suite.tr.Add("power")
	suite.Equal([]string{"point"}, suite.tr.Query("point", "op", "", false))
	suite.Equal([]string{"power"}, suite.tr.Query("power", "op", "", false))
	suite.Equal([]string{"point", "power"}, suite.tr.Query("po...", "op", "", false))
	suite.Equal([]string{"power"}, suite.tr.Query("po...", "opw", "", false))
}

func (suite *TreeTrieSuite) TestTwoWordsSharingPrefixWithDuplicatedInclude() {
	suite.tr.Add("skill")
	suite.tr.Add("skull")
	suite.Equal([]string{"skill"}, suite.tr.Query("skill", "ll", "", false))
	suite.Equal([]string{"skull"}, suite.tr.Query("skull", "ll", "", false))
	suite.Equal([]string{"skill", "skull"}, suite.tr.Query("sk...", "ll", "", false))
}

func (suite *TreeTrieSuite) TestOneWordWithExclude() {
	suite.tr.Add("hello")
	suite.Equal([]string{"hello"}, suite.tr.Query("hello", "", "x", false))
	suite.Equal(make([]string, 0), suite.tr.Query("hello", "", "h", false))
	suite.Equal(make([]string, 0), suite.tr.Query("hello", "", "xl", false))
}

func (suite *TreeTrieSuite) TestTwoWordsWithExclude() {
	suite.tr.Add("point")
	suite.tr.Add("power")
	suite.Equal([]string{"point"}, suite.tr.Query("point", "", "xyz", false))
	suite.Equal([]string{"power"}, suite.tr.Query("power", "", "xyz", false))
	suite.Equal([]string{"point", "power"}, suite.tr.Query("po...", "", "xyz", false))
	suite.Equal([]string{"power"}, suite.tr.Query("po...", "", "xt", false))
	suite.Equal([]string{"point"}, suite.tr.Query("po...", "", "xe", false))
	suite.Equal(make([]string, 0), suite.tr.Query("po...", "", "xte", false))
}

func (suite *TreeTrieSuite) TestTwoWordsContainedInOrder() {
	suite.tr.Add("tex")
	suite.Equal(make([]string, 0), suite.tr.Query("te", "", "", false))
	suite.Equal([]string{"tex"}, suite.tr.Query("tex", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("text", "", "", false))
	suite.tr.Add("text")
	suite.Equal(make([]string, 0), suite.tr.Query("te", "", "", false))
	suite.Equal([]string{"tex"}, suite.tr.Query("tex", "", "", false))
	suite.Equal([]string{"text"}, suite.tr.Query("text", "", "", false))
}

func (suite *TreeTrieSuite) TestTwoWordsContainedOutOfOrder() {
	suite.tr.Add("text")
	suite.Equal(make([]string, 0), suite.tr.Query("te", "", "", false))
	suite.Equal(make([]string, 0), suite.tr.Query("tex", "", "", false))
	suite.Equal([]string{"text"}, suite.tr.Query("text", "", "", false))
	suite.tr.Add("tex")
	suite.Equal(make([]string, 0), suite.tr.Query("te", "", "", false))
	suite.Equal([]string{"tex"}, suite.tr.Query("tex", "", "", false))
	suite.Equal([]string{"text"}, suite.tr.Query("text", "", "", false))
}

func (suite *TreeTrieSuite) TestThreeWordsDifferentLength() {
	suite.tr.Add("skll")
	suite.tr.Add("skill")
	suite.tr.Add("skull")
	suite.Equal([]string{"skll"}, suite.tr.Query("skll", "ll", "", false))
	suite.Equal([]string{"skill"}, suite.tr.Query("skill", "ll", "", false))
	suite.Equal([]string{"skull"}, suite.tr.Query("skull", "ll", "", false))
	suite.Equal([]string{"skill", "skull"}, suite.tr.Query("sk...", "ll", "", false))
}

func (suite *TreeTrieSuite) TestThreeWordsUniq() {
	suite.tr.Add("space")
	suite.tr.Add("speed")
	suite.tr.Add("split")
	suite.Equal([]string{"space", "split"}, suite.tr.Query("sp...", "", "", true))
}

/*******************************/
/* Run tests                   */
/*******************************/
func TestTreeTrieSuite(t *testing.T) {
	suite.Run(t, new(TreeTrieSuite))
}
