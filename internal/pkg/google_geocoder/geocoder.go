package google_geocoder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type GoogleGeoCoder struct {
	apiKey       string
	languageCode string
	baseRegion   string
	baseURL      string
}

func NewGoogleGeoCoder(apiKey, languageCode, baseRegion string) *GoogleGeoCoder {
	baseUrl := fmt.Sprintf(
		"https://maps.googleapis.com/maps/api/geocode/json?key=%s&region=%s&language=%s",
		apiKey,
		baseRegion,
		languageCode,
	)

	return &GoogleGeoCoder{
		apiKey:       apiKey,
		languageCode: languageCode,
		baseRegion:   baseRegion,
		baseURL:      baseUrl,
	}
}

func (g *GoogleGeoCoder) addParamsToBaseUrl(params ...string) string {
	finalUrl := g.baseURL
	for i := 0; i < len(params); i += 2 {
		finalUrl += fmt.Sprintf("&%s=%s", params[i], url.QueryEscape(params[i+1]))
	}
	return finalUrl
}

func (g *GoogleGeoCoder) GetGeoByAddress(address string) (GoogleGeoResponseResults, error) {
	finalUrl := g.addParamsToBaseUrl("address", address)

	resp, err := http.Get(finalUrl)
	if err != nil {
		return GoogleGeoResponseResults{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GoogleGeoResponseResults{}, err
	}

	var responseStruct GoogleGeoCoderResponse
	err = json.Unmarshal(body, &responseStruct)
	if err != nil {
		return GoogleGeoResponseResults{}, err
	}
	if responseStruct.Status != "OK" {
		return GoogleGeoResponseResults{}, fmt.Errorf(responseStruct.Status)
	}
	if len(responseStruct.Results) == 0 {
		return GoogleGeoResponseResults{}, fmt.Errorf("empty results in response for address: %s", address)
	}

	return responseStruct.Results[0], nil
}