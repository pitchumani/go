package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)
func getLatestNAV() {
		nav_url := "https://www.amfiindia.com/spages/NAVAll.txt"
		transport := http.DefaultTransport
		transport.(*http.Transport).Proxy = nil
		client := &http.Client{
			Transport: transport,
		}
		res, err := client.Get(nav_url)
		if err != nil {
			panic(err)
		}
		//fmt.Println("Response type: %T\n", res)
		defer res.Body.Close()
		bytes, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		contents := string(bytes)
		file, err := os.Create("nav.txt")
		if err != nil {
			panic(err)
		}
		io.WriteString(file, contents)
}

func getHistoricalNAV() {
	transport := http.DefaultTransport
	transport.(*http.Transport).Proxy = nil
	client := &http.Client{
		Transport: transport,
	}
	// API for historical nav for a mf
	// http://portal.amfiindia.com/DownloadNAVHistoryReport_Po.aspx?mf=<AMC Code>&tp=1&frmdt=31-Jan-2018&todt=31-Jan-2018
	nav_amc_date_api := "http://portal.amfiindia.com/DownloadNAVHistoryReport_Po.aspx?mf=21&tp=1&frmdt=31-Jan-2018&todt=31-Jan-2018"
	res, err := client.Get(nav_amc_date_api)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	contents := string(bytes)
	file, err := os.Create("nav-amc-date.txt")
	if err != nil {
		panic(err)
	}
	io.WriteString(file, contents)
}

func main() {
	fmt.Println("Fetch mutual fund's nav data")
  getLatestNAV()
	//getHistoricalNAV()
}
