package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type TokenType string

const (
	// TokenTypeAccess -
	TokenTypeAccess TokenType = "chirpy-access"
)

func HashPassword(password string) (string, error) {

	// Convertir le mot de passe en []byte
	passwordBytes := []byte(password)

	// Hasher le mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err // Retourner une chaîne vide et l'erreur
	}

	// Convertir le hash en string et le retourner
	return string(hashedPassword), nil
}

// CheckPasswordHash -
func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func MakeJWT(userID uuid.UUID, tokenSecret string) (string, error) {

	myClaim := jwt.RegisteredClaims{
		Issuer:    string(TokenTypeAccess),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour)),
		Subject:   (userID).String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)

	return token.SignedString([]byte(tokenSecret))

}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	// Définir une structure pour les claims
	claims := jwt.RegisteredClaims{}

	// Parser le token et valider la signature
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		// Vérifier la méthode de signature
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("méthode de signature invalide : %v", token.Header["alg"])
		}
		// Retourner la clé secrète pour la validation
		return []byte(tokenSecret), nil
	})

	// Gérer les erreurs de parsing ou de validation
	if err != nil {
		return uuid.Nil, err
	}

	userIDString, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, err
	}

	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return uuid.Nil, err
	}
	if issuer != string(TokenTypeAccess) {
		return uuid.Nil, errors.New("invalid issuer")
	}

	id, err := uuid.Parse(userIDString)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user ID: %w", err)
	}
	return id, nil
}

func GetBearerToken(headers http.Header) (string, error) {

	// Get the Authorization header
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is missing")
	}

	// Check if the header starts with "Bearer "
	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		return "", errors.New("authorization header is not in bearer token format")
	}

	// Extract and return the token (strip the "Bearer " prefix)
	token := strings.TrimSpace(authHeader[len(bearerPrefix):])
	if token == "" {
		return "", errors.New("bearer token is empty")
	}

	return token, nil
}

func MakeRefreshToken() (string, error) {

	mes_bits := make([]byte, 32)

	_, err := rand.Read(mes_bits)
	if err != nil {
		fmt.Errorf("probleme avec le random: %w", err)
		return "", err
	}

	my_string := hex.EncodeToString(mes_bits)

	return my_string, nil
}

func GetAPIKey(headers http.Header) (string, error) {

	// Get the Authorization header
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is missing")
	}

	println(authHeader)

	// Extraire la clé en découpant la chaîne
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 {
		return "", fmt.Errorf("format d'en-tête invalide")
	}

	println(parts[1])

	// Retourner la clé (la troisième partie)
	return parts[1], nil
}
