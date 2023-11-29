package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	var text string = fmt.Sprintf("<!DOCTYPE html><html><head><title>Webserver</title></head><body>HEllO handleSearch </body></html>")

	w.Write([]byte(text))

}

func HandleGeocode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	var text string = fmt.Sprintf("<!DOCTYPE html><html><head><title>Webserver</title></head><body>HEllO handleSearch </body></html>")

	w.Write([]byte(text))

}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start")
	w.Header().Set("Content-Type", "application/json")

	bodyR, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var jsonResult LogPas
	err = json.Unmarshal(bodyR, &jsonResult)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if CheckLogin(jsonResult.Username) && CheckPassword(jsonResult.Password) {
		tokenJWT := JwtCreate()
		response := JwtResponse{}
		response.Body.Token = tokenJWT

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем токен
		if ValidateToken(w, r) == nil {
			return
		}
		next.ServeHTTP(w, r)
	})
}
