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

func locationJson(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func location(url string) (*Location, error) {
	location := &Location{}

	raw_json, err := locationJson(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw_json, location)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func requestUrl(ip string, precision string, key string) string {
	baseUrl := countryUrl
	if strings.ToLower(precision) == "city" {
		baseUrl = cityUrl
	}

	return strings.Join([]string{baseUrl, `?format=json`, `&ip=`, ip, `&key=`, key}, ``)
}

// Public API

func LookupLocationJson(ip string, precision string, key string) (string, error) {
	location, err := locationJson(requestUrl(ip, precision, key))
	return string(location[:]), err
}

func LookupLocation(ip string, precision string, key string) (*Location, error) {
	return location(requestUrl(ip, precision, key))
}

func BulkLookupLocationJSON(ips []string, precision string, key string) ([]string, error) {
	locations := make([]string, len(ips)-1)
	var err error
	for i, ip := range ips {
		locations[i], err = LookupLocationJson(ip, precision, key)
		if err != nil {
			return nil, err
		}
	}

	return locations, nil
}

func BulkLookupLocation(ips []string, precision string, key string) ([]*Location, error) {
	locations := make([]*Location, len(ips)-1)
	var err error
	for i, ip := range ips {
		locations[i], err = LookupLocation(ip, precision, key)
		if err != nil {
			return nil, err
		}
	}

	return locations, nil
}
