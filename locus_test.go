package locus

import (
	"testing"
)

const country_precision_url_1 = `http://api.ipinfodb.com/v3/ip-country/?format=json&ip=1.1.1.1&key=apiKey`
const country_precision_url_2 = `http://api.ipinfodb.com/v3/ip-country/?format=json&ip=2.2.2.2&key=apiKey`

const city_precision_url_1 = `http://api.ipinfodb.com/v3/ip-city/?format=json&ip=1.1.1.1&key=apiKey`
const city_precision_url_2 = `http://api.ipinfodb.com/v3/ip-city/?format=json&ip=2.2.2.2&key=apiKey`

func TestRequestUrl(t *testing.T) {
	if url, err := requestUrl(`1.1.1.1`, `country`, `apiKey`); url != country_precision_url_1 || err != nil {
		t.Errorf("requestUrl(`1.1.1.1`, `country`, `apiKey`) returned (%v, %v).\n\tExpected  (%v, %v)", url, err, country_precision_url_1, nil)
	}

	if url, err := requestUrl(`1.1.1.1`, `city`, `apiKey`); url != city_precision_url_1 || err != nil {
		t.Errorf("requestUrl(`1.1.1.1`, `city`, `apiKey`) returned (%v, %v).\n\tExpected  (%v, %v)", url, err, city_precision_url_1, nil)
	}
}

// Placeholders until I mock http requests in case the IP's location changes, etc.
func TestLookupLocation(t *testing.T) {

}

func TestLookupLocations(t *testing.T) {

}

func TestLookupLocationsFile(t *testing.T) {

}
