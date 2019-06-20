package junitxml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

var androidTestCase = JUnitTestCase{
	XMLName:   xml.Name{Local: "testcase"},
	ClassName: "com.example.android.testing.espresso.RecyclerViewSample.RecyclerViewSampleTest",
	Name:      "testCase1",
	Time:      "2.005",
}

var iOSTestCase = JUnitTestCase{
	XMLName:   xml.Name{Local: "testcase"},
	ClassName: "AppleUITests",
	Name:      "testCase16()",
	Time:      "32.861647963523865",
}

func TestParseFile(t *testing.T) {
	reports, err := ParseSuites("test/")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(reports))
	suites := reports[0]
	assert.Equal(t, 1, len(suites.TestSuite))
	suite := suites.TestSuite[0]
	assert.Equal(t, 5, suite.Tests)
	assert.Equal(t, 1, suite.Failures)
	assert.Equal(t, 5, len(suite.TestCases))
	assert.Contains(t, suite.TestCases, androidTestCase)
}

func TestParseTestSummariesFile(t *testing.T) {
	reports, err := ParseSuites("test/iOS/")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(reports))
	report := reports[0]
	assert.Equal(t, 1, len(report.TestSuite))
	suite := report.TestSuite[0]
	assert.Equal(t, 20, suite.Tests)
	assert.Equal(t, 1, suite.Failures)
	assert.Equal(t, 20, len(suite.TestCases))
	assert.Contains(t, suite.TestCases, iOSTestCase)
}

func TestDirectoryIsNotExist(t *testing.T) {
	_, err := ParseSuites("directoryName/")
	assert.NotNil(t, err)
}

func TestParseSuite(t *testing.T) {
	report, err := ParseSuite("test/dir1/dir2/dir3/emulator-5554.xml")
	assert.Nil(t, err)
	assert.NotNil(t, report)
	assert.Equal(t, 1, len(report.TestSuite))
	suite := report.TestSuite[0]
	assert.Equal(t, 5, suite.Tests)
	assert.Equal(t, 1, suite.Failures)
	assert.Equal(t, 5, len(suite.TestCases))
	assert.Contains(t, suite.TestCases, androidTestCase)
}

func TestParseRecursive(t *testing.T) {
	reports, err := ParseSuitesRecursive("test/dir1")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(reports))
	suites := reports[0]
	assert.Equal(t, 1, len(suites.TestSuite))
	suite := suites.TestSuite[0]
	assert.Equal(t, 5, suite.Tests)
	assert.Equal(t, 1, suite.Failures)
	assert.Equal(t, 5, len(suite.TestCases))
	assert.Contains(t, suite.TestCases, androidTestCase)
}

func TestParseRecursiveWithTwoFiles(t *testing.T) {
	suites, err := ParseSuitesRecursive("test/")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(suites))
}
