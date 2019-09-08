package goconfig

import (
	"encoding/json"
	"io/ioutil"
)

// ParseConfig ..
func ParseConfig(file string, conf interface{}) (err error) {
	b, err := readFile(file)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &conf)
	if err != nil {
		return
	}
	return
}

func readFile(file string) ([]byte, error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return f, nil
}
