package junitxml

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

var firstTestCase = JUnitTestCase{
	XMLName:   xml.Name{Local: "testcase"},
	ClassName: "com.example.android.testing.espresso.RecyclerViewSample.RecyclerViewSampleTest",
	Name:      "testCase1",
	Time:      "2.005",
}

func TestParseFile(t *testing.T) {
	suites, err := ParseSuites("test/")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(suites))
	suite := suites[0]
	assert.Equal(t, 5, suite.Tests)
	assert.Equal(t, 1, suite.Failures)
	assert.Equal(t, 5, len(suite.TestCases))
	assert.Contains(t, suite.TestCases, firstTestCase)
}

func TestDirectoryIsNotExist(t *testing.T) {
	_, err := ParseSuites("directoryName/")
	assert.NotNil(t, err)
}

func TestParseSuite(t *testing.T) {
	suite, err := ParseSuite("test/dir1/dir2/dir3/emulator-5554.xml")
	assert.Nil(t, err)
	assert.NotNil(t, suite)
	assert.Equal(t, 5, suite.Tests)
	assert.Equal(t, 1, suite.Failures)
	assert.Equal(t, 5, len(suite.TestCases))
	assert.Contains(t, suite.TestCases, firstTestCase)
}

func TestParseRecursive(t *testing.T) {
	suites, err := ParseSuitesRecursive("test/dir1")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(suites))
	suite := suites[0]
	assert.Equal(t, 5, suite.Tests)
	assert.Equal(t, 1, suite.Failures)
	assert.Equal(t, 5, len(suite.TestCases))
	assert.Contains(t, suite.TestCases, firstTestCase)
}

func TestParseRecursiveWithTwoFiles(t *testing.T) {
	suites, err := ParseSuitesRecursive("test/")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(suites))
}
