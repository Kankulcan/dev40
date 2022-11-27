package main

import "errors"

func main() {
}

type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

func (d Dictionary) search(word string) (string,error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}