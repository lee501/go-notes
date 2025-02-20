package dictionary

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error){
	dict, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return dict, nil
}
