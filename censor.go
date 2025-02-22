package main

import (
	"errors"
	"strings"
)

func sensorBadWords(msg string) (string, error) {

	if msg == "" {
		return "", errors.New("message vide")
	}

	banned_words := []string{"kerfuffle", "sharbert", "fornax"}

	msg_array := strings.Split(msg, " ")

	for i, word := range msg_array {
		for _, banned_word := range banned_words {
			lower_work := strings.ToLower(word)
			if lower_work == banned_word {
				msg_array[i] = "****"
			}
		}
	}

	msg2 := strings.Join(msg_array, " ")

	return msg2, nil
}
