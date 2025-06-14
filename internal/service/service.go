package service

import (
	"errors"
	"strings"

	"go1fl-sprint6-final-tpl/pkg/morse"
)

const (
	morseCharacters = ".- "
	latinCharacters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func Convert(s string) (string, error) {
	if s == "" {
		return "", errors.New("empty string passed")
	}

	isMorse := true
	for _, val := range s {
		if strings.ContainsRune(latinCharacters, val) {
			return "", errors.New("the line contains latin characters")
		}
		if !strings.ContainsRune(morseCharacters, val) {
			isMorse = false
		}
	}

	if isMorse {
		return morse.ToText(s), nil
	}
	return morse.ToMorse(s), nil
}