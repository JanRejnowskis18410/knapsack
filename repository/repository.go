package repository

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Repository struct {
	Filename string
	Capacity int
	Items    []Item
	Size     int
}

type Item struct {
	Id     int
	Value  int
	Weight int
}

// New returns a created instance of Repository with data
func New(filename string) (*Repository, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	repository := Repository{}
	err = repository.parse(file)
	if err != nil {
		return nil, err
	}
	return &repository, nil
}

// parse makes parsing of provided file, that is in a format of:
// first line is a capacity of a knapsack
// next lines are indicating items (each line is an item)
// on each line there is a value and a weight of an item separated by space
// file ends with an empty line
func (repository *Repository) parse(file *os.File) error {
	scanner := bufio.NewScanner(file)
	isFirstLine := true
	var items []Item
	currentId := 1
	for scanner.Scan() {
		line := scanner.Text()
		if isFirstLine {
			isFirstLine = false
			capacity, err := strconv.Atoi(line)
			if err != nil {
				return err
			}
			repository.Capacity = capacity
		} else {
			s := strings.Split(line, " ")
			value, err := strconv.Atoi(s[0])
			if err != nil {
				return err
			}
			weight, err := strconv.Atoi(s[1])
			if err != nil {
				return err
			}
			item := Item{Value: value, Weight: weight, Id: currentId}
			currentId++
			items = append(items, item)
		}
	}
	repository.Items = items
	repository.Size = len(items)
	return nil
}
