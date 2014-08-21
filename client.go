package gofreegeoipclient

import (
	"encoding/json"
	"net/http"
)

type Country struct {
	Ip           string
	Country_code string
	Country_name string
	Region_code  string
	Region_name  string
	City         string
	Zipcode      string
	Latitude     int
	Longitude    int
	Metro_code   string
	Area_code    string
}

func GetCountryForIP(ip string) (Country, error) {
	res, err := http.Get("https://freegeoip.net/json/" + ip)
	if err != nil {
		return Country{}, err
	}
	var c Country
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&c); err != nil {
		return Country{}, err
	}
	return c, nil
}
