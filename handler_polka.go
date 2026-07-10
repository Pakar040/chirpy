package main

import (
	"encoding/json"
	"net/http"

	"github.com/Pakar040/chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerPolkaWebhook(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID uuid.UUID `json:"user_id"`
		} `json:"data"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, 500, "Error decoding request", err)
		return
	}

	if params.Event != "user.upgraded" {
		w.WriteHeader(204)
		return
	}

	_, err := cfg.db.UpdateUserIsChirpyRed(r.Context(), database.UpdateUserIsChirpyRedParams{
		ID:          params.Data.UserID,
		IsChirpyRed: true,
	})
	if err != nil {
		respondWithError(w, 404, "User not found", err)
		return
	}

	respondWithJSON(w, 204, struct{}{})
}
