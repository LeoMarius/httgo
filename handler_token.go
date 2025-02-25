package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/LeoMarius/httgo/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Missing or invalid token", err)
		return
	}

	fmt.Printf("Received token for refresh: %s\n", token)

	refresh_token, err := cfg.db.FindToken(r.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Failed to find refresh token", err)
		return
	}

	if refresh_token.RevokedAt.Valid {
		respondWithError(w, http.StatusUnauthorized, "Token has been revoked", nil)
		return
	}

	if refresh_token.ExpiresAt.Before(time.Now()) {
		respondWithError(w, http.StatusUnauthorized, "Token has expired", nil)
		return
	}

	jwtToken, err := auth.MakeJWT(refresh_token.UserID, cfg.secret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Failded to make token", err)
		return
	}

	type response struct {
		Token string `json:"token"`
	}

	respondWithJSON(w, http.StatusOK, response{
		Token: jwtToken,
	})

}

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Missing or invalid token", err)
		return
	}

	refresh_token, err := cfg.db.FindToken(r.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Failed to find refresh token", err)
		return
	}

	if refresh_token.RevokedAt.Valid {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if refresh_token.ExpiresAt.Before(time.Now()) {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err = cfg.db.RevokeRefreshToken(r.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to revoke token", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
