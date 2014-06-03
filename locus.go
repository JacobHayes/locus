package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

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

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Locus: GeoIP Lookup\n")

	api_file, err := os.Open("api")
	defer api_file.Close()
	check(err)
	api_scanner := bufio.NewScanner(api_file)
	check(err)
	api_scanner.Scan()
	if err = api_scanner.Err(); err != nil {
		log.Fatal(err)
	}
	api_key := api_scanner.Text()

	ip_file, err := os.Open("ips")
	defer ip_file.Close()
	check(err)
	ip_scanner := bufio.NewScanner(ip_file)
	check(err)
	ip_scanner.Scan()
	if err = ip_scanner.Err(); err != nil {
		log.Fatal(err)
	}
	ip := ip_scanner.Text()

	request_url := fmt.Sprintf("http://api.ipinfodb.com/v3/ip-city/?key=%s&ip=%s&format=json", api_key, ip)

	resp, err := http.Get(request_url)
	defer resp.Body.Close()
	check(err)

	var location Location
	err = json.NewDecoder(resp.Body).Decode(&location)
	check(err)

	fmt.Printf("\n%v\n", location)
}
