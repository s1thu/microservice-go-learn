package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	jwtSecret []byte
}

var (
	demoEmail    = "sithuwin@example.com"
	demoPassword = "password123"
	demoUserId   = "1"
)

func NewAuthService(jwtSecret []byte) *AuthService {
	return &AuthService{jwtSecret: jwtSecret}
}

type Claims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func (s *AuthService) Login(email, password string) (string, error) {
	if email != demoEmail || password != demoPassword {
		return "", errors.New("invalid credentials")
	}

	// 3️⃣ Create token
	claims := Claims{
		UserId: demoUserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 4️⃣ Sign token with secret
	signedToken, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (s *AuthService) ParseToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims := token.Claims.(*Claims)
	return claims.UserId, nil
}
