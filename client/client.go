package client

import (
	"coding_challenge/model"
	"encoding/json"
	"net/http"
)

func GetCryptoResponseFromURL(id string) (map[string]model.Cryptoresponse, error) {
	//Build The URL strings
	var URL string
	if id != "" {
		URL = "https://api.hitbtc.com/api/3/public/ticker" + id
	} else {
		URL = "https://api.hitbtc.com/api/3/public/ticker"
	}
	//We make HTTP request using the Get function
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	var cResp map[string]model.Cryptoresponse
	//Decode the data
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return nil, err
	}
	return cResp, nil
}
