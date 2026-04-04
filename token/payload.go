package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username, role string, duration time.Duration) (*Payload, error) {
	// for Session ID
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	// new payload
	p := &Payload{
		ID:        tokenId,
		Username:  username,
		Role:      role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return p, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func (p *Payload) GetAudience() (jwt.ClaimStrings, error) {
	// for now, because I use single backend service then I dont need this in concrete implementation
	// this is used when I use multiple service, and restrict which SERVICE (aud) can use a token
	// so i only need to SATISFY the interface
	return nil, nil
}

func (p *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	// for now, because I use single backend service then I dont need this in concrete implementation
	// so i only need to SATISFY the interface
	return nil, nil
}

func (p *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	// for now, because I use single backend service then I dont need this in concrete implementation
	// so i only need to SATISFY the interface
	return nil, nil
}
func (p *Payload) GetIssuer() (string, error) {
	// for now, because I use single backend service then I dont need this in concrete implementation
	// so i only need to SATISFY the interface
	return "", nil
}

func (p *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, nil
}
func (p *Payload) GetSubject() (string, error) {
	return "", nil
}
