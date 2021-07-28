package token

import (
	"butaforia.io/utils"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"github.com/cristalhq/jwt/v3"
)

func GenerateToken(claims interface{}) string {
	key := []byte(`secret`)
	signer, err := jwt.NewSignerHS(jwt.HS256, key)
	utils.CheckErr(err)

	builder := jwt.NewBuilder(signer)
	token, err := builder.Build(claims)
	utils.CheckErr(err)

	return token.String()
}

func GenerateTokenID() string {
	s := []byte("random-unique-string")
	h := sha1.New()
	h.Write(s)
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func ParseTokenClaims(tokenStr string) *jwt.StandardClaims {
	key := []byte(`secret`)
	var claims jwt.StandardClaims

	verifier, err := jwt.NewVerifierHS(jwt.HS256, key)
	utils.CheckErr(err)

	token, err := jwt.ParseAndVerifyString(tokenStr, verifier)
	utils.CheckErr(err)

	errClaims := json.Unmarshal(token.RawClaims(), &claims)
	utils.CheckErr(errClaims)

	return &claims
}
