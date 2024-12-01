package helpers

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Danielyilma/Food_Recipes/services/models"
	"github.com/golang-jwt/jwt"
	"github.com/square/go-jose/v3"
)

var KeyID string = "eec94410-3fe6-43a1-a06c-0630e1509efe"

func loadPrivateKey(path string) (*rsa.PrivateKey, error) {
	privateKeyBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(privateKeyBytes)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func CreateJwt(user models.User) (string, error) {
	privateKey, err := loadPrivateKey("../../rsa_key")
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"https://hasura.io/jwt/claims": map[string]interface{}{
			"x-hasura-allowed-roles": []string{"editor", "user", "mod"},
			"x-hasura-default-role":  "user",
			"x-hasura-user-id":       strconv.Itoa(user.Id),
			"x-hasura-org-id":        "123",
			"x-hasura-custom":        "custom-value",
		},
		"x-hasura-role": "user",
		"user_id":       user.Id,
		"iat":           time.Now().Unix(),
		"exp":           time.Now().Add(time.Hour * 24).Unix(),
	})

	token.Header["Kid"] = KeyID

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetJWK() (jose.JSONWebKeySet, error) {
	publicKeyBytes, err := os.ReadFile("../../rsa_key.pub")
	if err != nil {
		return jose.JSONWebKeySet{}, err
	}

	block, _ := pem.Decode(publicKeyBytes)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Print(block)
		return jose.JSONWebKeySet{}, err
	}
	// Parse the public key to a usable object
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return jose.JSONWebKeySet{}, err
	}

	jwk := jose.JSONWebKey{
		Key:       publicKey,
		KeyID:     KeyID,
		Algorithm: "RS256",
		Use:       "sig",
	}

	jwkSet := jose.JSONWebKeySet{
		Keys: []jose.JSONWebKey{jwk},
	}

	return jwkSet, nil
}
