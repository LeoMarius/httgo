package main

import (
	"encoding/json"
	"net/http"

	"github.com/LeoMarius/httgo/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerPolkaHooks(w http.ResponseWriter, r *http.Request) {

	type user struct {
		UserId string `json:"user_id"`
	}

	type parameters struct {
		Event string `json:"event"`
		Data  user   `json:"data"`
	}

	// Décoder le corps de la requête
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Bad API KEY parameters", err)
		return
	}

	if apiKey != cfg.polkaKey {
		respondWithError(w, http.StatusUnauthorized, "Clee API non valide", err)
		return
	}

	if params.Event != "user.upgraded" {
		// On ne prends pas en considération cette info mais on envoi un succès
		w.WriteHeader(http.StatusNoContent)
		return
	}

	userUUID, err := uuid.Parse(params.Data.UserId)
	if err != nil {
		// Gérer l'erreur si la conversion échoue
		respondWithError(w, http.StatusNotFound, "Invalid UUID", err)
		return
	}

	err = cfg.db.UpgradeRed(r.Context(), userUUID)
	if err != nil {
		// Gérer l'erreur si la conversion échoue
		respondWithError(w, http.StatusNotFound, "user not found", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
