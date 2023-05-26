package main

import (
	"coding_challenge/service"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/currency/all", GetAllSymbolPrices)
	http.HandleFunc("/currency/{symbol}", GetSymbolPrice)

	http.ListenAndServe(":8080", nil)
}

func GetSymbolPrice(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside The Get Symbol Price Flow")
	d := strings.TrimPrefix(r.URL.Path, "/currency/")
	fmt.Println(d)
	res, err := service.FetchCryptoByID(d)
	if err != nil {
		w.Write([]byte("Error in Fetching the Cryptos]"))
	}
	fmt.Fprintf(w, "%v:", res)
}

func GetAllSymbolPrices(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside the Get All Currency Flow")
	res, err := service.FetchAllCrypto()
	if err != nil {
		w.Write([]byte("Error in Fetching the Cryptos]"))
	}
	fmt.Fprintf(w, "%v:", res)
}
