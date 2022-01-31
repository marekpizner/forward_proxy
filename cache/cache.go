package cache

import (
	"errors"
	"fmt"
	"os"
)

type Cache interface {
	add(string, []byte)
	get(string)
}

type InMemorry struct {
	filePath string
}

func fileExist(filepath string) (bool, error) {
	if _, err := os.Stat(filepath); err == nil {
		return true, err
	} else if errors.Is(err, os.ErrNotExist) {
		_, err := os.Create(filepath)
		return true, err
	} else {
		return false, err
	}
}

func NewInMemorry(filepath string) *InMemorry {
	check, err := fileExist(filepath)
	if check == false {
		panic(err)
	}
	return &InMemorry{filePath: filepath}
}

func (c InMemorry) Add(hash string, object []byte) {
	err := os.WriteFile(c.filePath, object, 0644)
	if err != nil {
		fmt.Errorf(err.Error())
	}
}

func (c InMemorry) Get(hash string) {

}
