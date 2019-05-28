package main

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
)

const (
	miaChiave = "token"
)


func esempio(miaChiave []byte) (string, error) {
	// crea il token
	token := jwt.New(jwt.SigningMethodHS256)
	// set dell'autorizzazione
	autorizzazioni := make(jwt.MapClaims)
	autorizzazioni["hi"] = "hello"
	autorizzazioni["esempio"] = time.Now().Add(time.Hour * 72).Unix()
	token.Claims = autorizzazioni
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(miaChiave)
	return tokenString, err
}

func Parsa(mioToken string, chiaveMia string) {
	token, err := jwt.Parse(mioToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(chiaveMia), nil
	})

	if err == nil && token.Valid {
		fmt.Println(token.Raw)
		fmt.Println("il tuo token è valido ")
	} else {
		fmt.Println("il token non è valido")
	}
}

func main() {
	tokenCreato, err := esempio([]byte(miaChiave))
	if err != nil {
		fmt.Println("Non è possibile creare il token")
	}
	Parsa(tokenCreato, miaChiave)
}
