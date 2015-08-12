package locus

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type GeoLookup interface {
	IP(ip net.IP, precision Precision) (Location, error)
}

type Precision int

const (
	City Precision = iota
	Country
)

var (
	cityURL    *url.URL
	countryURL *url.URL

	InvalidAPIKey    error
	InvalidIPAddress error
	LookupError      error
)

func init() {
	var err error

	if cityURL, err = url.Parse(`http://api.ipinfodb.com/v3/ip-city/`); err != nil {
		log.Panicf("locus: init: %v", err)
	}

	if countryURL, err = url.Parse(`http://api.ipinfodb.com/v3/ip-country/`); err != nil {
		log.Panicf("locus: init: %v", err)
	}

	InvalidAPIKey = errors.New("Invalid API Key")
	LookupError = errors.New("IP Lookup Error: inspect Location.StatusCode")
	InvalidIPAddress = errors.New("Invalid IP Address")
}

// Locus provides the default implementation of the GeoLookup interface with extra accessor methods for implementation specific functionality.
type Locus struct {
	sync.Mutex

	apiKey     string
	httpClient http.Client
}

// Ensure Locus implements the GeoLookup interface
var _ GeoLookup = (*Locus)(nil)

func New(apiKey string) *Locus {
	return &Locus{
		apiKey: apiKey,
		httpClient: http.Client{
			Timeout: time.Second * 10,
		},
	}
}

// IP gets the geolocation of the provided IP with the given precision.
// It returns a Location struct.
func (l *Locus) IP(ip net.IP, precision Precision) (Location, error) {
	resp, err := l.httpClient.Get(resourceURL(ip, precision, l.apiKey))
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	location := Location{
		IP:        ip,
		Precision: precision,
	}
	if err = json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return Location{}, err
	}

	if location.StatusCode == "ERROR" {
		switch location.StatusMessage {
		case "Invalid API key.":
			return location, InvalidAPIKey
		case "Invalid IP Address.":
			return location, InvalidIPAddress
		default:
			return location, LookupError
		}
	}

	return location, nil
}

func (l *Locus) Timeout() time.Duration {
	l.Lock()
	defer l.Unlock()

	return l.httpClient.Timeout
}

func (l *Locus) SetTimeout(n time.Duration) {
	l.Lock()
	defer l.Unlock()

	l.httpClient.Timeout = n
}

// resourceURL creates the URL used for a lookup request.
// It returns a string containing a valid URL.
func resourceURL(ip net.IP, precision Precision, key string) string {
	baseURL := countryURL
	if precision == City {
		baseURL = cityURL
	}

	params := url.Values{}
	params.Set(`ip`, ip.String())
	params.Set(`key`, key)
	params.Set(`format`, `json`)
	baseURL.RawQuery = params.Encode()

	return baseURL.String()
}
