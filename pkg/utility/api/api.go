package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type LocationInfo struct {
	IP                 string  `json:"ip"`
	Network            string  `json:"network"`
	Version            string  `json:"version"`
	City               string  `json:"city"`
	Region             string  `json:"region"`
	RegionCode         string  `json:"region_code"`
	Country            string  `json:"country"`
	CountryName        string  `json:"country_name"`
	CountryCode        string  `json:"country_code"`
	CountryCodeISO3    string  `json:"country_code_iso3"`
	CountryCapital     string  `json:"country_capital"`
	CountryTLD         string  `json:"country_tld"`
	ContinentCode      string  `json:"continent_code"`
	InEU               bool    `json:"in_eu"`
	Postal             string  `json:"postal"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Timezone           string  `json:"timezone"`
	UTCOffset          string  `json:"utc_offset"`
	CountryCallingCode string  `json:"country_calling_code"`
	Currency           string  `json:"currency"`
	CurrencyName       string  `json:"currency_name"`
	Languages          string  `json:"languages"`
	CountryArea        float64 `json:"country_area"`
	CountryPopulation  int     `json:"country_population"`
	ASN                string  `json:"asn"`
	Org                string  `json:"org"`
}

func GetLocationInfo(ip string) (*LocationInfo, error) {
	url := fmt.Sprintf("https://ipapi.co/%s/json/", ip)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var locationInfo LocationInfo
	err = json.Unmarshal(body, &locationInfo)
	if err != nil {
		return nil, fmt.Errorf("JSON çözümleme hatası: %v", err)
	}

	return &locationInfo, nil
}
