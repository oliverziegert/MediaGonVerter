package service

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	error2 "pc-ziegert.de/media_service/common/error"
	"pc-ziegert.de/media_service/common/log"
	"strings"
	"time"
)

// The JWT will contains following information:
// Signing Algorithm: RS256
// Standard:
// iss: "dracoon auth service"
// iat: [now]
// exp: [now+5s]
// User Infos:
// tenantUuid (String): The UUID of the tenant
// customerUuid (String): The UUID of the customer
// customerId (long): The ID of the customer
// userId (long): The ID of the user
// username (String): The user's username
// firstName (String): The user's first name
// lastName (String): The user's last name
// email (Email address): The user's email address
// isProviderCustomerUser (boolean): Flag for an provider customer user
// roles (String[]): List of assigned user roles

type CustomClaims struct {
	Issuer                 string     `json:"iss,omitempty"`
	IssuedAt               *time.Time `json:"iat,omitempty"`
	ExpiresAt              *time.Time `json:"exp,omitempty"`
	TenantUuid             string     `json:"tenantUuid,omitempty"`
	CustomerUuid           string     `json:"customerUuid,omitempty"`
	CustomerId             uint64     `json:"customerId,omitempty"`
	UserId                 uint64     `json:"userId,omitempty"`
	Username               string     `json:"username,omitempty"`
	FirstName              string     `json:"firstName,omitempty"`
	LastName               string     `json:"lastName,omitempty"`
	Email                  string     `json:"email,omitempty"`
	IsProviderCustomerUser bool       `json:"isProviderCustomerUser,omitempty"`
	Roles                  []string   `json:"roles,omitempty"`
}

type JWTService struct {
	publicKey string
	pk        *rsa.PublicKey
}

func NewJWTService(publicKey string) *JWTService {
	publicKey = strings.ReplaceAll(publicKey, "\\n", "\n")
	pk, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		//ToDo
	}
	return &JWTService{
		pk:        pk,
		publicKey: publicKey,
	}
}

func (s *JWTService) ValidateToken(encodedToken string) (*jwt.Token, jwt.MapClaims, *error2.Error) {
	mapClaims := jwt.MapClaims{}
	jwtToken, err := jwt.ParseWithClaims(encodedToken, mapClaims, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodRSA); !isValid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return s.pk, nil
	})
	if err != nil {
		return nil, nil, error2.WrapError(error2.AuthCredentialsInvalid, "JWT token validation vaild", err)
	}
	return jwtToken, mapClaims, nil
}

func (s *JWTService) Claims(mapClaims *jwt.MapClaims) (*CustomClaims, *error2.Error) {
	cc := &CustomClaims{}
	err := mapstructure.Decode(mapClaims, cc)
	if err != nil {
		err := error2.WrapError(error2.ValIdInvalid, "Failed to register a consumer.", err)
		log.Debug(err.StackTrace())
		return nil, err
	}
	return cc, nil
}
