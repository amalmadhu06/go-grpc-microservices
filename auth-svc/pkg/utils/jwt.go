package utils

import (
	"errors"
	"github.com/amalmadhu06/go-grpc-microservices/auth-svc/pkg/models"
	"github.com/golang-jwt/jwt"
	"time"
)

type JWTWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type jwtClaims struct {
	jwt.StandardClaims
	Id    int64
	Email string
	Role  string
}

func (w *JWTWrapper) GenerateToken(user models.User, role string) (signedToken string, err error) {
	claims := &jwtClaims{
		Id:    user.Id,
		Email: user.Email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			//Audience:  "",
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours)).Unix(),
			//Id:        "",
			//IssuedAt:  0,
			Issuer: w.Issuer,
			//NotBefore: 0,
			//Subject:   "",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(w.SecretKey))

	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (w *JWTWrapper) ValidateToken(singedToken string, role string) (claims *jwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		singedToken,
		&jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}
	if claims.Role != role {
		return nil, errors.New("role not matching")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("JWT is expired")
	}
	return claims, nil
}
