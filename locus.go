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
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Locus: GeoIP Lookup\n")

	fmt.Printf("Enter your ipinfodb API key: ")
	key, err := reader.ReadString('\n'); check(err)
	key = strings.Trim(key, "\n")

	fmt.Printf("Enter an IP Address to lookup: ")
	ip, err := reader.ReadString('\n')
	ip = strings.Trim(ip, "\n"); check(err)

	request_url := fmt.Sprintf("http://api.ipinfodb.com/v3/ip-city/?key=%s&ip=%s&format=json", key, ip)

	resp, err := http.Get(request_url)
	defer resp.Body.Close(); check(err)

	var location Location
	err = json.NewDecoder(resp.Body).Decode(&location); check(err)

	fmt.Printf("\n%v\n", location)
}
