package locus

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
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

func requestUrl(ip string, precision string, key string) (string, error) {
	baseUrl := countryUrl
	if strings.ToLower(precision) == "city" {
		baseUrl = cityUrl
	}

	var request *url.URL
	request, err := url.Parse(baseUrl)
	if err != nil {
		return ``, err
	}

	params := url.Values{}
	params.Set(`ip`, ip)
	params.Set(`key`, key)
	params.Set(`format`, `json`)
	request.RawQuery = params.Encode()

	return request.String(), nil
}

func lookupLocation(ip string, precision string, key string) (Location, error) {
	location := Location{}
	request, err := requestUrl(ip, precision, key)
	if err != nil {
		return Location{}, err
	}

	resp, err := http.Get(request)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	raw_json, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	err = json.Unmarshal(raw_json, &location)
	if err != nil {
		return Location{}, err
	}

	return location, nil
}

func lookupLocations(ips []string, precision string, key string) ([]Location, error) {
	locations := make([]Location, len(ips))
	var err error
	for i, ip := range ips {
		locations[i], err = LookupLocation(ip, precision, key)
		if err != nil {
			return nil, err
		}
	}

	return locations, nil
}

func lookupLocationsFile(filename string, precision string, key string) ([]Location, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ips := make([]string, 0)
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		if scanner.Err() != nil {
			return nil, scanner.Err()
		}

		ips = append(ips, scanner.Text())
	}

	return LookupLocations(ips, precision, key)
}

// Public API

func LookupLocation(ip string, precision string, key string) (Location, error) {
	return lookupLocation(ip, precision, key)
}

func LookupLocations(ips []string, precision string, key string) ([]Location, error) {
	return lookupLocations(ips, precision, key)
}

func LookupLocationsFile(filename string, precision string, key string) ([]Location, error) {
	return lookupLocationsFile(filename, precision, key)
}
