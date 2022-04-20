package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cristalhq/jwt/v4"
)

func main() {
	key := []byte("12345678901234567890123456789012")
	sign, err := jwt.NewSignerHS(jwt.HS256, key)
	if err != nil {
		panic(err)
	}
	claims := new(jwt.RegisteredClaims)
	claims.Audience = []string{"admin"}
	claims.ID = "12345678901234567890123456789012"
	claims.Issuer = "cristalhq"
	claims.Subject = "12345678901234567890123456789012"
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour))

	builder := jwt.NewBuilder(sign)
	token, err := builder.Build(claims)
	if err != nil {
		panic(err)
	}

	verifier, err := jwt.NewVerifierHS(jwt.HS256, key)
	checkErr(err)
	tokenBytes := token.Bytes()
	newToken, err := jwt.Parse(tokenBytes, verifier)
	checkErr(err)
	err = verifier.Verify(newToken)
	checkErr(err)

	var newClaims jwt.RegisteredClaims
	errClaims := json.Unmarshal(newToken.Claims(), &newClaims)
	checkErr(errClaims)

	errParseClaims := jwt.ParseClaims(tokenBytes, verifier, &newClaims)
	checkErr(errParseClaims)
	fmt.Println(newClaims.IsForAudience("admin"))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
