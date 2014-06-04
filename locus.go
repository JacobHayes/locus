package locus

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const cityUrl string = `http://api.ipinfodb.com/v3/ip-city/`
const countryUrl string = `http://api.ipinfodb.com/v3/ip-country/`

type Location struct {
	StatusCode    string `json: "statusCode"`
	StatusMessage string `json: "statusMessage"`
	IpAddress     string `json: "ipAddress"`
	CountryCode   string `json: "countryCode"`
	CountryName   string `json: "countryName"`
	RegionName    string `json: "regionName"`
	CityName      string `json: "cityName"`
	ZipCode       string `json: "zipCode"`
	Latitude      string `json: "latitude"`
	Longitude     string `json: "longitude"`
	TimeZone      string `json: "timeZone"`
}

func locationJson(url string) (*[]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}

func location(url string) (*Location, error) {
	var location *Location

	raw_json, err := locationJson(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*raw_json, location)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func requestUrl(base_url string, key string, ip string) string {
	return strings.Join([]string{base_url, `?key=`, key, `&ip=`, ip, `&format=json`}, ``)
}

// Public API

func CityLocationJson(key string, ip string) (*[]byte, error) {
	return locationJson(requestUrl(cityUrl, key, ip))
}

func CountryLocationJson(key string, ip string) (*[]byte, error) {
	return locationJson(requestUrl(countryUrl, key, ip))
}

func CityLocation(key string, ip string) (*Location, error) {
	return location(requestUrl(cityUrl, key, ip))
}

func CountryLocation(key string, ip string) (*Location, error) {
	return location(requestUrl(countryUrl, key, ip))
}
