package main

import (
	"net/http"

	"github.com/LeoMarius/httgo/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerDeleteChirp(w http.ResponseWriter, r *http.Request) {

	chirpID := r.PathValue("id")
	chirp_uuid, err := uuid.Parse(chirpID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to pase chirp", err)
		return
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Missing or invalid token", err)
		return
	}

	userUUID, err := auth.ValidateJWT(token, cfg.secret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Pas authorisé", err)
		return
	}
	chirp, err := cfg.db.GetChirpsByID(r.Context(), chirp_uuid)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "chirp not found", err)
		return
	}

	if chirp.UserID != userUUID {
		respondWithError(w, http.StatusForbidden, "Toto : Pas votre chirp2", err)
		return
	}

	err = cfg.db.DeleteChirpsByID(r.Context(), chirp.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to delete", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
