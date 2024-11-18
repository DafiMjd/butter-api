package helper

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func ParseJwt(token string) (*jwt.Token, error) {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return parsed, err
	}

	return parsed, nil
}
