package wowapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type RequestFunc func(url string) ([]byte, error)

type BlizzardAPIBearerToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func Client(ApiClientId string, ApiSecret string) RequestFunc {
	client := &http.Client{}

	token, err := blizzardToken(ApiClientId, ApiSecret)
	if err != nil {
		log.Fatal(err)
	}

	return func(url string) ([]byte, error) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+token.AccessToken)

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return bodyText, nil
	}
}

func blizzardToken(ApiClientId string, ApiSecret string) (token BlizzardAPIBearerToken, err error) {
	client := &http.Client{}
	URL := "https://eu.battle.net/oauth/token?grant_type=client_credentials"
	v := url.Values{}
	v.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", URL, strings.NewReader(v.Encode()))
	if err != nil {
		return token, err
	}
	req.SetBasicAuth(ApiClientId, ApiSecret)
	resp, err := client.Do(req)
	if err != nil {
		return token, err
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return token, err
	}

	err = json.Unmarshal(bodyText, &token)
	if err != nil {
		return token, err
	}

	// fmt.Println(string(bodyText))
	return token, nil
}
