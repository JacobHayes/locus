package locus

import (
	"testing"
)

const ip_1 = `1.1.1.1`
const ip_2 = `2.2.2.2`

const country_url_1 = `http://api.ipinfodb.com/v3/ip-country/?format=json&ip=1.1.1.1&key=apiKey`
const country_url_2 = `http://api.ipinfodb.com/v3/ip-country/?format=json&ip=2.2.2.2&key=apiKey`
const city_url_1 = `http://api.ipinfodb.com/v3/ip-city/?format=json&ip=1.1.1.1&key=apiKey`
const city_url_2 = `http://api.ipinfodb.com/v3/ip-city/?format=json&ip=2.2.2.2&key=apiKey`

func TestRequestUrl(t *testing.T) {
	if url, err := requestUrl(ip_1, `country`, `apiKey`); url != country_url_1 || err != nil {
		t.Errorf("requestUrl(`%v`, `country`, `apiKey`) returned (%v, %v).\n\tExpected  (%v, %v)", ip_1, url, err, country_url_1, nil)
	}

	if url, err := requestUrl(ip_1, `city`, `apiKey`); url != city_url_1 || err != nil {
		t.Errorf("requestUrl(`%v`, `city`, `apiKey`) returned (%v, %v).\n\tExpected  (%v, %v)", ip_1, url, err, city_url_1, nil)
	}
}

// BUG(JacobHayes): Need to mock http requests to not require real API key and allow for IP's changing location, etc.
func TestLookupLocation(t *testing.T) {
}

// BUG(JacobHayes): Need to mock http requests to not require real API key and allow for IP's changing location, etc.
func TestLookupLocations(t *testing.T) {

}

// BUG(JacobHayes): Need to mock http requests to not require real API key and allow for IP's changing location, etc.
func TestLookupLocationsFile(t *testing.T) {

}
