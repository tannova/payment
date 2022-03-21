package token

import (
	"crypto/rsa"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	config "gitlab.com/greyhole/night-kit/pkg/carbon"
	"go.uber.org/zap"
)

type Token interface {
	Create(mexID int64) (jti string, token string, err error)
	Parse(token string) (*jwt.Token, *jwt.StandardClaims, error)
}

type token struct {
	logger *zap.Logger

	expiresTime     time.Duration
	privateKey      *rsa.PrivateKey
	issuer          string
	defaultAudience string
}

func New(logger *zap.Logger, jwtSigning *config.JwtSigning) (Token, error) {
	expiresTime := time.Duration(jwtSigning.ExpiresTime)
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(jwtSigning.PrivateKey))
	if err != nil {
		return nil, err
	}

	return &token{
		logger:          logger,
		expiresTime:     expiresTime,
		privateKey:      privateKey,
		issuer:          jwtSigning.Issuer,
		defaultAudience: jwtSigning.Audience,
	}, nil
}

func (t *token) Create(mexID int64) (jti string, token string, err error) {
	id := uuid.New().String()
	iat := time.Now().UTC()
	exp := iat.Add(t.expiresTime * time.Second)

	tk := &jwt.StandardClaims{
		Audience:  t.defaultAudience,
		ExpiresAt: exp.Unix(),
		Id:        id,
		IssuedAt:  iat.Unix(),
		Issuer:    t.issuer,
		NotBefore: iat.Unix(),
		Subject:   strconv.FormatInt(mexID, 10),
	}

	token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, tk).SignedString(t.privateKey)

	return id, token, err
}

func (t *token) Parse(token string) (*jwt.Token, *jwt.StandardClaims, error) {
	claims := &jwt.StandardClaims{}
	tk, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return t.privateKey.Public(), nil
	})

	return tk, claims, err
}
