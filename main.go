package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/codegangsta/martini"
)

type Configuration struct {
	Routes []Route
}

type Route struct {
	Verb        string
	Source      string
	Destination string
}

func main() {
	// Config
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
	}
	config := new(Configuration)
	err = json.Unmarshal(file, config)
	if err != nil {
		fmt.Println(err)
	}

	// Martini
	m := martini.Classic()
	for _, route := range config.Routes {
		switch route.Verb {
		case "get":
			fmt.Printf("Mapping GET %v to %v\n", route.Source, route.Destination)
			m.Get(route.Source, handler(route.Destination))
		case "post":
			fmt.Printf("Mapping POST %v to %v\n", route.Source, route.Destination)
			m.Post(route.Source, handler(route.Destination))
		}
	}
	m.Run()
}

// Reverse Proxy Handler
func handler(destination string) func(http.ResponseWriter, *http.Request) {
	url, err := url.Parse(destination)
	if err != nil {
		panic(err)
	}
	p := httputil.NewSingleHostReverseProxy(url)
	return func(w http.ResponseWriter, r *http.Request) {
		p.ServeHTTP(w, r)
	}
}
