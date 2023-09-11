package main

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
)

// CustomClaims is a common type to allow the SignToken method to handle both admin and other user types.
type CustomClaims interface {
	// ToClaims will ensure all claim types implement MapClaims interface that the JWT Library uses.
	ToClaims() jwt.MapClaims
}

var (
	Admin = AdminRole{
		value: "admin",
	}
)

type AdminClaims struct {
	ID   int       `json:"id,omitempty"`
	Role AdminRole `json:"role,omitempty"`
	jwt.MapClaims
}

func (a AdminClaims) ToClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"id":   a.ID,
		"role": a.Role.Value(),
	}
}

type UserClaims struct {
	ID   int      `json:"id,omitempty"`
	Role UserRole `json:"role,omitempty"`
	jwt.MapClaims
}

func (u UserClaims) ToClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"id":   u.ID,
		"role": u.Role.Value(),
	}
}

func SignToken(claims CustomClaims) (string, error) {
	var (
		key []byte
		t   *jwt.Token
	)

	tokenClaims := claims.ToClaims()

	switch claims.(type) {
	case AdminClaims:
		key = []byte(os.Getenv("JWT_ADMIN_SECRET_KEY"))
	case UserClaims:
		key = []byte(os.Getenv("JWT_SECRET_KEY"))
	}

	t = jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	s, err := t.SignedString(key)
	return s, err
}

func VerifyToken(tokenString string, role AuthRoles) (jwt.MapClaims, error) {
	if tokenString == "" {
		return nil, EmptyTokenError{}
	}

	var (
		secret string
		claim  jwt.Claims
	)

	switch role.(type) {
	case AdminRole:
		secret = os.Getenv("JWT_ADMIN_SECRET_KEY")
		claim = &AdminClaims{}
	case UserRole:
		secret = os.Getenv("JWT_SECRET_KEY")
		claim = &UserClaims{}
	}

	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}, jwt.WithValidMethods([]string{"HS256"}))

	if err != nil {
		return nil, JWTParsingError{}
	}

	switch claim.(type) {
	case *AdminClaims:
		if claims, ok := token.Claims.(*AdminClaims); ok && token.Valid {
			return claims.ToClaims(), nil
		} else {
			return nil, UnknownClaimTypeError{}
		}
	case *UserClaims:
		if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
			return claims.ToClaims(), nil
		} else {
			return nil, UnknownClaimTypeError{}
		}
	default:
		return nil, UnknownClaimTypeError{}
	}
}
