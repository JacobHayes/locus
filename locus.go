package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Location struct {
	StatusCode   string `json: "statusCode"`
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Locus: GeoIP Lookup\n")

	fmt.Printf("Enter your ipinfodb API key: ")
	key, err := reader.ReadString('\n')
	key = strings.Trim(key, "\n")
	if err != nil {
		fmt.Printf("\nKey failed...\n")
	}

	fmt.Printf("Enter an IP Address to lookup: ")
	ip, err := reader.ReadString('\n')
	ip = strings.Trim(ip, "\n")
	if err != nil {
		fmt.Printf("\nIP failed...\n\n")
	}

	request_url := fmt.Sprintf("http://api.ipinfodb.com/v3/ip-city/?key=%s&ip=%s&format=json", key, ip)

	resp, err := http.Get(request_url)
	if err != nil {
		resp.Body.Close()
		log.Fatalf("'Get' returned error: %s", err)
	}
	defer resp.Body.Close()

	var location Location
	if err = json.NewDecoder(resp.Body).Decode(&location); err != nil {
		log.Fatalf("'Decode' returned an error: %v", err)
	} else {
		fmt.Printf("\n%v", location)
	}
}
