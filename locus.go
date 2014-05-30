package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	type Location struct {
		StatusCode    string
		StatusMessage string
		IpAddress     string
		CountryCode   string
		CountryName   string
		RegionName    string
		CityName      string
		ZipCode       string
		Latitude      string
		Longitude     string
		TimeZone      string
	}

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
		fmt.Printf("'Get' returned error: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("'ReadAll' returned an error: %s", err)
	}

	fmt.Printf("\nCity Location JSON: %s\n\n", body)

	var location_struct []Location
	err = json.Unmarshal(body, &location_struct)
	if err != nil {
		fmt.Printf("'json.Unmarshal' returned an error: %s\n\n", err)
	} else {
		fmt.Printf("City Location Struct: %v\n\n", location_struct)
	}

	var location_interface interface{}
	err = json.Unmarshal(body, &location_interface)
	if err != nil {
		fmt.Printf("'json.Unmarshal' returned an error: %s\n\n", err)
	} else {
		fmt.Printf("City Location Interface: %v\n", location_interface)
	}
}
