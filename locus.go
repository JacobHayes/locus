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

// Location contains fields common to an IP's geolocation.
// When using country precision, only StatusCode, StatusMessage, IpAddress, CountryCode, and CountryName are present.
// When using city precision, all fields are present.
type Location struct {
	CityName      string `json:"cityName"`
	CountryCode   string `json:"countryCode"`
	CountryName   string `json:"countryName"`
	IpAddress     string `json:"ipAddress"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	RegionName    string `json:"regionName"`
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
	TimeZone      string `json:"timeZone"`
	ZipCode       string `json:"zipCode"`
}

// LookupLocation gets the geolocation of the provided IP with the given precision.
// It returns a Location struct containing geolocation data and any encountered error.
func LookupLocation(ip string, precision string, key string) (Location, error) {
	return lookupLocation(ip, precision, key)
}

// LookupLocations gets the geolocation of the provided IPs with the given precision.
// It returns a slice of Location structs containing geolocation data and any encountered error.
func LookupLocations(ips []string, precision string, key string) ([]Location, error) {
	return lookupLocations(ips, precision, key)
}

// LookupLocationsFile gets the geolocation of the IPs in the provided file with the given precision.
// The file format expects a single IP address per line.
// It returns a slice of Location structs containing geolocation data and any encountered error.
func LookupLocationsFile(filename string, precision string, key string) ([]Location, error) {
	return lookupLocationsFile(filename, precision, key)
}

// requestUrl creates the URL used for a lookup request.
// It returns a string containing a valid URL and any encountered error.
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

// See the documentation for LookupLocation.
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

// See the documentation for LookupLocations.
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

// See the documentation for LookupLocationsFile.
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
