package handler

import (
	"crypto/rsa"
	"net/http"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

var _ http.Handler = &JWT{}

type JWT struct {
	PrivateKey *rsa.PrivateKey
}

func (j *JWT) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := jwt.New()
	signedToken, err := token.Sign(jwa.RS256, j.PrivateKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(signedToken)
}
