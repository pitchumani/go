package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

/*
getNav response json example:
{
    "meta":{
        "fund_house":"PPFAS Mutual Fund",
        "scheme_type":"360 ONE Mutual Fund",
        "scheme_category":"Formerly Known as IIFL Mutual Fund",
        "scheme_code":122639,
        "scheme_name":"Parag Parikh Flexi Cap Fund - Direct Plan - Growth"
    },
    "data":[
        {
            "date":"07-03-2024",
            "nav":"75.46660"
        }
    ],
    "status":"SUCCESS"
}
*/
type NData struct {
	Date string `json:"date"`
	Nav string `json:"nav"`
}
type NavData struct {
	Meta struct {
		FundHouse string `json:"fund_house"`
		SchemeType string `json:"scheme_type"`
		SchemeCategory string `json:"scheme_type"`
		SchemeCode int `json:"scheme_code"`
		SchemeName string `json:"scheme_name"`
	}
	NData []NData `json:"data"`
	Status string `json:"status"`
}

func decodeNavResponseJson(content []byte) NavData {
	var navdata NavData
	err := json.Unmarshal(content, &navdata)
	checkError(err)
	return navdata
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// get nav for given mutual fund
// use mfapi.in apis
func getNav(schemeCode int) NavData {
    latest_nav_url := "https://api.mfapi.in/mf/" + strconv.Itoa(schemeCode) + "/latest"

    // set http proxy to nil - computers in organization may have http proxy set,
    // which may not allow http requests from apps.
    transport := http.DefaultTransport
    transport.(*http.Transport).Proxy = nil
    client := &http.Client{
        Transport: transport,
    }
    res, err := client.Get(latest_nav_url)
    if err != nil {
        panic(err)
    }

    defer res.Body.Close()
    bytes, err := io.ReadAll(res.Body)
    if err != nil {
        panic(err)
    }
    contents := string(bytes)
	fmt.Println(contents)
	return decodeNavResponseJson(bytes) //contents)
}

func main() {
	fmt.Println("Fetch mutual fund's nav data")
	// Parag Parikh Long Term Equity Fund - Direct Growth - 122639
	nav := getNav(122639)
	fmt.Printf("%12v | %12v | %10v\n", "Date", "Scheme Code", "NAV")
	fmt.Printf("%12v | %12v | %10v\n", nav.NData[0].Date, nav.Meta.SchemeCode, nav.NData[0].Nav)
}
