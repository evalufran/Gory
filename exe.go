package main

import (
    "io"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
	"html/template"
)
func queryParamDisplayHandler(res http.ResponseWriter, req *http.Request) {
    io.WriteString(res, "nome: "+req.FormValue("nome"))
    io.WriteString(res, "\nnumero: "+req.FormValue("numero"))
}
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

const (
	miaChiave = "token"
)

func main() {

	tokenCreato, err := esempio([]byte(miaChiave))
	if err != nil {
		fmt.Println("Non è possibile creare il token")

	} else { Parsa(tokenCreato, miaChiave)
		tmpl1 := template.Must(template.ParseFiles("home.html"))
		http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		tmpl1.Execute(w, struct{ Success bool }{true})
	})
		http.HandleFunc("/example", func(res http.ResponseWriter, req *http.Request) {
		queryParamDisplayHandler(res, req)

	})
	
    http.ListenAndServe(":9090", nil)
}
}