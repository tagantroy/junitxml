package junitxml

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"path/filepath"
)

func ParseSuite(file string) (*JUnitTestSuite, error) {
	if filepath.Ext(file) != ".xml" {
		return nil, errors.New("wrong file extension")
	}
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	suite := &JUnitTestSuite{}
	err = xml.Unmarshal(bytes, &suite)
	if err != nil {
		return nil, err
	}
	return suite, nil
}

func ParseSuites(dir string) ([]JUnitTestSuite, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var suites []JUnitTestSuite
	for _, file := range files {
		suite, err := ParseSuite(dir + file.Name())
		if err == nil {
			suites = append(suites, *suite)
		}
	}
	return suites, nil
}
