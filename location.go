package locus

import "net"

// Location contains fields common to an IP's geolocation.
// The IP and Precision members are guaranteed to be present, however other members may not, based on the precision and any encountered errors.
// City precision will populate all fields, however Country precision will only populate the following fields: {StatusCode, StatusMessage, CountryCode, CountryName}
type Location struct {
	IP            net.IP
	Precision     Precision
	CityName      string `json:"cityName"`
	CountryCode   string `json:"countryCode"`
	CountryName   string `json:"countryName"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	RegionName    string `json:"regionName"`
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
	TimeZone      string `json:"timeZone"`
	ZipCode       string `json:"zipCode"`
}

func (l Location) Equals(x Location) bool {
	return l.IP.String() == x.IP.String() &&
		l.Precision == x.Precision &&
		l.CityName == x.CityName &&
		l.CountryCode == x.CountryCode &&
		l.CountryName == x.CountryName &&
		l.Latitude == x.Latitude &&
		l.Longitude == x.Longitude &&
		l.RegionName == x.RegionName &&
		l.StatusCode == x.StatusCode &&
		l.StatusMessage == x.StatusMessage &&
		l.TimeZone == x.TimeZone &&
		l.ZipCode == x.ZipCode
}
