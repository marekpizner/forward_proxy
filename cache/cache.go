package cache

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"os"
)

type Cache interface {
	Add(string, []byte)
	Get(string) *bytes.Buffer
}

type InMemorry struct {
	filePath string
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func fileExist(filepath string) error {
	if _, err := os.Stat(filepath); err == nil {
		_, err := os.OpenFile(filepath, os.O_RDWR, 0644)
		return err
	} else if errors.Is(err, os.ErrNotExist) {
		_, err := os.Create(filepath)
		return err
	} else {
		return err
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func NewInMemorry(filepath string) *InMemorry {
	return &InMemorry{filePath: filepath}
}

func (c InMemorry) pathBuilder(hash string) string {
	return c.filePath + "/" + hash
}

func (c InMemorry) getFileDescriptor(hash string, type_of_describtor int) (*os.File, error) {
	filePath := c.pathBuilder(hash)
	err := fileExist(filePath)
	check(err)

	file, err := os.OpenFile(filePath, type_of_describtor, 0644)
	return file, err
}

func (c InMemorry) Add(hash string, object []byte) {
	file, err := c.getFileDescriptor(hash, os.O_WRONLY)
	check(err)

	file.Write(object)
}

func (c InMemorry) Get(hash string) *bytes.Buffer {
	buff := new(bytes.Buffer)
	file, err := c.getFileDescriptor(hash, os.O_RDONLY)
	check(err)

	buff.ReadFrom(file)
	check(err)
	return buff
}
