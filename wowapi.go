package wowapi

import (
	"encoding/json"
	"fmt"
	"io"
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

// Client returns a RequestFunc that can be used to make requests to the Blizzard API.
// The returned RequestFunc will use the provided API key to authenticate requests.
// ApiClientId and ApiSecret are the client zh_CNid and secret for the Blizzard API.
//
//	region: us | eu | kr | tw | cn
//	lang: en_US | es_MX | pt_BR | en_GB | es_ES | fr_FR | ru_RU | de_DE | pt_PT | it_IT | zh_TW | ko_KR | zh_CN
func Client(ApiClientId string, ApiSecret string, region string, lang string) (RequestFunc, error) {
	acceptedRegion := []string{"us", "eu", "kr", "tw", "cn"}
	acceptedLang := []string{"en_US", "es_MX", "pt_BR", "en_GB", "es_ES", "fr_FR", "ru_RU", "de_DE", "pt_PT", "it_IT", "zh_TW", "ko_KR", "zh_CN"}

	if !contains(acceptedRegion, region) {
		return nil, fmt.Errorf("invalid region: %s", region)
	}

	if !contains(acceptedLang, lang) {
		return nil, fmt.Errorf("invalid lang: %s", lang)
	}

	client := &http.Client{}

	token, err := blizzardToken(ApiClientId, ApiSecret, region)
	if err != nil {
		return nil, err
	}

	var urlStart string
	if region == "cn" {
		urlStart = "https://gateway.battlenet.com.cn"
	} else {
		urlStart = "https://" + region + ".api.blizzard.com"
	}
	urlEnd := "?namespace=profile-" + region + "&locale=" + lang

	return func(url string) ([]byte, error) {
		url = urlStart + url + urlEnd
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

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return bodyText, nil
	}, nil
}

func blizzardToken(ApiClientId string, ApiSecret string, region string) (token BlizzardAPIBearerToken, err error) {
	client := &http.Client{}
	var URL string
	if region == "cn" {
		URL = "https://oauth.battlenet.com.cn/token?grant_type=client_credentials"
	} else {
		URL = "https://oauth.battle.net/token?grant_type=client_credentials"
	}
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

	if resp.StatusCode != http.StatusOK {
		return token, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

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

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
