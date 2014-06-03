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

func open(filename string) *os.File {
  file, err := os.Open(filename)
  check(err)

  return file
}

func scanner(file *os.File) *bufio.Scanner {
  return bufio.NewScanner(file)
}

func read_line(scanner *bufio.Scanner) string {
	check(scanner.Err())
	scanner.Scan()
	line := scanner.Text()
	check(scanner.Err())

	return line
}

func main() {
	fmt.Println("Locus: GeoIP Lookup")

	api_file := open("api")
	defer api_file.Close()
	api_key := read_line(scanner(api_file))

	ip_file := open("ips")
	defer ip_file.Close()
	ip := read_line(scanner(ip_file))

	request_url := fmt.Sprintf("http://api.ipinfodb.com/v3/ip-city/?key=%s&ip=%s&format=json", api_key, ip)

	resp, err := http.Get(request_url)
	defer resp.Body.Close()
	check(err)

	var location Location
	err = json.NewDecoder(resp.Body).Decode(&location)
	check(err)

	fmt.Printf("\n%v\n", location)
}
