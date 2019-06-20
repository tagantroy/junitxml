package junitxml

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ParseSuite(file string) (*JUnitTestSuiteList, error) {
	if filepath.Ext(file) != ".xml" {
		return nil, errors.New("wrong file extension")
	}
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	suites := &JUnitTestSuiteList{}
	err = xml.Unmarshal(bytes, &suites)
	if err != nil {
		return nil, err
	}
	return suites, nil
}

func ParseSuitesRecursive(dir string) ([]JUnitTestSuiteList, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".xml" {
			files = append(files, path)
		}
		return nil
	})
	var suites []JUnitTestSuiteList
	if err != nil {
		return suites, err
	}
	for _, path := range files {
		s, err := ParseSuite(path)
		if err == nil {
			suites = append(suites, *s)
		}
	}
	return suites, nil
}

func ParseSuites(dir string) ([]JUnitTestSuiteList, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var suites []JUnitTestSuiteList
	for _, file := range files {
		suite, err := ParseSuite(dir + file.Name())
		if err == nil {
			suites = append(suites, *suite)
		}
	}
	return suites, nil
}
