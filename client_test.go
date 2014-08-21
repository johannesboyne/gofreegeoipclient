package gofreegeoipclient

import "testing"

func TestGetCountryForIP(t *testing.T) {
	if country, _ := GetCountryForIP("92.50.115.170"); country.Country_code != "DE" {
		t.Errorf("GetCountryForIP(%v) = %v, want %v", "92.50.115.170", country.Country_code, "DE")
	}
	if country, err := GetCountryForIP("92.50.115.a"); err == nil {
		t.Errorf("GetCountryForIP(%v) = %v, want %v", "92.50.115.a", country, nil)
	}
}
