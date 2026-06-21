package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type LookupError struct {
	Key string
}

func (e *LookupError) Error() string {
	return fmt.Sprintf("lookup %q: %v", e.Key, ErrNotFound)
}

func (e *LookupError) Unwrap() error {
	return ErrNotFound
}

func lookup(key string) (string, error) {
	return "", &LookupError{Key: key}
}

func main() {
	_, err := lookup("language")
	if errors.Is(err, ErrNotFound) {
		fmt.Println("missing value")
	}

	var lookupErr *LookupError
	if errors.As(err, &lookupErr) {
		fmt.Println("missing key:", lookupErr.Key)
	}
}
