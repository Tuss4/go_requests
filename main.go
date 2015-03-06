package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
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
	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Status)
	defer resp.Body.Close()
	fmt.Println(p.Login)
}
