package goconfig

import (
	"testing"
)

type TestConfig struct {
	String string `json:"string"`
	Bla    string `json:"bla"`
	Number int    `json:"number"`
	False  bool   `json:"false"`
	True   bool   `json:"true"`
}

func TestReadFile(t *testing.T) {
	loc := "test_config.json"
	f, err := readFile(loc)
	fileString := `{
		"string": "string",
		"true": true,
		"false": false,
		"number": 3
	}`
	if err != nil {
		t.Error(err)
	}
	if string(f) == fileString {
		t.Error("File read but output not what expected")
	}
}

func TestParseConfig(t *testing.T) {
	loc := "./test_config.json"
	testConfig := TestConfig{}
	ParseConfig(loc, &testConfig)
	if testConfig.Number != 3 {
		t.Error("Inputs incorrectly parsed")
	}
	if testConfig.False != false {
		t.Error("Inputs incorrectly parsed")
	}
	if testConfig.True != true {
		t.Error("Inputs incorrectly parsed")
	}
	if testConfig.String != "string" {
		t.Error("Inputs incorrectly parsed")
	}
}
