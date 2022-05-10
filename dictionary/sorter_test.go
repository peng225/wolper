package dictionary

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

/*******************************/
/* Test set up                 */
/*******************************/
type SorterSuite struct {
	suite.Suite
}

func (suite *SorterSuite) SetupTest() {
}

/*******************************/
/* Expose private methods      */
/*******************************/

/*******************************/
/* Test cases                  */
/*******************************/
func (suite *SorterSuite) TestGetBucketIndex() {
	target := "abcde"
	colorPattern := []Color{GREEN, GREEN, GREEN, GREEN, GREEN}
	suite.Equal(convToBucketIndex(colorPattern), getBucketIndex(target, target))
	colorPattern = []Color{GREEN, GREEN, GREEN, GREEN, GRAY}
	suite.Equal(convToBucketIndex(colorPattern), getBucketIndex(target, "abcdf"))
	colorPattern = []Color{YELLOW, YELLOW, GREEN, GREEN, GRAY}
	suite.Equal(convToBucketIndex(colorPattern), getBucketIndex(target, "bacdf"))
	colorPattern = []Color{GREEN, GRAY, GREEN, GREEN, GREEN}
	suite.Equal(convToBucketIndex(colorPattern), getBucketIndex(target, "aacde"))

	target = "abbcd"
	colorPattern = []Color{YELLOW, GREEN, YELLOW, GREEN, GREEN}
	suite.Equal(convToBucketIndex(colorPattern), getBucketIndex(target, "bbacd"))
	colorPattern = []Color{GRAY, GREEN, GREEN, GREEN, GREEN}
	suite.Equal(convToBucketIndex(colorPattern), getBucketIndex(target, "bbbcd"))

	target = "bbbcd"
	colorPattern = []Color{GRAY, GREEN, GREEN, GREEN, GREEN}
	suite.Equal(convToBucketIndex(colorPattern), getBucketIndex(target, "abbcd"))

	target = "bbbcd"
	colorPattern = []Color{YELLOW, GRAY, GREEN, YELLOW, GREEN}
	suite.Equal(convToBucketIndex(colorPattern), getBucketIndex(target, "acbbd"))

	target = "abc"
	colorPattern = []Color{GREEN, YELLOW, GRAY}
	suite.Equal(convToBucketIndex(colorPattern), getBucketIndex(target, "adb"))

	target = "abcdef"
	colorPattern = []Color{GREEN, GREEN, YELLOW, YELLOW, YELLOW, GRAY}
	suite.Equal(convToBucketIndex(colorPattern), getBucketIndex(target, "abgcde"))
}

func (suite *SorterSuite) TestSortWithEntropy() {
	words := []string{"aaaaa", "aaaaz", "aaaax", "aaaay", "aaxyz"}
	sortWithEntropy(words)
	suite.Equal("aaxyz", words[0])
	suite.Equal("aaaaa", words[1])
	suite.Equal("aaaax", words[2])
	suite.Equal("aaaay", words[3])
	suite.Equal("aaaaz", words[4])
}

/*******************************/
/* Run tests                   */
/*******************************/
func TestSorterSuite(t *testing.T) {
	suite.Run(t, new(SorterSuite))
}
