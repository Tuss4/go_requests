// An exercise with net/http and encoding/json
// Retrieves my profile
// Retrieves a list of my repos
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

type Repo struct {
	Html_url    string
	Name        string
	Id          int
	Description string
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

	resp2, err := http.Get("https://api.github.com/users/tuss4/repos")
	if err != nil {
		log.Fatal(err)
	}
	if resp2.StatusCode != 200 {
		log.Fatal(resp.Status)
	}

	repos := []Repo{}
	err = json.NewDecoder(resp2.Body).Decode(&repos)
	defer resp2.Body.Close()
	fmt.Println(repos)
}
