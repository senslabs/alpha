package jwt

import (
	"encoding/json"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/logger"
	"github.com/senslabs/alpha/sens/types"
)

const accessSigningKey = "L7zBtmHybEkjubZfvkAs-3gklypIZGO5WZKZZQuLQ"
const refreshSigningKey = "T8YmLUVa9BR66RbNR5YV-zLDM3cuQMz8gKismkDCw"

func generateToken(subject interface{}, duration time.Duration, signingMethod *jwt.SigningMethodHMAC, signingKey string) (string, error) {
	s, err := json.Marshal(subject)
	if err != nil {
		logger.Error(err)
		return "", errors.FromError(errors.GO_ERROR, err)
	}
	issuedAt := time.Now().UTC()
	claims := &jwt.StandardClaims{
		IssuedAt:  issuedAt.Unix(),
		ExpiresAt: issuedAt.Unix() + int64(duration.Seconds()),
		Issuer:    "senslabs.io",
		Subject:   string(s),
	}
	token := jwt.NewWithClaims(signingMethod, claims)
	return token.SignedString([]byte(signingKey))
}

func GenerateAccessToken(subject interface{}, expiry time.Duration) (string, error) {
	return generateToken(subject, expiry, jwt.SigningMethodHS256, accessSigningKey)
}

func GenerateRefreshToken(subject interface{}) (string, error) {
	return generateToken(subject, 90*24*time.Hour, jwt.SigningMethodHS512, refreshSigningKey)
}

func GenerateTemporaryToken(subject interface{}) (string, error) {
	return generateToken(subject, 15*time.Minute, jwt.SigningMethodHS512, refreshSigningKey)
}

func verifyToken(tokenText string, signingMethod *jwt.SigningMethodHMAC, signingKey string) (map[string]interface{}, error) {
	var m map[string]interface{}
	token, err := jwt.Parse(tokenText, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.Errorf("Method: %v, Unexpected signing method: %v", method, token.Header["alg"])
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		err = types.JsonUnmarshal(claims["sub"].([]byte), &m)
	}
	return m, err
}

func VerifyToken(tokenText string) (map[string]interface{}, error) {
	return verifyToken(tokenText, jwt.SigningMethodHS256, accessSigningKey)
}

func RefreshAccessToken(tokenText string) (string, error) {
	if subject, err := verifyToken(tokenText, jwt.SigningMethodHS512, refreshSigningKey); err != nil {
		return "", err
	} else {
		return GenerateAccessToken(subject, 24*time.Hour)
	}
}