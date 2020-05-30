package util

import (
	"io/ioutil"
	"os"
)

func LoadFile(path string) ([]byte, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	return ioutil.ReadFile(path)
}
