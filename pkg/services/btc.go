package services

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type BTCService struct {
	EndpointAPI string
}

func NewBTCService(endpoint string) *BTCService {
	return &BTCService{
		EndpointAPI: endpoint,
	}
}

func (s *BTCService) GetPriceBTCInUAH() (float64, error) {
	type Response struct {
		Date string `json:"date"`
		BTC  struct {
			UAH float64 `json:"uah"`
		} `json:"btc"`
	}

	resp, err := http.Get(s.EndpointAPI)
	if err != nil {
		return .0, err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Println(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return .0, err
	}

	response := new(Response)

	if err = json.Unmarshal(body, response); err != nil {
		return .0, err
	}

	return response.BTC.UAH, nil
}
