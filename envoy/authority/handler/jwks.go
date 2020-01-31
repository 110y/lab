package handler

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jws"
)

type JWKS struct {
	PrivateKey *rsa.PrivateKey
}

func (h *JWKS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ACCESS: JWKS")

	key, err := jwk.New(h.PrivateKey)
	if err != nil {
		fmt.Printf("failed to create jwk: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := key.Set(jws.KeyIDKey, kid); err != nil {
		fmt.Printf("failed to set kid: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	keys := jwk.Set{
		Keys: []jwk.Key{key},
	}

	b, err := json.Marshal(keys)
	if err != nil {
		fmt.Printf("failed to marshal json: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
