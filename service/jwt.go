package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTService interface {
	GenerateToken(name string, admin bool) string
	ValidateToken(tokenString string) (*jwt.Token, error)
	Meta(token string) (meta JwtMeta, err error)
}

type JwtMeta struct {
	Name  string
	Admin bool
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: "MY-TOP-SECRET",
		issuer:    "dv16888.com",
	}
}

func (j *jwtService) GenerateToken(name string, isAdmin bool) string {
	claims := &jwtCustomClaims{
		name,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}

	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpacted signing method %v", t.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}

func (j *jwtService) Meta(token string) (meta JwtMeta, err error) {
	validToken, error := j.ValidateToken(token)
	if error != nil {
		err = error
	}
	if claims, ok := validToken.Claims.(jwt.MapClaims); ok {
		meta = JwtMeta{Name: claims["name"].(string), Admin: claims["admin"].(bool)}
	}
	return meta, err
}
