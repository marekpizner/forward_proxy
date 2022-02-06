package proxy

import (
	"bufio"
	"os"
)

type Blacklist struct {
	list []string
}

func readBlacklistFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func NewBlacklist(filepath string) *Blacklist {
	list, _ := readBlacklistFile(filepath)
	return &Blacklist{list: list}
}

func (b Blacklist) Check(str string) bool {
	for _, s := range b.list {
		if s == str {
			return true
		}
	}
	return false
}
