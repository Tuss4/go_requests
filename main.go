package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Profile struct {
	Login    string
	Id       int
	Url      string
	Html_url string
}

func main() {
	var p Profile
	resp, err := http.Get("https://api.github.com/users/tuss4")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal(resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(p)
}
