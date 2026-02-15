package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ProcessData(data string) (string, error) {
	trimmed := strings.TrimSpace(data)

	if trimmed == "" {
		return "", nil
	}

	isMorse := true
	for _, char := range trimmed {
		if char != '.' && char != '-' && char != ' ' && char != '\n' && char != '\r' && char != '\t' {
			isMorse = false
			break
		}
	}

	if isMorse {
		return morse.ToText(trimmed), nil
	}
	return morse.ToMorse(trimmed), nil
}
