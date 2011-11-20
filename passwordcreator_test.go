package main

import (
	"testing"
	"os"
)

func TestDigitsCharset(t *testing.T) {
	c, e := NewCreator(os.Stdout, false, false, true, false, "")

	testDigits := "0123456789"
	if e != nil || c.characters != testDigits {
		t.Errorf("Characters not distinct.\nExpected \"%s\", but got \"%s\"", testDigits, c.characters)
	}
}

func TestSomeChars(t *testing.T) {
	c, err := NewCreator(os.Stdout, true, false, true, false, ",.-_")

	testCharacters := "abcdefghijklmnopqrstuvwxyz0123456789,.-_"

	if err != nil || c.characters !=  testCharacters {
		t.Errorf("Characters not distinct.\nExpected \"%s\", but got \"%s\"", testCharacters, c.characters)
	}
}
