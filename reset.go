package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {

	cfg.fileserverHits.Store(0)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0"))

	err := cfg.db.DeleteUsers(r.Context())
	if err != nil {
		fmt.Println("couldn't delete users: %w", err)
	}

	err = cfg.db.DeleteAllChirps(r.Context())
	if err != nil {
		fmt.Println("couldn't delete chirps: %w", err)
	}

	err = cfg.db.DeleteAllTokens(r.Context())
	if err != nil {
		fmt.Println("couldn't delete Tokens: %w", err)
	}

	fmt.Println("Database reset successfully!")

}
