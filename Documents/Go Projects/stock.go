package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type apiStruct struct {
	MetaData struct {
		Symbol    string `json:"2. Symbol"`         // stock symbol
		Refreshed string `json:"3. Last Refreshed"` //last refreshed
		TimeZone  string `json:"6. Time Zone"`      //timezone
	} `json:"Meta Data"`
}

type dateStruct struct {
	MonthlyAdjustedTimeSeries struct {
		Day struct {
			Open   string `json:"1. open"`   //opened amount
			High   string `json:"2. high"`   //day's high
			Low    string `json:"3. low"`    //day's low
			Close  string `json:"4. close"`  //day's close
			Volume string `json:"6. volume"` //day's volume
		} `json:"2018-09-10"`
	} `json:"Weekly Adjusted Time Series"`
}

func get(smbl string) {

	var apiData apiStruct //calling struct
	var apiDay dateStruct

	apiKey := "L3S8N827U95W5NM7" // apiKEY
	symbol := smbl               // Stock symbol (passed through argument)
	endpoint := "https://www.alphavantage.co/query?function=TIME_SERIES_WEEKLY_ADJUSTED&symbol=" + symbol + "&apikey=" + apiKey

	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("CHeck URL?")
	}
	defer resp.Body.Close()

	read, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading the response! %s", err)
	}

	// fmt.Println(string(read))

	er := json.Unmarshal(read, &apiData)
	if er != nil {
		fmt.Println("Unable to unmarshall.")
	}

	er = json.Unmarshal(read, &apiDay)
	if er != nil {
		fmt.Println("Unmarshalling error for apiDay struct. %s", er)
	}

	fmt.Println(apiData.MetaData)
	fmt.Println(apiDay.MonthlyAdjustedTimeSeries.Day) // prints out open, high, low, close, volume

}

func main() {
	get("AMZN") // Price of Amazon stock

}
