package file

import (
	"bufio"
	"fmt"
	"os"
)

type repository struct {
	save    bool
	uuid    int8
	data    map[string]string
	file    *os.File
	writer  *bufio.Writer
	scanner *bufio.Scanner
}

func New(file *os.File, save bool) (*repository, error) {
	writer := bufio.NewWriter(file)
	scanner := bufio.NewScanner(file)

	repo := &repository{
		save:    save,
		data:    make(map[string]string),
		file:    file,
		writer:  writer,
		scanner: scanner,
	}

	err := repo.load()
	if err != nil {
		return nil, fmt.Errorf("load error: %v", err)
	}

	return repo, nil
}
