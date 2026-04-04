package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTToken struct {
	secretKey string
}

func NewJWTTOken(secret string) Tokenizer {
	return &JWTToken{
		secretKey: secret,
	}
}

func (j *JWTToken) CreateToken(username string, role string, duration time.Duration) (string, *Payload, error) {
	newPayload, errPayload := NewPayload(username, role, duration)
	if errPayload != nil {
		return "", nil, errPayload
	}
	// create token with claims => Payload{}
	tokenWithClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, newPayload, nil)
	// sign token
	// // signed Claim with secretKey <<save at env variable>>
	signedToken, errSign := tokenWithClaim.SignedString([]byte(j.secretKey))
	return signedToken, newPayload, errSign

}

func (j *JWTToken) VerifyToken(token string) (*Payload, error) {

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(j.secretKey), nil
	}

	tokenWithClaim, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		return nil, err
	}
	// note: assert into *Payload type, but first need to check if assertion process is success
	payload, ok := tokenWithClaim.Claims.(*Payload)
	// check if assertion is failed
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
