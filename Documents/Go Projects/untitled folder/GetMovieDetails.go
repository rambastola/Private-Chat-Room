package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type search struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
}

type respData struct {
	Search []search `json:"Search"`
}

type movieInfo struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Actor    string `json:"Actors"`
	Country  string `json:"Country"`
	Awards   string `json:"Awards"`
}

func getData(movieName string) string {
	var data respData

	resp, err := http.Get("http://www.omdbapi.com/?s=" + movieName + "&apikey=99a10a4f")
	if err != nil {
		fmt.Println("Error while getting a response. %s", err)
	}
	defer resp.Body.Close()
	read, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading Bytes. %s", err)
	}

	e := json.Unmarshal(read, &data)
	if e != nil {
		fmt.Println("Error while decoding data. %s", e)
	}
	length := (len(data.Search))
	fmt.Println("%d serach results were found. %s", length)

	fmt.Println(data.Search)

	var choose int
	fmt.Println("\n Which movie are you looking for? ie- For the second one, type '2'.")
	fmt.Scan(&choose)
	choose--
	return data.Search[choose].ImdbID //return IMDB ID of chosen movie

}

func movieDetails(choose string) {
	var details movieInfo
	endpoint := "http://www.omdbapi.com/?i=" + choose + "&apikey=99a10a4f"
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("There was an error getting response. %s", err)
	}
	defer resp.Body.Close()
	read, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Unable to read the response. %s", err)
	}
	err = json.Unmarshal(read, &details)
	if err != nil {
		fmt.Println("Unable to unmarshal. %s", err)
	}
	fmt.Println(details)

}

func main() {
	fmt.Println("Type the movie you are searching for!")
	var text string
	fmt.Scan(&text)         //read input from standard input
	imdbID := getData(text) //stores the movieID
	movieDetails(imdbID)
}
