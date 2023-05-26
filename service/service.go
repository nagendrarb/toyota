package service

import (
	"coding_challenge/client"
	"coding_challenge/model"
)

// Fetch is exported ...
func FetchCryptoByID(crypto string) (string, error) {
	//Call the External Client
	resp, err := client.GetCryptoResponseFromURL(crypto)
	if err != nil {
		return "", err
	}
	value := resp[crypto]
	//Invoke the text output function & return it with nil as the error value
	return value.TextOutput(), nil
}

func FetchAllCrypto() ([]model.Cryptoresponse, error) {
	cResp, err := client.GetCryptoResponseFromURL("")
	if err != nil {
		return []model.Cryptoresponse{}, err
	}
	result := make([]model.Cryptoresponse, 0)

	for key, _ := range cResp {
		price := cResp[key]

		cry := model.Cryptoresponse{
			ID:       key[:3],
			FullName: key,
			Ask:      price.Ask,
			Bid:      price.Bid,
			Last:     price.Last,
			Open:     price.Open,
			Low:      price.Low,
			High:     price.High,
		}

		result = append(result, cry)
	}

	return result, nil
}
