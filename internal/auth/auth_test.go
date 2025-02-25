package auth

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func TestMakeJWT(t *testing.T) {
	userID := uuid.New()
	tokenSecret := "supersecretkey"

	token, err := MakeJWT(userID, tokenSecret)
	if err != nil {
		t.Fatalf("MakeJWT a échoué : %v", err)
	}
	if token == "" {
		t.Error("Le token est vide")
	}
}

func TestValidateJWT(t *testing.T) {
	userID := uuid.New()
	tokenSecret := "supersecretkey"

	// Créer un token
	token, err := MakeJWT(userID, tokenSecret)
	if err != nil {
		t.Fatalf("MakeJWT a échoué : %v", err)
	}

	// Valider le token
	parsedUserID, err := ValidateJWT(token, tokenSecret)
	if err != nil {
		t.Fatalf("ValidateJWT a échoué : %v", err)
	}

	// Vérifier que l'ID utilisateur correspond
	if parsedUserID != userID {
		t.Errorf("ID utilisateur incorrect : attendu %v, obtenu %v", userID, parsedUserID)
	}
}

func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedToken string
		expectError   bool
	}{
		{
			name:          "Valid Bearer Token",
			headers:       http.Header{"Authorization": []string{"Bearer valid_token"}},
			expectedToken: "valid_token",
			expectError:   false,
		},
		{
			name:          "Missing Authorization Header",
			headers:       http.Header{},
			expectedToken: "",
			expectError:   true,
		},
		{
			name:          "Malformed Authorization Header",
			headers:       http.Header{"Authorization": []string{"InvalidHeader"}},
			expectedToken: "",
			expectError:   true,
		},
		{
			name:          "Empty Bearer Token",
			headers:       http.Header{"Authorization": []string{"Bearer "}},
			expectedToken: "",
			expectError:   true,
		},
		{
			name:          "Whitespace in Bearer Token",
			headers:       http.Header{"Authorization": []string{"Bearer  token_with_whitespace  "}},
			expectedToken: "token_with_whitespace",
			expectError:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := GetBearerToken(tt.headers)
			if (err != nil) != tt.expectError {
				t.Errorf("GetBearerToken() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if token != tt.expectedToken {
				t.Errorf("GetBearerToken() token = %v, expected %v", token, tt.expectedToken)
			}
		})
	}
}
