package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/LeoMarius/httgo/internal/auth"
	"github.com/LeoMarius/httgo/internal/database"
	"github.com/google/uuid"
)

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerChirps(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body   string    `json:"body"`
		UserId uuid.UUID `json:"user_id"`
	}

	// Décoder le corps de la requête
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
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

	my_chirp := database.CreateChirpParams{
		Body:   params.Body,
		UserID: userUUID,
	}
	// print(r.Body)

	chirp, err := cfg.db.CreateChirp(r.Context(), my_chirp)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create chirp", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.CreatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})

}

func (cfg *apiConfig) handlerAllChirps(w http.ResponseWriter, r *http.Request) {
	// Get the author_id from query parameters, not request body
	authorIDStr := r.URL.Query().Get("author_id")

	sort := r.URL.Query().Get("sort")

	var err error

	var dbChiprs []database.Chirp

	if authorIDStr != "" {
		// Only try to parse the UUID if author_id was provided
		authorID, err := uuid.Parse(authorIDStr)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid author_id format", err)
			return
		}
		if sort == "desc" {
			dbChiprs, err = cfg.db.GetChirpsByUserIDDESC(r.Context(), authorID)
		} else {
			dbChiprs, err = cfg.db.GetChirpsByUserIDASC(r.Context(), authorID)
		}
		// If we have a valid author_id, get chirps for that author

	} else {
		// If no author_id was provided, get all chirps
		if sort == "desc" {
			dbChiprs, err = cfg.db.GetAllChirpsDESC(r.Context())
		} else {
			dbChiprs, err = cfg.db.GetAllChirpsASC(r.Context())
		}

	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps", err)
		return
	}

	var chirps []Chirp
	for _, dbChirp := range dbChiprs {
		chirps = append(chirps, Chirp(dbChirp)) // Conversion et ajout au slice
	}

	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handlerGetOneChirp(w http.ResponseWriter, r *http.Request) {

	chirpID := r.PathValue("id")

	print(chirpID)

	chirp_uuid, err := uuid.Parse(chirpID)
	if err != nil {
		// Gérer l'erreur si la conversion échoue
		respondWithError(w, http.StatusNotFound, "Invalid UUID", err)
		return
	}

	chirp, err := cfg.db.GetChirpsByID(r.Context(), chirp_uuid)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp not found", err)
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.CreatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}
