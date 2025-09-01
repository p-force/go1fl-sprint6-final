package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoDetect(fileData []byte) (string, error) {
	input := strings.TrimSpace(string(fileData))
	
	// Проверяем, содержит ли строка только символы Морзе
	isMorse := true
	for _, ch := range input {
		if ch != '.' && ch != '-' && ch != ' ' {
			isMorse = false
			break
		}
	}

	if isMorse {
		// Конвертируем Морзе в текст
		result := morse.ToText(input)
		if result == "" {
			return "", errors.New("failed to convert Morse code to text")
		}
		return result, nil
	} else {
		// Конвертируем текст в Морзе
		result := morse.ToMorse(input)

		if result == "" {
			return "", errors.New("failed to convert text to Morse code")
		}

		return result, nil
	}
}
