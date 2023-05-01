package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var (
	addr  = "http://172.27.195.147:8080"
	paths = map[string]string{
		"/selections":            "GET",
		"/selections/1":          "GET",
		"/persons/1":             "GET",
		"/films/1":               "GET",
		"/series/1":              "GET",
		"/search?query=чернобы":  "GET",
		"/genres":                "GET",
		"/genres/1":              "GET",
		"/user/profile":          "GET",
		"/favorites/content":     "GET",
		"/user/signin":           "POST",
		"/user/signup":           "POST",
		"/user/logout":           "POST",
		"/user/update":           "POST",
		"/favorites/content/add": "POST",
	}
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		pathsSlice := make([]string, 0, len(paths))
		for k := range paths {
			pathsSlice = append(pathsSlice, k)
		}
		path := pathsSlice[rand.Intn(len(pathsSlice))]
		method := paths[path]

		url := fmt.Sprintf("%s%s", addr, path)
		fmt.Println("Sending", method, "request to", url)

		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			fmt.Printf("[%s %s] ERROR: %s\n", method, url, err)
			continue
		}

		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			fmt.Printf("[%s %s] ERROR: %s\n", method, url, err)
		} else {
			fmt.Printf("[%s %s] STATUS: %s\n", method, url, resp.Status)
			resp.Body.Close()
		}
		time.Sleep(700 * time.Millisecond)
	}
}
