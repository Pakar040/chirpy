package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, 500, "Something went wrong", err)
		return
	}

	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long", nil)
		return
	}

	wordsToClean := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}
	cleanedBody := cleanMessage(params.Body, wordsToClean)

	respondWithJSON(w, 200, returnVals{
		CleanedBody: cleanedBody,
	})
}

func cleanMessage(msg string, wordsToClean map[string]struct{}) string {
	msgWords := strings.Split(msg, " ")
	for i, word := range msgWords {
		if _, ok := wordsToClean[strings.ToLower(word)]; ok {
			msgWords[i] = "****"
		}
	}
	return strings.Join(msgWords, " ")
}
