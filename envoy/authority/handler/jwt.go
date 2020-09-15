package handler

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
)

var _ http.Handler = &JWT{}

const (
	// issuer = "authority.authority.svc.cluster.local"
	kid = "c0e21c71-f442-4340-a994-48f648fa88c2"
)

type JWT struct {
	PrivateKey *rsa.PrivateKey
}

func (j *JWT) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := jwt.New()

	host, err := os.Hostname()
	if err != nil {
		fmt.Printf("failed to get host: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := token.Set(jwt.IssuerKey, host); err != nil {
		fmt.Printf("failed to set iss: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := token.Set("foo", "bar"); err != nil {
		fmt.Printf("failed to set foo: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	headers := &jws.StandardHeaders{}
	if err := headers.Set(jws.KeyIDKey, kid); err != nil {
		fmt.Printf("failed to set kid: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := headers.Set(jws.AlgorithmKey, jwa.RS256.String()); err != nil {
		fmt.Printf("failed to set alg: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := headers.Set(jws.TypeKey, "JWT"); err != nil {
		fmt.Printf("failed to set typ: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(token)
	if err != nil {
		fmt.Printf("failed to marshal json: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	signedToken, err := jws.Sign(b, jwa.RS256, j.PrivateKey, jws.WithHeaders(headers))
	if err != nil {
		fmt.Printf("failed to sign jwt: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(signedToken)
}
