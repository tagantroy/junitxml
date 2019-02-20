package junitxml

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	suites, err := ParseSuites("test/")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(suites))
	suite := suites[0]
	assert.Equal(t, 5, suite.Tests)
	assert.Equal(t, 1, suite.Failures)
	assert.Equal(t, 5, len(suite.TestCases))
	assert.Contains(t, suite.TestCases, JUnitTestCase{
		XMLName:   xml.Name{Local: "testcase"},
		ClassName: "com.example.android.testing.espresso.RecyclerViewSample.RecyclerViewSampleTest",
		Name:      "testCase1",
		Time:      "2.005",
	})
}

func TestDirectoryIsNotExist(t *testing.T) {
	_, err := ParseSuites("directoryName/")
	assert.NotNil(t, err)
}
