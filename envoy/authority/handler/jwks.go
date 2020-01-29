package handler

import (
	"crypto/rsa"
	"encoding/json"
	"net/http"

	"github.com/lestrrat-go/jwx/jwk"
)

type JWKS struct {
	PrivateKey *rsa.PrivateKey
}

func (h *JWKS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key, err := jwk.New(h.PrivateKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
