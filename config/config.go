package config

import (
	"fmt"
	"io"
	"io/ioutil"
)

const (
	// JSONCFG --
	JSONCFG = iota + 1
)

// ParseConfigFromReader --
func ParseConfigFromReader(em int, r io.Reader, v interface{}) (err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return parseConfig(em, data, v)
}

// ParseConfigFromFile --
func ParseConfigFromFile(em int, name string, v interface{}) (err error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return err
	}
	return parseConfig(em, data, v)
}

func parseConfig(em int, data []byte, v interface{}) (err error) {
	switch em {
	case JSONCFG:
		err = ParseJSONConfig(string(data), v)
	default:
		err = fmt.Errorf("unsport this format config")
	}
	return err
}
